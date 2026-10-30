// Package booklooker is a client for the Booklooker REST API 2.0.
package booklooker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

const DefaultBaseURL = "https://api.booklooker.de/2.0"

// tokenLifetime stays below the 10 minutes after which Booklooker drops an unused token.
const tokenLifetime = 9 * time.Minute

// Client authenticates on the first call and renews the token when it has expired.
// It is safe for concurrent use.
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	mutex      sync.Mutex
	token      string
	used       time.Time
}

// Error is a response with the status NOK, Code is the returned value like NOT_FOUND.
type Error struct {
	Interface string
	Code      string
}

type request struct {
	method      string
	path        string
	params      url.Values
	body        []byte
	contentType string
}

type response struct {
	Status      string          `json:"status"`
	ReturnValue json.RawMessage `json:"returnValue"`
}

func New(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: DefaultBaseURL,
		// File uploads may have up to 80 MB.
		HTTPClient: &http.Client{Timeout: 5 * time.Minute},
	}
}

func (e *Error) Error() string {
	return "booklooker: " + e.Interface + ": " + e.Code
}

// Account returns the seller account of the API key and whether it can sell the media type.
func (c *Client) Account(mediaType MediaType) (Account, error) {
	var account Account
	params := url.Values{"mediaType": {strconv.Itoa(int(mediaType))}}
	err := c.call(request{method: http.MethodGet, path: "account", params: params}, &account)
	return account, err
}

func (c *Client) call(r request, result any) error {
	value, err := c.send(r)
	if expired(err) {
		c.reset()
		value, err = c.send(r)
	}
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	return json.Unmarshal(value, result)
}

// callText returns the value as text, numbers included, because the API documents most values only by example.
func (c *Client) callText(r request) (string, error) {
	var value json.RawMessage
	err := c.call(r, &value)
	if err != nil {
		return "", err
	}
	var text string
	err = json.Unmarshal(value, &text)
	if err != nil {
		return string(value), nil
	}
	return text, nil
}

func (c *Client) send(r request) (json.RawMessage, error) {
	token, err := c.currentToken()
	if err != nil {
		return nil, err
	}
	params := url.Values{"token": {token}}
	for key, values := range r.params {
		params[key] = values
	}
	r.params = params
	return c.do(r)
}

func (c *Client) currentToken() (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.token != "" && time.Since(c.used) < tokenLifetime {
		c.used = time.Now()
		return c.token, nil
	}
	value, err := c.do(request{method: http.MethodPost, path: "authenticate", params: url.Values{"apiKey": {c.APIKey}}})
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(value, &c.token)
	if err != nil {
		return "", err
	}
	c.used = time.Now()
	return c.token, nil
}

func (c *Client) reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.token = ""
}

func (c *Client) do(r request) (json.RawMessage, error) {
	httpRequest, err := http.NewRequest(r.method, c.BaseURL+"/"+r.path+"?"+r.params.Encode(), bytes.NewReader(r.body))
	if err != nil {
		return nil, err
	}
	if r.contentType != "" {
		httpRequest.Header.Set("Content-Type", r.contentType)
	}
	httpResponse, err := c.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("booklooker: %s: HTTP status %s", r.path, httpResponse.Status)
	}
	var result response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("booklooker: %s: %w", r.path, err)
	}
	if result.Status != "OK" {
		return nil, &Error{Interface: r.path, Code: errorCode(result.ReturnValue)}
	}
	return result.ReturnValue, nil
}

func errorCode(value json.RawMessage) string {
	var code string
	err := json.Unmarshal(value, &code)
	if err != nil {
		return string(value)
	}
	return code
}

// Booklooker may drop a token before tokenLifetime, the call is then repeated once with a new token.
func expired(err error) bool {
	var apiError *Error
	if !errors.As(err, &apiError) {
		return false
	}
	return apiError.Code == "TOKEN_EXPIRED" || apiError.Code == "TOKEN_UNKNOWN"
}

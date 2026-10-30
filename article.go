package booklooker

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ArticleListOptions are optional, the zero value lists the order numbers of all articles.
type ArticleListOptions struct {
	Field ArticleField
	// ShowPrice appends the price to each line, separated by a tab.
	ShowPrice bool
	// ShowStock lists only articles in stock with order number and stock and ignores the other options.
	ShowStock bool
	// MediaType restricts the list, nil lists all media types.
	MediaType *MediaType
}

// DeleteArticle marks an article for deletion, Booklooker deletes it within 2 hours.
func (c *Client) DeleteArticle(orderNo string) error {
	return c.call(request{method: http.MethodDelete, path: "article", params: url.Values{"orderNo": {orderNo}}}, nil)
}

// ArticleList returns one line per active article. Deleted articles stay in the list for 1 to 2 hours, ArticleStatus is current.
func (c *Client) ArticleList(options ArticleListOptions) ([]string, error) {
	params := url.Values{}
	if options.Field != "" {
		params.Set("field", string(options.Field))
	}
	if options.ShowPrice {
		params.Set("showPrice", "1")
	}
	if options.ShowStock {
		params.Set("showStock", "1")
	}
	if options.MediaType != nil {
		params.Set("mediaType", strconv.Itoa(int(*options.MediaType)))
	}
	text, err := c.callText(request{method: http.MethodGet, path: "article_list", params: params})
	if err != nil {
		return nil, err
	}
	return strings.FieldsFunc(text, func(r rune) bool { return r == '\n' }), nil
}

func (c *Client) ArticleStatus(orderNo string) (ArticleStatus, error) {
	text, err := c.callText(request{method: http.MethodGet, path: "article_status", params: url.Values{"orderNo": {orderNo}}})
	return ArticleStatus(text), err
}

// Search returns the found articles unchanged, because Booklooker documents neither the parameters nor the result.
func (c *Client) Search(params url.Values) (json.RawMessage, error) {
	var result json.RawMessage
	err := c.call(request{method: http.MethodGet, path: "search", params: params}, &result)
	return result, err
}

// DeleteImage deletes the image at the position, the following images move up by one.
func (c *Client) DeleteImage(orderNo string, position int) error {
	params := url.Values{"orderNo": {orderNo}, "position": {strconv.Itoa(position)}}
	return c.call(request{method: http.MethodDelete, path: "image", params: params}, nil)
}

// DeleteImages deletes all images of an article.
func (c *Client) DeleteImages(orderNo string) error {
	return c.call(request{method: http.MethodDelete, path: "image", params: url.Values{"orderNo": {orderNo}}}, nil)
}

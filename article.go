package booklooker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// centsInEuro converts the cents of an Article into the euros Booklooker reads.
const centsInEuro = 100

// Article is a book for CreateArticle and UpdateArticle, Booklooker checks it like the form on its website.
// UpdateArticle sends every field, an empty text clears it at Booklooker.
type Article struct {
	// OrderNo identifies the article, CreateArticle without one lets Booklooker assign one like BA00042.
	OrderNo     string
	Title       string
	Author      string
	Publisher   string
	Edition     string
	Year        string
	Place       string
	Binding     string
	Condition   Condition
	Description string
	// Language is a code like de or a name like deutsch, ValueList with ValueLanguage lists them.
	Language string
	ISBN     string
	Pages    string
	Format   string
	// Weight in grams, zero leaves it out. Without a weight Booklooker rejects an article whose shipping it cannot
	// calculate.
	Weight int64
	// Price in cents, at least 25.
	Price int64
	// Category is an ID of ValueList with ValueCategory, zero leaves it out.
	Category     int64
	CoverURL     string
	Keyword      string
	Note         string
	Unlimited    bool
	FirstEdition bool
	Signed       bool
}

// ArticleResult is the order number of the article and its status, a new or changed one is ArticlePending.
type ArticleResult struct {
	OrderNo string        `json:"orderNo"`
	Status  ArticleStatus `json:"articleStatus"`
}

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

// Search returns the found articles unchanged, because Booklooker documents the result only as a self-explanatory
// list. Each call counts against the search quota of 50 calls in 10 minutes, conditionMin=2 finds only offers with
// light wear or better.
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

// CreateArticle offers a book. ORDER_NO_EXISTS answers an order number Booklooker knows, also a paused one.
func (c *Client) CreateArticle(article Article) (ArticleResult, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for name, values := range article.values() {
		err := writer.WriteField(name, values[0])
		if err != nil {
			return ArticleResult{}, fmt.Errorf("write the field %s of article %s: %w", name, article.OrderNo, err)
		}
	}
	err := writer.Close()
	if err != nil {
		return ArticleResult{}, fmt.Errorf("close the form of article %s: %w", article.OrderNo, err)
	}
	var result ArticleResult
	err = c.call(request{method: http.MethodPost, path: "article", body: body.Bytes(), contentType: writer.FormDataContentType()}, &result)
	return result, err
}

// UpdateArticle changes a book, the change takes effect with the next update run of Booklooker. NOT_EDITABLE answers
// a sold, deleted or paused article, IN_PROGRESS one that the update run handles right now.
func (c *Client) UpdateArticle(article Article) (ArticleResult, error) {
	values := article.values()
	values.Del("orderNo")
	r := request{
		method:      http.MethodPut,
		path:        "article",
		params:      url.Values{"orderNo": {article.OrderNo}},
		body:        []byte(values.Encode()),
		contentType: "application/x-www-form-urlencoded",
	}
	var result ArticleResult
	err := c.call(r, &result)
	return result, err
}

func (a Article) values() url.Values {
	values := url.Values{
		"title":        {a.Title},
		"author":       {a.Author},
		"publisher":    {a.Publisher},
		"edition":      {a.Edition},
		"year":         {a.Year},
		"place":        {a.Place},
		"binding":      {a.Binding},
		"condition":    {strconv.Itoa(int(a.Condition))},
		"description":  {a.Description},
		"language":     {a.Language},
		"isbn":         {a.ISBN},
		"pages":        {a.Pages},
		"format":       {a.Format},
		"price":        {fmt.Sprintf("%d.%02d", a.Price/centsInEuro, a.Price%centsInEuro)},
		"coverUrl":     {a.CoverURL},
		"keyword":      {a.Keyword},
		"note":         {a.Note},
		"unlimited":    {flag(a.Unlimited)},
		"firstEdition": {flag(a.FirstEdition)},
		"signed":       {flag(a.Signed)},
	}
	if a.OrderNo != "" {
		values.Set("orderNo", a.OrderNo)
	}
	if a.Weight > 0 {
		values.Set("weight", strconv.FormatInt(a.Weight, 10))
	}
	if a.Category > 0 {
		values.Set("category", strconv.FormatInt(a.Category, 10))
	}
	return values
}

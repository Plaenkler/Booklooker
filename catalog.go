package booklooker

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Value is an allowed value of a select field. A category has its subcategories, a medium of music its group.
type Value struct {
	ID         ValueID `json:"id"`
	Name       string  `json:"name"`
	Group      string  `json:"group"`
	Categories []Value `json:"categories"`
}

// ValueID is a number like the ID of a category or a text like the code de of a language.
type ValueID string

// PriceOverview counts the offers of an ISBN or EAN that Search finds, Conditions has all six states in order.
type PriceOverview struct {
	Offers     int               `json:"offers"`
	Conditions []ConditionOffers `json:"conditions"`
}

// ConditionOffers are the offers in exactly one state. Condition is empty for offers without a state of the scale,
// MinPrice in euros without shipping is empty without offers.
type ConditionOffers struct {
	Condition json.Number `json:"condition"`
	Name      string      `json:"name"`
	Offers    int         `json:"offers"`
	MinPrice  json.Number `json:"minPrice"`
}

func (v *ValueID) UnmarshalJSON(data []byte) error {
	var text string
	err := json.Unmarshal(data, &text)
	if err == nil {
		*v = ValueID(text)
		return nil
	}
	var number json.Number
	err = json.Unmarshal(data, &number)
	if err != nil {
		return err
	}
	*v = ValueID(number.String())
	return nil
}

// ValueList returns the allowed values in the order of the form of Booklooker. They change rarely, Booklooker asks
// to fetch them about once a day.
func (c *Client) ValueList(field ValueField, mediaType MediaType) ([]Value, error) {
	params := url.Values{"field": {string(field)}, "mediaType": {strconv.Itoa(int(mediaType))}}
	var values []Value
	err := c.call(request{method: http.MethodGet, path: "value_list", params: params}, &values)
	return values, err
}

// PriceOverview counts the book offers of an ISBN per state. Each call counts against the search quota of 50 calls
// in 10 minutes.
func (c *Client) PriceOverview(isbn string) (PriceOverview, error) {
	var overview PriceOverview
	err := c.call(request{method: http.MethodGet, path: "price_overview", params: url.Values{"isbn": {isbn}}}, &overview)
	return overview, err
}

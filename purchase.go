package booklooker

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Purchase buys articles as the authenticated user. Booklooker has to enable the interface first.
type Purchase struct {
	// ArticleIDs are up to 100 order numbers, also of different sellers, each seller gets an order of its own.
	ArticleIDs []string
	// PaymentMethods are tried in this order per seller, a method below the minimum order value is skipped.
	PaymentMethods []PaymentMethod
	// OrderToken identifies the purchase with 1 to 64 letters, digits, - or _.
	// A repeated token is rejected, so a purchase cannot be placed twice after a connection loss.
	OrderToken string
	// AcceptTerms has to be true, it accepts the terms and conditions of Booklooker.
	AcceptTerms bool
	// AddressID is the delivery address of all orders, 0 is the invoice address.
	AddressID int
	// RequireCompleteOrder only orders from a seller if all of its articles are available.
	RequireCompleteOrder bool
	// Message is sent to the sellers.
	Message string
}

// PurchaseResult lists every article ID once, either in InvalidArticles or at the order of its seller.
type PurchaseResult struct {
	InvalidArticles []string        `json:"invalidArticles"`
	Orders          []PurchaseOrder `json:"orders"`
}

// PurchaseOrder is the order at one seller. Amounts are in euro.
type PurchaseOrder struct {
	SellerID int `json:"sellerId"`
	// Status is NOK if no order was placed, Error then has the reason, e.g. NO_PAYMENT_METHOD_AVAILABLE.
	Status          string            `json:"status"`
	Error           string            `json:"error"`
	OrderID         int               `json:"orderId"`
	Price           json.Number       `json:"price"`
	ShippingCost    json.Number       `json:"shippingCost"`
	AdditionalCosts json.Number       `json:"additionalCosts"`
	TotalCost       json.Number       `json:"totalCost"`
	PaymentID       PaymentMethod     `json:"paymentId"`
	PaymentNotice   string            `json:"paymentNotice"`
	Articles        []PurchaseArticle `json:"articles"`
}

type PurchaseArticle struct {
	ArticleID string `json:"articleId"`
	// Status is NOK if the article is not part of the order, Error then has the reason, e.g. ARTICLE_SOLD.
	Status string `json:"status"`
	Error  string `json:"error"`
}

// Purchase returns an error only if Booklooker did not process the call, the orders of the result can still have failed.
func (c *Client) Purchase(purchase Purchase) (PurchaseResult, error) {
	params := url.Values{
		"articleId":   purchase.ArticleIDs,
		"orderToken":  {purchase.OrderToken},
		"acceptTerms": {flag(purchase.AcceptTerms)},
		"addressId":   {strconv.Itoa(purchase.AddressID)},
	}
	for _, method := range purchase.PaymentMethods {
		params.Add("paymentId", strconv.Itoa(int(method)))
	}
	if purchase.RequireCompleteOrder {
		params.Set("requireCompleteOrder", "1")
	}
	if purchase.Message != "" {
		params.Set("message", purchase.Message)
	}
	var result PurchaseResult
	err := c.call(request{method: http.MethodPost, path: "purchase", params: params}, &result)
	return result, err
}

func flag(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

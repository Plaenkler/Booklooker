package booklooker

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const dateLayout = "2006-01-02"

// Order is an order of a customer. Amounts are decimal numbers as text, e.g. "12.50".
type Order struct {
	OrderID   int    `json:"orderId"`
	OrderDate string `json:"orderDate"`
	OrderTime string `json:"orderTime"`
	// Status is one of the OrderStatus values.
	Status   OrderStatus `json:"status"`
	Email    string      `json:"email"`
	Phone    string      `json:"tel"`
	VATID    string      `json:"ustIdNr"`
	Comments string      `json:"comments"`
	// CalculatedShippingCost is only set if Booklooker could calculate it, e.g. from the weight.
	CalculatedShippingCost string        `json:"calculatedShippingCost"`
	PaymentID              PaymentMethod `json:"paymentId"`
	// PaymentConfirmed and TransactionID are only set for PayPal, the time has the format "2006-01-02 15:04:05".
	PaymentConfirmed string `json:"paymentConfirmed"`
	TransactionID    string `json:"transactionId"`
	// The bank account is only set for direct debit.
	AccountHolder                 string `json:"accountHolder"`
	AccountIBAN                   string `json:"accountIban"`
	AccountBIC                    string `json:"accountBic"`
	BuyerPositiveRatingPercentage string `json:"buyerPositiveRatingPercentage"`
	BuyerTotalRatingNumber        int    `json:"buyerTotalRatingNumber"`
	// BuyerUsername is the company name for dealers.
	BuyerUsername string `json:"buyerUsername"`
	// The commission without VAT, the current one considers cancellations.
	OriginalProvisionNet string  `json:"originalProvisionNet"`
	CurrentProvisionNet  string  `json:"currentProvisionNet"`
	InvoiceAddress       Address `json:"invoiceAddress"`
	// DeliveryAddress is empty if it is the invoice address.
	DeliveryAddress Address     `json:"deliveryAddress"`
	OrderItems      []OrderItem `json:"orderItems"`
}

type Address struct {
	Title             string `json:"title"`
	Name              string `json:"name"`
	FirstName         string `json:"firstName"`
	Company           string `json:"company"`
	AddressSupplement string `json:"addressSupplement"`
	Street            string `json:"street"`
	Zip               string `json:"zip"`
	City              string `json:"city"`
	Country           string `json:"country"`
}

type OrderItem struct {
	// OrderItemID is needed for CancelOrderItem.
	OrderItemID int    `json:"orderItemId"`
	Amount      int    `json:"amount"`
	OrderNo     string `json:"orderNo"`
	OrderTitle  string `json:"orderTitle"`
	// Author is the director for movies and the artist for music.
	Author string `json:"author"`
	// Infotext has to be enabled by the Booklooker support.
	Infotext string `json:"infotext"`
	// Note is the personal note or storage location.
	Note      string    `json:"note"`
	MediaType MediaType `json:"mediaType"`
	// SinglePrice is the price of one piece, TotalPriceRebated of the whole amount less discounts.
	SinglePrice       string `json:"singlePrice"`
	TotalPriceRebated string `json:"totalPriceRebated"`
	// Status is "Ordered" or "Canceled".
	Status string `json:"status"`
}

func (c *Client) Orders(date time.Time) ([]Order, error) {
	return c.orders(url.Values{"date": {date.Format(dateLayout)}})
}

// OrdersBetween returns the orders of an interval of at most one month, both days included.
func (c *Client) OrdersBetween(from, to time.Time) ([]Order, error) {
	return c.orders(url.Values{"dateFrom": {from.Format(dateLayout)}, "dateTo": {to.Format(dateLayout)}})
}

func (c *Client) Order(orderID int) (Order, error) {
	orders, err := c.orders(url.Values{"orderId": {strconv.Itoa(orderID)}})
	if err != nil {
		return Order{}, err
	}
	if len(orders) == 0 {
		return Order{}, fmt.Errorf("booklooker: order: order %d not found", orderID)
	}
	return orders[0], nil
}

// CancelOrder cancels a whole order without commission, at most 6 weeks after it was placed.
func (c *Client) CancelOrder(orderID int) error {
	return c.call(request{method: http.MethodPut, path: "order_cancel", params: orderParams(orderID)}, nil)
}

// CancelOrderItem cancels a single article of an order without commission, at most 6 weeks after it was placed.
func (c *Client) CancelOrderItem(orderItemID int, mediaType MediaType) error {
	params := url.Values{"orderItemId": {strconv.Itoa(orderItemID)}, "mediaType": {strconv.Itoa(int(mediaType))}}
	return c.call(request{method: http.MethodPut, path: "order_item_cancel", params: params}, nil)
}

// SendOrderMessage mails a message to the customer and returns its text.
// Payment information and reminders can be sent every 7 days, a reminder at the earliest 7 days after the order.
func (c *Client) SendOrderMessage(orderID int, messageType MessageType, additionalText string) (string, error) {
	params := orderParams(orderID)
	params.Set("messageType", string(messageType))
	if additionalText != "" {
		params.Set("additionalText", additionalText)
	}
	return c.callText(request{method: http.MethodPut, path: "order_message", params: params})
}

// RateOrder rates an order within 6 months, a given rating can only be improved.
func (c *Client) RateOrder(orderID int, rating Rating, comment string) error {
	params := orderParams(orderID)
	params.Set("rating", strconv.Itoa(int(rating)))
	if comment != "" {
		params.Set("comment", comment)
	}
	return c.call(request{method: http.MethodPut, path: "order_rating", params: params}, nil)
}

func (c *Client) SetOrderStatus(orderID int, status OrderStatus) error {
	params := orderParams(orderID)
	params.Set("status", string(status))
	return c.call(request{method: http.MethodPut, path: "order_status", params: params}, nil)
}

func (c *Client) orders(params url.Values) ([]Order, error) {
	var orders []Order
	err := c.call(request{method: http.MethodGet, path: "order", params: params}, &orders)
	return orders, err
}

func orderParams(orderID int) url.Values {
	return url.Values{"orderId": {strconv.Itoa(orderID)}}
}

package model

type OrderRequest struct {
	OrderID  string
	Date     string
	DateFrom string
	DateTo   string
}

type Order struct {
	OrderID                       int                 `json:"orderId"`
	OrderDate                     string              `json:"orderDate"`
	OrderTime                     string              `json:"orderTime"`
	Status                        string              `json:"status"`
	Email                         string              `json:"email"`
	CalculatedShippingCost        string              `json:"calculatedShippingCost"`
	PaymentID                     int                 `json:"paymentId"`
	BuyerPositiveRatingPercentage string              `json:"buyerPositiveRatingPercentage"`
	BuyerTotalRatingNumber        int                 `json:"buyerTotalRatingNumber"`
	BuyerUsername                 string              `json:"buyerUsername"`
	OriginalProvisionNet          string              `json:"originalProvisionNet"`
	CurrentProvisionNet           string              `json:"currentProvisionNet"`
	InvoiceAddress                OrderInvoiceAddress `json:"invoiceAddress"`
	OrderItems                    []OrderItem         `json:"orderItems"`
	PaymentConfirmed              string              `json:"paymentConfirmed,omitempty"`
	TransactionID                 string              `json:"transactionId,omitempty"`
}

type OrderInvoiceAddress struct {
	Title     string `json:"title"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Street    string `json:"street"`
	Zip       string `json:"zip"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

type OrderItem struct {
	OrderItemID       int    `json:"orderItemId"`
	Amount            int    `json:"amount"`
	OrderNo           string `json:"orderNo"`
	OrderTitle        string `json:"orderTitle"`
	Author            string `json:"author"`
	SinglePrice       string `json:"singlePrice"`
	TotalPriceRebated string `json:"totalPriceRebated"`
	MediaType         int    `json:"mediaType"`
	Status            string `json:"status"`
}

type ErrorResponse struct {
	Status      string      `json:"status"`
	ReturnValue interface{} `json:"returnValue"`
}

type OrderResponse struct {
	Status      string  `json:"status"`
	ReturnValue []Order `json:"returnValue"`
}

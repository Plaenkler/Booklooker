package model

type OrderRequest struct {
	OrderID  string `json:"orderId"`
	Date     string `json:"date"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

type Order struct {
	OrderID                       string `json:"orderId"`
	CalculatedShippingCosts       string `json:"calculatedShippingCost"`
	Comments                      string `json:"comments"`
	OrderItemID                   string `json:"orderItemId"`
	Amount                        string `json:"amount"`
	OrderNo                       string `json:"orderNo"`
	OrderTitle                    string `json:"orderTitle"`
	Author                        string `json:"author"`
	SinglePrice                   string `json:"singlePrice"`
	TotalPriceRebate              string `json:"totalPriceRebated"`
	MediaType                     string `json:"mediaType"`
	Status                        string `json:"status"`
	PaymentID                     string `json:"paymentId"`
	AccountHolder                 string `json:"accountHolder"`
	AccountIBAN                   string `json:"accountIban"`
	AccountBIC                    string `json:"accountBic"`
	PaymentConfirmed              string `json:"paymentConfirmed"`
	TransactionId                 string `json:"transactionId"`
	BuyerPositiveRatingPercentage string `json:"buyerPositiveRatingPercentage"`
	BuyerTotalRatingNumber        string `json:"buyerTotalRatingNumber"`
}

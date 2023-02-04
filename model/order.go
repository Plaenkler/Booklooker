package model

type OrderRequest struct {
	OrderID  string
	Date     string
	DateFrom string
	DateTo   string
}

type Order struct {
	OrderID                       string
	CalculatedShippingCosts       string
	Comments                      string
	OrderItemID                   string
	Amount                        string
	OrderNo                       string
	OrderTitle                    string
	Author                        string
	SinglePrice                   string
	TotalPriceRebate              string
	MediaType                     MediaType
	Status                        string
	PaymentID                     string
	AccountHolder                 string
	AccountIBAN                   string
	AccountBIC                    string
	PaymentConfirmed              string
	TransactionId                 string
	BuyerPositiveRatingPercentage string
	BuyerTotalRatingNumber        string
}

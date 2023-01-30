package models

const OrderPath = "order"

type OrderRequest struct {
	OrderID  string `json:"orderId"`
	Date     string `json:"date"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

type OrderResponse struct {
	Status string `json:"status"`
	// TODO: Add Order struct
}

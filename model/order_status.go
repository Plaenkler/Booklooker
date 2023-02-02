package model

type OrderStatusRequest struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

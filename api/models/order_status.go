package models

const OrderStatusPath = "order_status"

type OrderStatusRequest struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

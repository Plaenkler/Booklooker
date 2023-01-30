package models

const OrderStatusPath = "order_status"

type OrderStatusRequest struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

type OrderStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

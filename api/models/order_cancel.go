package models

const OrderCancelPath = "order_cancel"

type OrderCancelRequest struct {
	OrderID string `json:"orderId"`
}

type OrderCancelResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

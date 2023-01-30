package models

const OrderCancelPath = "order_cancel"

type OrderCancelRequest struct {
	OrderID string `json:"orderId"`
}

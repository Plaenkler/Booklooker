package models

const OrderItemCancelPath = "order_item_cancel"

type OrderItemCancelRequest struct {
	OrderItemID string `json:"orderItemId"`
	MediaType   byte   `json:"mediaType"`
}

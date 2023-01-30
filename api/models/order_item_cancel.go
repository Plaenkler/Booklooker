package models

const OrderItemCancelPath = "order_item_cancel"

type OrderItemCancelRequest struct {
	OrderItemId string `json:"orderItemId"`
	MediaType   byte   `json:"mediaType"`
}

type OrderItemCancelResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

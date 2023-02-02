package model

type OrderItemCancelRequest struct {
	OrderItemID string `json:"orderItemId"`
	MediaType   string `json:"mediaType"`
}

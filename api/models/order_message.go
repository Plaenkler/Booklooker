package models

const OrderMessagePath = "order_message"

type OrderMessageRequest struct {
	OrderID        string `json:"orderId"`
	MessageType    string `json:"messageType"`
	AdditionalText string `json:"additionalText,omitempty"`
}

type OrderMessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

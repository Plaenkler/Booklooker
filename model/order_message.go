package model

type OrderMessageRequest struct {
	OrderID        string `json:"orderId"`
	MessageType    string `json:"messageType"`
	AdditionalText string `json:"additionalText,omitempty"`
}

package model

type OrderMessageRequest struct {
	OrderID        string
	MessageType    MessageType
	AdditionalText string
}

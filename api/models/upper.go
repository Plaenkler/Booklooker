package models

// Upper response implemented by many endpoints
type GeneralResponse struct {
	Status      string `json:"status"`
	ReturnValue string `json:"returnValue,omitempty"`
}

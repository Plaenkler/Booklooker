package models

// Every other endpoint expects a token (?token=REST_API_TOKEN) obtained by authentication
const AuthenticatePath = "authenticate"

type AuthenticateRequest struct {
	APIKey string `json:"apiKey"`
}

type AuthenticateResponse struct {
	Status      string `json:"status"`
	ReturnValue string `json:"returnValue,omitempty"`
}

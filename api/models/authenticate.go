package models

// Every other endpoint expects a token (?token=REST_API_TOKEN) obtained by authentication
const AuthenticatePath = "authenticate"

type AuthenticateRequest struct {
	APIKey string `json:"apiKey"`
}

type AuthenticateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

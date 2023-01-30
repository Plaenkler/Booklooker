package models

// Every other endpoint expects a token (?token=REST_API_TOKEN) obtained by authentication
const AuthenticatePath = "authenticate"

type AuthenticateRequest struct {
	APIKey string `json:"apiKey"`
}

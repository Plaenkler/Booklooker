package models

const ImportStatusPath = "import_status"

type ImportStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

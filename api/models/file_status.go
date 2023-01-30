package models

const FileStatusPath = "file_status"

type FileStatusRequest struct {
	Filename string `json:"filename"`
}

type FileStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

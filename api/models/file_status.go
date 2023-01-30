package models

const FileStatusPath = "file_status"

type FileStatusRequest struct {
	Filename string `json:"filename"`
}

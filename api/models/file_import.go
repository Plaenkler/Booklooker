package models

const FileImportPath = "file_import"

type FileImportRequest struct {
	File      []byte `json:"file"`
	FileType  string `json:"fileType,omitempty"`
	DataType  byte   `json:"dataType,omitempty"`
	MediaType byte   `json:"mediaType,omitempty"`
	FormatID  int    `json:"formatId,omitempty"`
	Encoding  string `json:"encoding,omitempty"`
}

type FileImportResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

package model

type FileImportRequest struct {
	File      []byte `json:"file"`
	FileType  string `json:"fileType,omitempty"`
	DataType  string `json:"dataType,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	FormatID  string `json:"formatId,omitempty"`
	Encoding  string `json:"encoding,omitempty"`
}

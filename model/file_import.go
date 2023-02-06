package model

import "os"

type FileImportRequest struct {
	File      *os.File
	FileType  FileType
	DataType  DataType
	MediaType MediaType
	FormatID  string
	Encoding  Encoding
}

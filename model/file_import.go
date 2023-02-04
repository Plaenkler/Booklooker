package model

import "os"

type FileImportRequest struct {
	File      *os.File
	FileType  string
	DataType  int
	MediaType int
	FormatID  string
	Encoding  string
}

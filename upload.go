package booklooker

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// Upload is a file for ImportFile.
type Upload struct {
	// Name has to end with .zip, .txt, .tab, .csv or .xls, for images with .zip.
	Name     string
	Content  io.Reader
	FileType FileType
	DataType DataType
	// MediaType defaults to Books.
	MediaType MediaType
	// FormatID sets the order of the columns, 1 is the Booklooker format and 0 keeps the format of the last upload.
	FormatID int
	Encoding Encoding
}

// ImportFile uploads an article or image file of up to 80 MB, Booklooker processes the files in the order of upload.
func (c *Client) ImportFile(upload Upload) error {
	r, err := uploadRequest(upload)
	if err != nil {
		return err
	}
	return c.call(r, nil)
}

// FileStatus refers to the last uploaded file with that name.
func (c *Client) FileStatus(filename string) (FileStatus, error) {
	text, err := c.callText(request{method: http.MethodGet, path: "file_status", params: url.Values{"filename": {filename}}})
	return FileStatus(text), err
}

// FileErrors lists the rejected articles of a completely processed file.
// Each line has the error ID, record, order number, title and details.
func (c *Client) FileErrors(filename string) (string, error) {
	params := url.Values{"filename": {filename}, "showErrors": {"1"}}
	return c.callText(request{method: http.MethodGet, path: "file_status", params: params})
}

// ImportStatus returns the number of uploaded files that are not processed yet.
func (c *Client) ImportStatus() (int, error) {
	text, err := c.callText(request{method: http.MethodGet, path: "import_status"})
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(text)
}

// The body is buffered, so the upload can be repeated with a new token.
func uploadRequest(upload Upload) (request, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", upload.Name)
	if err != nil {
		return request{}, err
	}
	_, err = io.Copy(part, upload.Content)
	if err != nil {
		return request{}, err
	}
	err = writer.Close()
	if err != nil {
		return request{}, err
	}
	params := url.Values{
		"dataType":  {strconv.Itoa(int(upload.DataType))},
		"mediaType": {strconv.Itoa(int(upload.MediaType))},
	}
	if upload.FileType != "" {
		params.Set("fileType", string(upload.FileType))
	}
	if upload.FormatID != 0 {
		params.Set("formatID", strconv.Itoa(upload.FormatID))
	}
	if upload.Encoding != "" {
		params.Set("encoding", string(upload.Encoding))
	}
	return request{method: http.MethodPost, path: "file_import", params: params, body: body.Bytes(), contentType: writer.FormDataContentType()}, nil
}

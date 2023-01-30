package models

// API endpoints
const (
	ArticleListPath     = "article_list"
	ArticleStatusPath   = "article_status"
	ArticlePath         = "article"
	AuthenticatePath    = "authenticate"
	FileImportPath      = "file_import"
	FileStatusPath      = "file_status"
	ImportStatusPath    = "import_status"
	OrderItemCancelPath = "order_item_cancel"
	OrderCancelPath     = "order_cancel"
	OrderMessagePath    = "order_message"
	OrderStatusPath     = "order_status"
	OrderPath           = "order"
)

// Request models
type ArticleListRequest struct {
	ReturnType string `json:"field"`
	ShowPrice  bool   `json:"showPrice"`
	ShowStock  bool   `json:"showStock"`
	MediaType  byte   `json:"mediaType,omitempty"`
}

type ArticleStatusRequest struct {
	OrderNo string `json:"orderNo"`
}

type ArticleRequest struct {
	OrderNo string `json:"orderNo"`
}

type AuthenticateRequest struct {
	APIKey string `json:"apiKey"`
}

type FileImportRequest struct {
	File      []byte `json:"file"`
	FileType  string `json:"fileType,omitempty"`
	DataType  byte   `json:"dataType,omitempty"`
	MediaType byte   `json:"mediaType,omitempty"`
	FormatID  int    `json:"formatId,omitempty"`
	Encoding  string `json:"encoding,omitempty"`
}

type FileStatusRequest struct {
	Filename string `json:"filename"`
}

type OrderCancelRequest struct {
	OrderID string `json:"orderId"`
}

type OrderItemCancelRequest struct {
	OrderItemID string `json:"orderItemId"`
	MediaType   byte   `json:"mediaType"`
}

type OrderMessageRequest struct {
	OrderID        string `json:"orderId"`
	MessageType    string `json:"messageType"`
	AdditionalText string `json:"additionalText,omitempty"`
}

type OrderStatusRequest struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

type OrderRequest struct {
	OrderID  string `json:"orderId"`
	Date     string `json:"date"`
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

// Response models
type OrderResponse struct {
	Status string `json:"status"`
	// TODO: Add Order struct
}

// Implemented from many endpoints
type GlobalResponse struct {
	Status      string `json:"status"`
	ReturnValue string `json:"returnValue,omitempty"`
}

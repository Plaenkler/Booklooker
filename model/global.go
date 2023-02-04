package model

import "time"

const BaseURL = "https://api.booklooker.de/2.0/"

// Endpoints
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

// Encodings
type Encoding string

const (
	UTF8  Encoding = "UTF-8"
	CP437 Encoding = "IBMPC/CR"
	MacOS Encoding = "macintosh"
)

// Data types
type DataType int

const (
	AddChangeDelete DataType = 0
	Replace         DataType = 1
	Delete          DataType = 2
)

// File types
type FileType string

const (
	Article FileType = "article"
	Picture FileType = "pic"
)

// Message types
type MessageType string

const (
	PaymentInformation MessageType = "PAYMENT_INFORMATION"
	PaymentReminder    MessageType = "PAYMENT_REMINDER"
	ShippingNotice     MessageType = "SHIPPING_NOTICE"
)

// Media types
type MediaType string

const (
	Books      MediaType = "0"
	Movies     MediaType = "1"
	Music      MediaType = "2"
	AudioBooks MediaType = "3"
	Games      MediaType = "4"
)

// Tokens have a lifetime of 10 minutes
type Token struct {
	Value  string `json:"token"`
	Expiry time.Time
}

// Implemented by most endpoints
type GlobalResponse struct {
	Status      string `json:"status"`
	ReturnValue string `json:"returnValue,omitempty"`
}

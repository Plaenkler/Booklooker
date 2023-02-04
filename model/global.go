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

// Media types
type MediaType int

const (
	Books MediaType = iota
	Movies
	Music
	AudioBooks
	Games
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

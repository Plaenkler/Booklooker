package models

const ArticlePath = "article"

type ArticleRequest struct {
	OrderNo string `json:"orderNo"`
}

type ArticleResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

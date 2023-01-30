package models

const ArticleStatusPath = "article_status"

type ArticleStatusRequest struct {
	OrderNo string `json:"orderNo"`
}

type ArticleStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

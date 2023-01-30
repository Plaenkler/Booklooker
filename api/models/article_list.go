package models

const ArticleListPath = "article_list"

type ArticleListRequest struct {
	ReturnType string `json:"field"`
	ShowPrice  bool   `json:"showPrice"`
	ShowStock  bool   `json:"showStock"`
	MediaType  byte   `json:"mediaType,omitempty"`
}

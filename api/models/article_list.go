package models

const ArticleListPath = "article_list"

type ArticleListRequest struct {
	Field     string `json:"field"`
	ShowPrice bool   `json:"showPrice"`
	ShowStock bool   `json:"showStock"`
	MediaType byte   `json:"mediaType,omitempty"`
}

type ArticleListResponse struct {
	Status      string   `json:"status"`
	Message     string   `json:"message,omitempty"`
	ArticleList []string `json:"value,omitempty"`
}

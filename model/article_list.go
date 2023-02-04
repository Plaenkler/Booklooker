package model

type ArticleListRequest struct {
	ReturnType string `json:"field"`
	ShowPrice  bool   `json:"showPrice"`
	ShowStock  bool   `json:"showStock"`
	MediaType  string `json:"mediaType,omitempty"`
}

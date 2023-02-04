package model

type ArticleListRequest struct {
	ReturnType string
	ShowPrice  bool
	ShowStock  bool
	MediaType  MediaType
}

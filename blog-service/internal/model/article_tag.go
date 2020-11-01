package model

type ArticleTag struct {
	*Model
	TagID     uint64 `json:"tag_id"`
	articleID uint64 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

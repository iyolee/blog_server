package model

import "blog_server/pkg/app"

// Article 文章model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

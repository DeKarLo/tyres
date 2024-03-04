package models

import (
	"time"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	Img       string
	CreatedAt time.Time
	UpdatetAt time.Time
}

func (p *Post) NewPost(title, content, img string, createdAt, updatedAt time.Time) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		Img:       img,
		CreatedAt: createdAt,
		UpdatetAt: updatedAt,
	}
}

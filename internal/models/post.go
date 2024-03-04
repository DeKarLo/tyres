package models

import (
	"time"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	Img       string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) NewPost(title, content, img string, price int, createdAt, updatedAt time.Time) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		Img:       img,
		Price:     price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

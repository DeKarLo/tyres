package models

type Post struct {
	ID      int
	Title   string
	Content string
	img     string
}

func (p *Post) NewPost(title, content, img string) *Post {
	return &Post{
		Title:   title,
		Content: content,
		img:     img,
	}
}

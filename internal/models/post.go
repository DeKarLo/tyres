package models

type Post struct {
	ID      int
	Title   string
	Content string
	Img     string
	Price   int
}

func (p *Post) NewPost(title, content, img string, price int) *Post {
	return &Post{
		Title:   title,
		Content: content,
		Img:     img,
		Price:   price,
	}
}

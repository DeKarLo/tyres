package models

type Post struct {
	ID      int
	UserID  int
	Title   string
	Content string
	Img     string
	Price   int
}

func (p *Post) NewPost(title, content, img string, userID, price int) *Post {
	return &Post{
		UserID:  userID,
		Title:   title,
		Content: content,
		Img:     img,
		Price:   price,
	}
}

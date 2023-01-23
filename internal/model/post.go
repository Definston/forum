package model

type Post struct {
	Id       int
	UserId   int
	UserNick string
	ParentId int
	Tag      map[string]bool
	Content  string
	Likes    int
	Dislikes int
	Link     string
	Resiever *int
}

package models

type Reaction struct {
	UserId       int
	PostId       int
	Islike       int
	CommentId    int
	CountLike    int
	CountDislike int
}

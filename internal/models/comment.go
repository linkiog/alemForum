package models

type Comment struct {
	IdComment int
	IdPost    int
	IdAuth    int
	Author    string
	Content   string
	Like      int
	Dislike   int
	Date      string
}

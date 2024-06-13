package models

type Post struct {
	IdPost     int
	IdAuth     int
	Author     string
	Title      string
	Content    string
	Category   []string
	Like       int
	Dislike    int
	CreateDate string
}

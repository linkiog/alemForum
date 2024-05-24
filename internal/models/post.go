package models

import "time"

type Post struct {
	IdPost     int
	IdAuth     int
	Author     string
	Title      string
	Content    string
	Category   []string
	Like       int
	Dislike    int
	CreateDate time.Time
}

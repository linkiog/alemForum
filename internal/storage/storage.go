package storage

import "database/sql"

type Storage struct {
	Auth
	User
	Post
	Comment
	Reaction
}

func InitStorage(db *sql.DB) *Storage {
	return &Storage{
		Auth:     InitAuth(db),
		User:     InitUserStorage(db),
		Post:     InitPost(db),
		Comment:  InitComment(db),
		Reaction: InitReaction(db),
	}

}

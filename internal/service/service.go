package service

import "forum/internal/storage"

type Service struct {
	Auth
	User
	PostSer
	Comment
	Reaction
}

func InitService(storage *storage.Storage) *Service {
	return &Service{
		Auth:     InitAuthService(storage),
		User:     InitUserService(storage),
		PostSer:  InitPostService(storage.Post),
		Comment:  InitComment(storage.Comment),
		Reaction: InitReactionService(storage.Reaction),
	}
}

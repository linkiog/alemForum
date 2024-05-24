package service

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type Comment interface {
	CreateComment(comment, author string, idAuth, idPost int, time string) error
	GetAllComment(idPost int) ([]models.Comment, error)
	GetOneCommentByIdComment(idComment int) (models.Comment, error)
}
type CommentService struct {
	db storage.Comment
}

func InitComment(db storage.Comment) Comment {
	return &CommentService{
		db: db,
	}

}

func (c *CommentService) CreateComment(comment, author string, idAuth, idPost int, time string) error {
	return c.db.CreateComment(comment, author, idAuth, idPost, time)

}
func (c *CommentService) GetAllComment(idPost int) ([]models.Comment, error) {
	return c.db.GetAllComment(idPost)

}

func (c *CommentService) GetOneCommentByIdComment(idComment int) (models.Comment, error) {
	return c.db.GetOneCommentByIdComment(idComment)

}

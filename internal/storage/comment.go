package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

type Comment interface {
	CreateComment(content, author string, idAuth, idPost int, time string) error
	GetAllComment(postId int) ([]models.Comment, error)
	GetOneCommentByIdComment(idComment int) (models.Comment, error)
}

type CommentStorage struct {
	db *sql.DB
}

func InitComment(db *sql.DB) Comment {
	return &CommentStorage{
		db: db,
	}

}

func (c *CommentStorage) CreateComment(content, author string, idAuth, idPost int, time string) error {
	query := `INSERT INTO comment(idPost,idAuth,author,content,like,dislike,createDate) VALUES($1,$2,$3,$4,$5,$6,$7)`
	var comment models.Comment
	_, err := c.db.Exec(query, idPost, idAuth, author, content, comment.Like, comment.Dislike, time)
	if err != nil {
		return fmt.Errorf("Create Comment %w" + err.Error())
	}
	return nil

}

func (c *CommentStorage) GetAllComment(postId int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT idComment,idAuth,author,content,like,dislike,createDate FROM comment WHERE idPost=$1;`
	rows, err := c.db.Query(query, postId)
	if err != nil {
		return []models.Comment{}, fmt.Errorf("storage GetAllComment func:%w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.IdComment, &comment.IdAuth, &comment.Author, &comment.Content, &comment.Like, &comment.Dislike, &comment.Date); err != nil {
			return []models.Comment{}, fmt.Errorf("scan GetAllComment%w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil

}
func (c *CommentStorage) GetOneCommentByIdComment(idComment int) (models.Comment, error) {
	var comment models.Comment
	query := `SELECT idComment, idAuth, author, content, like, dislike, createDate FROM comment WHERE idComment = $1;`
	rows, err := c.db.Query(query, idComment)
	if err != nil {
		return models.Comment{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&comment.IdComment, &comment.IdAuth, &comment.Author, &comment.Content, &comment.Like, &comment.Dislike, &comment.Date); err != nil {
			return models.Comment{}, fmt.Errorf("scan GetOneCommentById: %w", err)
		}
	}
	if err := rows.Err(); err != nil {
		return models.Comment{}, fmt.Errorf("rows.Err in GetOneCommentById: %w", err)
	}

	return comment, nil
}

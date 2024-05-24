package service

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/storage"
	"strings"
)

type PostSer interface {
	CreatePost(post models.Post) error
	GetCategories() ([]models.Category, error)
	GetAllPost() ([]models.Post, error)
	GetOnePost(id int) (models.Post, error)
	GetMyPosts(id int) ([]models.Post, error)
	GetMyLikedPost(id int) ([]models.Post, error)
	Category() ([]models.Category, error)
	GetPostsByCategory(category string) ([]models.Post, error)
	EditPost(post models.Post) error
}

type PostService struct {
	db storage.Post
}

func InitPostService(db storage.Post) PostSer {
	return &PostService{
		db: db,
	}

}
func (p *PostService) CreatePost(post models.Post) error {
	for x := range post.Category {
		post.Category[x] = strings.TrimSpace(post.Category[x])
		if len(post.Category[x]) == 0 {
			return fmt.Errorf("Empty categary")
		}
		if len(post.Category[x]) == 0 || len(post.Category[x]) > 6 || len(post.Category) == 0 {
			return fmt.Errorf("INVALID CATEGORY, category should be shorter than 35 symbols and not empty")

		}
	}
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	if len(post.Content) > 500 || len(post.Content) == 0 {
		return fmt.Errorf("content should be shorter than 500 symbols and not empty")

	}
	if len(post.Title) == 0 || len(post.Title) > 35 {
		return fmt.Errorf("INVALID TITLE, title should be shorter than 35 symbols and not empty")

	}
	if len(post.Category) == 0 {
		return fmt.Errorf("INVALID CATEGORY, please select existing categories ")
	}
	return p.db.CreatePost(post)

}
func (p *PostService) GetCategories() ([]models.Category, error) {
	return p.db.Category()

}
func (p *PostService) GetAllPost() ([]models.Post, error) {
	return p.db.GetAllPosts()
}
func (p *PostService) GetOnePost(id int) (models.Post, error) {
	return p.db.GetOnePost(id)

}
func (p *PostService) GetMyPosts(id int) ([]models.Post, error) {
	return p.db.GetMyPosts(id)

}
func (p *PostService) GetMyLikedPost(id int) ([]models.Post, error) {
	return p.db.GetMyLikedPost(id)

}
func (p *PostService) Category() ([]models.Category, error) {
	return p.db.Category()

}
func (p *PostService) GetPostsByCategory(category string) ([]models.Post, error) {
	return p.db.GetPostsByCategory(category)

}
func (p *PostService) EditPost(post models.Post) error {
	return p.db.EditPost(post)

}

package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"strings"
)

type Post interface {
	CreatePost(post models.Post) error
	Category() ([]models.Category, error)
	GetAllPosts() ([]models.Post, error)
	GetOnePost(id int) (models.Post, error)
	GetMyPosts(id int) ([]models.Post, error)
	GetMyLikedPost(id int) ([]models.Post, error)
	GetPostsByCategory(category string) ([]models.Post, error)
	EditPost(post models.Post) error
}

type PostStr struct {
	db *sql.DB
}

func InitPost(db *sql.DB) Post {
	return &PostStr{
		db: db,
	}

}
func (p *PostStr) CreatePost(post models.Post) error {
	query := `INSERT INTO post (idAuth, author, title, content, category, createDate) VALUES ($1,$2,$3,$4,$5,$6)`
	var categoriesStr string
	if len(post.Category) == 1 {
		categoriesStr = post.Category[0]
	} else {
		categoriesStr = strings.Join(post.Category, ", ")
	}
	_, err := p.db.Exec(query, post.IdAuth, post.Author, post.Title, post.Content, categoriesStr, post.CreateDate)
	if err != nil {
		return err
	}
	return nil

}

func (p *PostStr) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT * FROM post ORDER BY createDate DESC`
	row, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("storage get all post (query) %w", err)
	}
	defer row.Close()
	for row.Next() {
		var post models.Post
		var categoriesStr string
		if err := row.Scan(&post.IdPost, &post.IdAuth, &post.Author, &post.Title, &post.Content, &categoriesStr, &post.Like, &post.Dislike, &post.CreateDate); err != nil {
			return nil, fmt.Errorf("storage get all post,scan %w", err)
		}
		post.Category = strings.Split(categoriesStr, ", ")
		posts = append(posts, post)

	}
	return posts, nil

}
func (p *PostStr) GetOnePost(id int) (models.Post, error) {
	query := `SELECT idPost,idAuth,author,title,content,category,like,dislike,createDate FROM post WHERE idPost=$1;`
	var post models.Post
	var categoriesStr string
	row := p.db.QueryRow(query, id)
	if err := row.Scan(&post.IdPost, &post.IdAuth, &post.Author, &post.Title, &post.Content, &categoriesStr, &post.Like, &post.Dislike, &post.CreateDate); err != nil {
		return models.Post{}, fmt.Errorf("GetOnePost with id %w", err)
	}
	post.Category = strings.Split(categoriesStr, ", ")
	return post, nil

}
func (p *PostStr) GetMyPosts(id int) ([]models.Post, error) {
	query := `SELECT * FROM post WHERE idAuth=$1 ORDER BY idPost DESC;`
	var posts []models.Post
	var categoriesStr string
	row, err := p.db.Query(query, id)
	if err != nil {
		return []models.Post{}, fmt.Errorf("problem with GetMyPost%w", err)
	}
	defer row.Close()
	for row.Next() {
		var post models.Post
		if err := row.Scan(&post.IdPost, &post.IdAuth, &post.Author, &post.Title, &post.Content, &categoriesStr, &post.Like, &post.Dislike, &post.CreateDate); err != nil {
			return []models.Post{}, fmt.Errorf("Problm with GetMyPosts (row.Next()): %w", err)

		}
		post.Category = strings.Split(categoriesStr, ", ")
		posts = append(posts, post)

	}
	return posts, nil

}
func (p *PostStr) GetMyLikedPost(id int) ([]models.Post, error) {
	query := `SELECT post.* FROM post
	JOIN reaction
	ON reaction.postId=post.idPost
	WHERE reaction.userID=$1 AND reaction.reaction=1`
	var posts []models.Post
	var categoriesStr string
	row, err := p.db.Query(query, id)
	if err != nil {
		return []models.Post{}, fmt.Errorf("Problem with GetMyLikedPost: %w", err)

	}
	for row.Next() {
		var post models.Post
		if err := row.Scan(&post.IdPost, &post.IdAuth, &post.Author, &post.Title, &post.Content, &categoriesStr, &post.Like, &post.Dislike, &post.CreateDate); err != nil {
			return []models.Post{}, fmt.Errorf("Problem with GetMyLikedPost: (row.Next()), %w", err)
		}
		post.Category = strings.Split(categoriesStr, ", ")
		posts = append(posts, post)

	}
	return posts, nil

}

func (p *PostStr) Category() ([]models.Category, error) {
	query := `SELECT name FROM categories;`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {

		return nil, err
	}
	return categories, nil

}
func (p *PostStr) GetPostsByCategory(category string) ([]models.Post, error) {
	query := `SELECT * FROM  post WHERE category LIKE '%' || $1 || '%';`
	var posts []models.Post
	var categoriesStr string
	row, err := p.db.Query(query, category)
	if err != nil {
		return []models.Post{}, fmt.Errorf("Problem with GetPostsByCategory: %w", err)
	}
	for row.Next() {
		var post models.Post
		if err := row.Scan(&post.IdPost, &post.IdAuth, &post.Author, &post.Title, &post.Content, &categoriesStr, &post.Like, &post.Dislike, &post.CreateDate); err != nil {
			return []models.Post{}, fmt.Errorf("Problem with GetMyLikedPost: (row.Next()), %w", err)
		}
		post.Category = strings.Split(categoriesStr, ", ")
		posts = append(posts, post)

	}
	return posts, nil

}

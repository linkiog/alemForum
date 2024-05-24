package storage

import (
	"fmt"
	"forum/internal/models"
	"strings"
)

func (p *PostStr) EditPost(post models.Post) error {
	query := `UPDATE post SET title=$1, content=$2, category=$3, createDate=$4 WHERE idPost=$5;`
	var categoriesStr string
	if len(post.Category) == 1 {
		categoriesStr = post.Category[0]
	} else {
		categoriesStr = strings.Join(post.Category, ", ")
	}
	if _, err := p.db.Exec(query, post.Title, post.Content, categoriesStr, post.CreateDate, post.IdPost); err != nil {
		return fmt.Errorf("Problem withs editPost sql %w", err)

	}
	return nil

}

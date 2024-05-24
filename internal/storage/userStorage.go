package storage

import (
	"database/sql"
	"forum/internal/models"
)

type User interface {
	CheckUserFromdb(email string) (bool, error)
	GetUserByToken(token string) (models.User, error)
}

type UserStorage struct {
	db *sql.DB
}

func InitUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) CheckUserFromdb(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = ? AS UE_exists;"

	row := u.db.QueryRow(query, email)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (u *UserStorage) GetUserByToken(token string) (models.User, error) {
	query := `SELECT id,email,username,token_duration FROM user WHERE token=$1;`
	row := u.db.QueryRow(query, token)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Token_duration); err != nil {
		return models.User{}, err

	}
	return user, nil
}

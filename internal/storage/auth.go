package storage

import (
	"database/sql"
	"forum/internal/models"
	"time"
)

type Auth interface {
	CreateUser(user models.User) error
	CheckUserFromdb(email string) (bool, error)
	GetUserByEmail(email string) (string, error)
	SaveToken(string, time.Time, string) error
	DeleteToken(token string) error
}

type AuthStorage struct {
	db *sql.DB
}

func InitAuth(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (a *AuthStorage) CreateUser(user models.User) error {
	query := `INSERT INTO user(email,username,password) VALUES ($1,$2,$3);`
	_, err := a.db.Exec(query, user.Email, user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthStorage) CheckUserFromdb(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user WHERE email = ?) AS UE_exists;"

	row := a.db.QueryRow(query, email)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (a *AuthStorage) GetUserByEmail(email string) (string, error) {
	query := `SELECT password from user where email=$1;`
	row := a.db.QueryRow(query, email)
	var password string
	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil

}
func (a *AuthStorage) SaveToken(token string, expired time.Time, email string) error {
	query := `UPDATE user SET token=$1, token_duration=$2 WHERE email=$3;`
	if _, err := a.db.Exec(query, token, expired, email); err != nil {
		return err
	}
	return nil

}
func (a *AuthStorage) DeleteToken(token string) error {
	query := `UPDATE user SET token = NULL,token_duration=NULL WHERE token=$1`
	if _, err := a.db.Exec(query, token); err != nil {
		return err
	}
	return nil
}

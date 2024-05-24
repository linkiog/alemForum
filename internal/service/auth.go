package service

import (
	"errors"
	"forum/internal/models"
	"forum/internal/storage"
	"regexp"
	"time"
	"unicode"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	CreateUser(models.User) error
	CheckUserFormDb(models.User) (string, time.Time, error)
	DeleteToken(token string) error
}

type AuthService struct {
	storage *storage.Storage
}

func InitAuthService(storage *storage.Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (a *AuthService) CreateUser(user models.User) error {
	if err := IsValidUser(user); err != nil {
		return err
	}
	check, err := a.storage.Auth.CheckUserFromdb(user.Email)
	if err != nil {
		return err
	}
	if check {
		return errors.New("Such an email exists")
	}
	user.Password, err = generateHash(user.Password)
	if err != nil {
		return err
	}
	return a.storage.CreateUser(user)
}

func IsValidUser(user models.User) error {
	for _, char := range user.Name {
		if char <= 32 || char >= 127 {
			return models.ErrInvalidUserName
		}
	}
	validEmail, err := regexp.MatchString(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, user.Email)
	if err != nil {
		return err
	}
	if !validEmail {
		return models.ErrInvalidEmail
	}
	if len(user.Name) < 6 || len(user.Name) >= 30 {
		return models.ErrInvalidUserName
	}
	if !IsValidPass(user.Password) {
		return models.ErrShortPassword
	}
	return nil
}

func IsValidPass(password string) bool {
	var (
		length     = false
		upp        = false
		lower      = false
		number     = false
		specialSym = false
	)
	if len(password) >= 7 || len(password) <= 20 {
		length = true
	}
	for i := range password {
		switch {
		case unicode.IsUpper(rune(password[i])):
			upp = true
		case unicode.IsLower(rune(password[i])):
			lower = true
		case unicode.IsNumber(rune(password[i])):
			number = true
		case unicode.IsPunct(rune(password[i])) || unicode.IsSymbol(rune(password[i])):
			specialSym = true

		}
	}
	return length && upp && lower && number && specialSym
}

func generateHash(passw string) (string, error) {
	hashPassw, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.DefaultCost)
	return string(hashPassw), err
}

func (a *AuthService) CheckUserFormDb(user models.User) (string, time.Time, error) {
	password, err := a.storage.GetUserByEmail(user.Email)
	if err != nil {
		return "", time.Time{}, errors.New("password or email address is incorrect")
	}
	if err := comparePassw(password, user.Password); err != nil {
		return "", time.Time{}, err
	}
	token := uuid.NewGen()
	d, err := token.NewV4()
	if err != nil {
		return "", time.Time{}, err
	}
	expired := time.Now().Add(time.Hour * 12)
	if err := a.storage.SaveToken(d.String(), expired, user.Email); err != nil {
		return "", time.Time{}, err
	}
	return d.String(), expired, nil

}
func comparePassw(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return errors.New("password or email address is incorrect")
	}

	return nil
}

func (a *AuthService) DeleteToken(token string) error {
	return a.storage.Auth.DeleteToken(token)
}

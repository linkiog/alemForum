package models

import "errors"

var (
	ErrUserNotFound      = errors.New("User does not exist or password incorrect")
	ErrInvalidUserName   = errors.New("Invalid username - your username should consist at least 6 characters")
	ErrInvalidEmail      = errors.New("Invalid email")
	ErrPasswordDontMatch = errors.New("Password didn't match")
	ErrShortPassword     = errors.New("Incorrect password - your password should be a minimum of 8 characters and consist of at least:1 lower case letter, 1 upper case letter, 1 number, 1 special symbol")
	Err                  = errors.New("You have successfully registered")
)

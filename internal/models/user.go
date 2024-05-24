package models

import "time"

type User struct {
	ID             int
	Email          string
	Name           string
	Password       string
	RPassw         string
	Error          error
	Status         string
	Token_duration time.Time
	IsAuth         bool
}

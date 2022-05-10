package models

import "time"

type Authentication struct {
	Email    string
	Password string
}

type Token struct {
	Role  int
	Email string
	Token string
}

type UserAcccessInfo struct {
	Token     string
	ExpDate   time.Time
	CreatedAt time.Time
}

package model

import "time"

type User struct {
	ID       string `json:"id"`
	CreateAt time.Time
}

func NewUser() *User {
	return &User{
		CreateAt: time.Now(),
	}
}

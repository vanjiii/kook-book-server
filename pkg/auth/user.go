package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID

	Email  string
	Name   string
	Avatar string

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

func (u User) TableName() string {
	return "users"
}

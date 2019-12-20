package pkg

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uuid.UUID

	Email    string
	Name     string
	Avatar   string
	Password []byte

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

func (u User) TableName() string {
	return "users"
}

func (u *User) SetPassword(password string) error {
	if !validatePassword(password) {
		return errors.ErrBadPass
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hash
}

func validatePassword(password string) bool {
	// Some complex logic here
	return len(password) >= 4
}

package auth

import (
	"context"
	"vanjiii/kook-book-server/pkg"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) CreateUser(ctx context.Context, email, password string) error {
	user := pkg.User{Email: email}
	user.SetPassword(password)

	return h.db.Create(&user).Error
}

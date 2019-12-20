package auth

import (
	"log"

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

func (h *Handler) CreateUser() error {
	log.Println("I am creating a user")
	return nil
}

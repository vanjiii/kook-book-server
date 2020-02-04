package pkg

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID uuid.UUID

	Photo       string
	Ingredients []Ingredient
	Description string
	PrepTime    time.Time

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

package recipe

import (
	"time"
)

type Recipe struct {
	ID uint

	Photo       string
	Ingredients []Ingredient
	Description string
	PrepTime    time.Duration

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

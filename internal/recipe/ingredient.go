package recipe

import (
	"time"
)

type Ingredient struct {
	ID uint

	Name     string
	Quantity string

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

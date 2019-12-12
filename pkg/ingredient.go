package pkg

import (
	"time"

	"github.com/google/uuid"
)

type Ingredient struct {
	ID uuid.UUID

	Name     string
	Quantity string

	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

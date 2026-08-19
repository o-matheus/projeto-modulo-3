package domain

import (
	"time"

	"github.com/google/uuid"
)

type Produto struct {
	ID        uuid.UUID
	Name      string
	Price     int64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

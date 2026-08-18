package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Cliente struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCliente(name, email, passwordHash string) (Cliente, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return Cliente{}, ErrNomeClienteRequired
	}

	if email == "" {
		return Cliente{}, ErrEmailClienteRequired
	}

	if passwordHash == "" {
		return Cliente{}, ErrPasswordClienteRequired
	}

	id, err := uuid.NewV7()
	if err != nil {
		return Cliente{}, err
	}

	now := time.Now()

	return Cliente{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

package application

import (
	"context"
	"projeto-modulo-3/internal/domain"

	"github.com/google/uuid"
)

type ClienteRepository interface {
	Create(ctx context.Context, cliente domain.Cliente) error
	List(ctx context.Context) ([]domain.Cliente, error)
	GetById(ctx context.Context, id uuid.UUID) (domain.Cliente, error)
	Update(ctx context.Context, cliente domain.Cliente) error
	Delete(ctx context.Context, id uuid.UUID) error
}

package application

import (
	"context"
	"projeto-modulo-3/internal/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ClienteService struct {
	repo ClienteRepository
}

func NewClienteService(repo ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) Create(ctx context.Context, name, email, password string) (domain.Cliente, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Cliente{}, err
	}

	cliente, err := domain.NewCliente(name, email, string(hash))
	if err != nil {
		return domain.Cliente{}, err
	}

	if err := s.repo.Create(ctx, cliente); err != nil {
		return domain.Cliente{}, err
	}

	return cliente, nil
}

func (s *ClienteService) List(ctx context.Context) ([]domain.Cliente, error) {
	return s.repo.List(ctx)
}

func (s *ClienteService) GetById(ctx context.Context, id uuid.UUID) (domain.Cliente, error) {
	return s.repo.GetById(ctx, id)
}

func (s *ClienteService) Update(ctx context.Context, id uuid.UUID, name, email string) (domain.Cliente, error) {
	cliente, erro := s.repo.GetById(ctx, id)
	if erro != nil {
		return domain.Cliente{}, erro
	}

	if err := cliente.Update(name, email); err != nil {
		return domain.Cliente{}, err
	}

	if err := s.repo.Update(ctx, cliente); err != nil {
		return domain.Cliente{}, err
	}

	return cliente, nil
}

func (s *ClienteService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

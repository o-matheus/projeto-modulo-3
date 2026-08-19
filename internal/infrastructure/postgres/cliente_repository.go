package postgres

import (
	"context"
	"errors"
	"fmt"
	"projeto-modulo-3/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClienteRepository struct {
	db *pgxpool.Pool
}

func NewClienteRepository(db *pgxpool.Pool) *ClienteRepository {
	return &ClienteRepository{db: db}
}

func (r *ClienteRepository) Create(ctx context.Context, cliente domain.Cliente) error {
	_, err := r.db.Exec(ctx, `INSERT INTO clientes (id, name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`, cliente.ID, cliente.Name, cliente.Email, cliente.PasswordHash, cliente.CreatedAt, cliente.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrEmailAlreadyExists
		}
		return fmt.Errorf("criar cliente: %w", err)
	}
	return nil
}

func (r *ClienteRepository) List(ctx context.Context) ([]domain.Cliente, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email, password_hash, created_at, updated_at FROM clientes ORDER BY created_at`)
	if err != nil {
		return nil, fmt.Errorf("listar clientes: %w", err)
	}
	defer rows.Close()

	var clientes []domain.Cliente
	for rows.Next() {
		var cliente domain.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Name, &cliente.Email, &cliente.PasswordHash, &cliente.CreatedAt, &cliente.UpdatedAt); err != nil {
			return nil, fmt.Errorf("ler clientes: %w", err)
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("percorrer clientes: %w", err)
	}

	return clientes, nil
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Pedido struct {
	ID        uuid.UUID
	ClienteID uuid.UUID
	Status    PedidoStatus
	Items     []ItemPedido
	CreatedAt time.Time
}

type ItemPedido struct {
	ProdutoID  uuid.UUID
	Quantity   int
	PriceCents int64
}

type PedidoStatus string

const (
	StatusPending  PedidoStatus = "PENDING"
	StatusPaid     PedidoStatus = "PAID"
	StatusCanceled PedidoStatus = "CANCELED"
)

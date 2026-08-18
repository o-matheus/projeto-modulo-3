package domain

import "errors"

var (
	ErrNomeClienteRequired     = errors.New("nome do cliente não pode ser vazio")
	ErrEmailClienteRequired    = errors.New("email do cliente não pode ser vazio")
	ErrPasswordClienteRequired = errors.New("senha do cliente não pode ser vazia")
	ErrClienteNotFound         = errors.New("cliente não encontrado")
	ErrProdutoNotFound         = errors.New("produto não encontrado")
	ErrPedidoNotFound          = errors.New("pedido não encontrado")
	ErrEmailAlreadyExists      = errors.New("email já cadastrado")
	ErrEstoqueInsuficiente     = errors.New("estoque insuficiente")
	ErrPedidoSemItens          = errors.New("pedido deve conter pelo menos um item")
	ErrQuantidadeInvalida      = errors.New("quantidade deve ser maior que zero")
	ErrStatusInvalido          = errors.New("status do pedido não permite essa operação")
	ErrClienteObrigatorio      = errors.New("cliente do pedido é obrigatório")
)

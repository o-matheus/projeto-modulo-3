CREATE TABLE clientes
(
    id            UUID PRIMARY KEY      DEFAULT uuidv7(),
    name          VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMPZ   NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPZ   NOT NULL DEFAULT now(),
)

CREATE TABLE produtos
(
    id         UUID PRIMARY KEY        DEFAULT uuidv7(),
    name       VARCHAR(255)   NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    stock      INT            NOT NULL CHECK (stock >= 0),
    created_at TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ     NOT NULL DEFAULT now()
)

CREATE TABLE pedidos
(
    id         UUID PRIMARY KEY     DEFAULT uuidv7(),
    cliente_id UUID        NOT NULL REFERENCES clientes (id),
    status     VARCHAR(20) NOT NULL DEFAULT 'PENDING'
        CHECK (status IN ('PENDING', 'PAID', 'CANCELED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)

CREATE TABLE itens_pedido
(
    id            UUID PRIMARY KEY        DEFAULT uuidv7(),
    pedido_id     UUID           NOT NULL REFERENCES pedidos (id) ON DELETE CASCADE,
    produto_id    UUID           NOT NULL REFERENCES produtos (id),
    quantity      INT            NOT NULL CHECK (quantity > 0),
    price_at_time DECIMAL(10, 2) NOT NULL,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
)
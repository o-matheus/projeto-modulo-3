GOCACHE ?= /tmp/go-build-cache
TOOLS_DIR ?= $(CURDIR)/.bin
MIGRATE ?= $(TOOLS_DIR)/migrate
MIGRATE_VERSION ?= v4.19.1

.PHONY: help up down install-migrate

help:
	@printf "Comandos:\n"
	@printf " make up - Iniciar serviços\n"
	@printf " make down - Parar serviços\n"

up:
	@docker compose up -d

down:
	@docker compose down

install-migrate:
	@mkdir -p "$(TOOLS_DIR)"
	@GOBIN="$(TOOLS_DIR)" go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)


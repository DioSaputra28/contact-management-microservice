# Load environment variables from .env
include .env
export

.PHONY: migrate-up migrate-down migrate-fresh migrate-status migrate-create help

# Database connection string for MySQL
DB_URL := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

help:
	@echo "Available commands:"
	@echo "  make migrate-up      - Run all pending migrations"
	@echo "  make migrate-down    - Rollback last migration"
	@echo "  make migrate-fresh   - Drop all tables and re-run migrations"
	@echo "  make migrate-status  - Show current migration version"
	@echo "  make migrate-create  - Create new migration file"

migrate-up:
	@echo "Running migrations..."
	migrate -path migrations -database $(DB_URL) up
	@echo "✓ Migrations completed"

migrate-down:
	@echo "Rolling back last migration..."
	migrate -path migrations -database $(DB_URL) down 1
	@echo "✓ Rollback completed"

migrate-fresh:
	@echo "⚠️  WARNING: This will drop all tables!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		migrate -path migrations -database $(DB_URL) drop -f; \
		migrate -path migrations -database $(DB_URL) up; \
		echo "✓ Fresh migration completed"; \
	else \
		echo "Cancelled"; \
	fi

migrate-status:
	@echo "Current migration status:"
	migrate -path migrations -database $(DB_URL) version

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name
	@echo "✓ Migration files created"

.PHONY: help build run test clean migrate rollback status sync setup-db

# Default target
help:
	@echo "Available targets:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build artifacts"
	@echo "  migrate     - Run database migrations"
	@echo "  rollback    - Rollback last migration"
	@echo "  status      - Show migration status"
	@echo "  sync        - Sync asset platforms from CoinGecko API"
	@echo "  setup-db    - Setup local PostgreSQL database"

# Build the application
build:
	@echo "Building application..."
	go build -o bin/cgoffline ./cmd/server

# Run the application
run: build
	@echo "Running application..."
	./bin/cgoffline

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/

# Database operations
migrate: build
	@echo "Running database migrations..."
	./bin/cgoffline -migrate

rollback: build
	@echo "Rolling back last migration..."
	./bin/cgoffline -rollback

status: build
	@echo "Showing migration status..."
	./bin/cgoffline -status

sync: build
	@echo "Syncing asset platforms..."
	./bin/cgoffline -sync

# Database setup
setup-db:
	@echo "Setting up local PostgreSQL database..."
	@echo "Please ensure PostgreSQL is installed and running on your system"
	@echo "Create a database and user with the following commands:"
	@echo ""
	@echo "sudo -u postgres psql"
	@echo "CREATE DATABASE your_database_name;"
	@echo "CREATE USER your_postgres_user WITH PASSWORD 'your_postgres_password';"
	@echo "GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_postgres_user;"
	@echo "\\q"
	@echo ""
	@echo "Then copy env.example to env and update the database credentials:"
	@echo "cp env.example env"
	@echo "Edit env file with your database credentials"

# Development setup
dev-setup: setup-db
	@echo "Running migrations..."
	@make migrate
	@echo "Development setup complete!"
	@echo "Make sure to configure your database credentials in the env file"

# Full sync with fresh data
full-sync: build
	@echo "Performing full synchronization..."
	./bin/cgoffline -sync

# Show help
help:
	@make help

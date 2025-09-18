# cgoffline

A Go application that fetches Asset Platforms data from the CoinGecko API and stores it in a PostgreSQL database using GORM. Built with Clean Architecture principles and modern Go best practices.

## Features

- 🚀 **Clean Architecture**: Modular design with clear separation of concerns
- 🗄️ **GORM Integration**: Modern ORM with automatic migrations
- 🔄 **API Synchronization**: Fetches and syncs asset platforms from CoinGecko API
- 📊 **PostgreSQL Support**: Robust database storage with indexing
- 🔧 **Configuration Management**: Environment-based configuration
- 📝 **Structured Logging**: JSON-formatted logs with different levels
- 🐳 **Docker Support**: Easy PostgreSQL setup with Docker Compose
- ⚡ **Retry Logic**: Resilient API calls with exponential backoff
- 🛠️ **Migration Management**: Database schema versioning with rollback support

## Project Structure

```
cgoffline/
├── cmd/server/           # Application entrypoint
├── internal/
│   ├── domain/          # Domain models and interfaces
│   ├── repository/      # Data access layer
│   ├── service/         # Business logic layer
│   └── handler/         # HTTP handlers (future)
├── pkg/
│   ├── config/          # Configuration management
│   └── logger/          # Logging utilities
├── migrations/          # Database migrations
├── docker-compose.yml   # PostgreSQL and Adminer setup
├── Makefile            # Common tasks
└── env.example         # Environment variables template
```

## Quick Start

### Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later (installed locally)
- Make (optional, for using Makefile commands)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd cgoffline
go mod download
```

### 2. Setup Local PostgreSQL

```bash
# Install PostgreSQL (Ubuntu/Debian)
sudo apt update
sudo apt install postgresql postgresql-contrib

# Start PostgreSQL service
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Create database and user
sudo -u postgres psql
CREATE DATABASE your_database_name;
CREATE USER your_postgres_user WITH PASSWORD 'your_postgres_password';
GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_postgres_user;
\q
```

### 3. Configure Environment

```bash
# Copy environment template
cp env.example env

# Edit the env file with your database credentials
nano env
```

### 4. Run Migrations

```bash
make migrate
```

### 5. Sync Data

```bash
# Sync asset platforms only
make sync-platforms

# Sync coin categories only
make sync-categories

# Sync both platforms and categories
make sync-all
```

### 6. Run the Application

```bash
make run
```

## Usage

### Command Line Options

```bash
# Run migrations and exit
./bin/cgoffline -migrate

# Rollback last migration
./bin/cgoffline -rollback

# Show migration status
./bin/cgoffline -status

# Sync asset platforms and exit
./bin/cgoffline -sync-platforms

# Sync coin categories and exit
./bin/cgoffline -sync-categories

# Sync both platforms and categories and exit
./bin/cgoffline -sync-all

# Run application normally (with initial sync)
./bin/cgoffline
```

### Makefile Commands

```bash
make help           # Show available commands
make build          # Build the application
make run            # Run the application
make test           # Run tests
make clean          # Clean build artifacts
make migrate        # Run database migrations
make rollback       # Rollback last migration
make status         # Show migration status
make sync-platforms # Sync asset platforms
make sync-categories # Sync coin categories
make sync-all       # Sync both platforms and categories
make setup-db       # Setup local PostgreSQL database
make dev-setup      # Complete development setup
```

## Configuration

The application uses environment variables for configuration. Copy `env.example` to `env` and modify as needed:

```bash
cp env.example env
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `your_postgres_user` |
| `DB_PASSWORD` | Database password | `your_postgres_password` |
| `DB_NAME` | Database name | `your_database_name` |
| `DB_SSLMODE` | SSL mode | `disable` |
| `DB_TIMEZONE` | Database timezone | `UTC` |
| `COINGECKO_BASE_URL` | CoinGecko API base URL | `https://api.coingecko.com/api/v3` |
| `API_TIMEOUT` | API timeout | `30s` |
| `API_RETRY_ATTEMPTS` | Retry attempts | `3` |
| `API_RETRY_DELAY` | Retry delay | `1s` |
| `LOG_LEVEL` | Log level | `info` |
| `LOG_FORMAT` | Log format | `json` |

## Database Schema

The application creates two main tables:

### Asset Platforms Table

```sql
CREATE TABLE asset_platforms (
    id VARCHAR(50) PRIMARY KEY,
    chain_identifier BIGINT,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(100),
    native_coin_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes
CREATE INDEX idx_asset_platforms_chain_identifier ON asset_platforms(chain_identifier);
CREATE INDEX idx_asset_platforms_name ON asset_platforms(name);
CREATE INDEX idx_asset_platforms_native_coin_id ON asset_platforms(native_coin_id);
```

### Coin Categories Table

```sql
CREATE TABLE coin_categories (
    id SERIAL PRIMARY KEY,
    coingecko_id VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes
CREATE UNIQUE INDEX idx_coin_categories_coingecko_id ON coin_categories(coingecko_id);
CREATE INDEX idx_coin_categories_name ON coin_categories(name);
```

## API Integration

The application integrates with the CoinGecko API to fetch data from two endpoints:

### Asset Platforms
- **Endpoint**: `https://api.coingecko.com/api/v3/asset_platforms`
- **Method**: GET
- **Response**: Array of asset platform objects
- **Data**: Blockchain platforms with chain identifiers and native coins

### Coin Categories
- **Endpoint**: `https://api.coingecko.com/api/v3/coins/categories/list`
- **Method**: GET
- **Response**: Array of category objects
- **Data**: Coin categories like DeFi, Stablecoins, NFTs, etc.

### Features
- **Rate Limiting**: Built-in retry logic with exponential backoff
- **Health Check**: API connectivity verification
- **Error Handling**: Comprehensive error handling and logging

## Development

### Architecture

The application follows Clean Architecture principles:

- **Domain Layer**: Core business entities and interfaces
- **Repository Layer**: Data access abstraction
- **Service Layer**: Business logic and external API integration
- **Handler Layer**: HTTP request handling (future)

### Adding New Features

1. Define domain models in `internal/domain/`
2. Implement repository interfaces in `internal/repository/`
3. Add business logic in `internal/service/`
4. Create migrations in `migrations/`
5. Update configuration in `pkg/config/`

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/service/...
```

## Monitoring and Observability

- **Structured Logging**: JSON-formatted logs with correlation IDs
- **Error Tracking**: Comprehensive error handling and logging
- **Health Checks**: API and database connectivity monitoring
- **Metrics**: Request timing and success rates (future)

## Local PostgreSQL Setup

### Database Setup

```bash
# Setup database and user
make setup-db

# Or manually create database
sudo -u postgres psql
CREATE DATABASE your_database_name;
CREATE USER your_postgres_user WITH PASSWORD 'your_postgres_password';
GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_postgres_user;
\q
```

### Database Access

- **PostgreSQL**: `localhost:5432`
- **Credentials**: Use your configured database credentials from the `env` file

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Ensure PostgreSQL is running: `sudo systemctl status postgresql`
   - Check connection parameters in the `env` file
   - Verify database and user exist

2. **API Request Failed**
   - Verify internet connectivity
   - Check CoinGecko API status
   - Review retry configuration

3. **Migration Errors**
   - Check database permissions
   - Verify migration status: `make status`
   - Rollback if needed: `make rollback`

### Logs

The application provides detailed logging:

```bash
# View application logs
./bin/cgoffline 2>&1 | jq

# Filter by log level
./bin/cgoffline 2>&1 | jq 'select(.level == "error")'
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run the test suite
6. Submit a pull request

## License

This project is licensed under the MIT License.

## Related Projects

- [CoinGecko MCP Server](https://github.com/nic0xflamel/coingecko-mcp-server)
- [Official CoinGecko MCP Server](https://mcp.api.coingecko.com)
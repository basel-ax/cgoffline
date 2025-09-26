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

# Sync exchanges only
make sync-exchanges

# Sync coins and their market data only
make sync-coins

# Sync coin details and tickers (filtered by volume)
make sync-coins-data

# Sync all data (platforms, categories, exchanges, and coins)
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

# Sync exchanges and exit
./bin/cgoffline -sync-exchanges

# Sync coins and their market data and exit
./bin/cgoffline -sync-coins

# Sync coin details and tickers (filtered by volume) and exit
./bin/cgoffline -sync-coins-data

# Sync all data (platforms, categories, exchanges, and coins) and exit
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
make sync-exchanges  # Sync exchanges
make sync-coins      # Sync coins and their market data
make sync-coins-data # Sync coin details and tickers (filtered by volume)
make sync-all        # Sync all data (platforms, categories, exchanges, and coins)
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
| `COINS_MIN_TOTAL_VOLUME` | Minimum total_volume to include in coins-data sync | `1000000` |
| `LOG_LEVEL` | Log level | `info` |
| `LOG_FORMAT` | Log format | `json` |

## Database Schema

The application creates three main tables:

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

### Exchanges Table

```sql
CREATE TABLE exchanges (
    id SERIAL PRIMARY KEY,
    coingecko_id VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    year_established INTEGER,
    country VARCHAR(100),
    description TEXT,
    url VARCHAR(500),
    image VARCHAR(500),
    has_trading_incentive BOOLEAN,
    trust_score INTEGER,
    trust_score_rank INTEGER,
    trade_volume_24h_btc DOUBLE PRECISION,
    trade_volume_24h_btc_normalized DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes
CREATE UNIQUE INDEX idx_exchanges_coingecko_id ON exchanges(coingecko_id);
CREATE INDEX idx_exchanges_name ON exchanges(name);
CREATE INDEX idx_exchanges_trust_score ON exchanges(trust_score);
CREATE INDEX idx_exchanges_country ON exchanges(country);
```

### Coins Table

```sql
CREATE TABLE coins (
    id SERIAL PRIMARY KEY,
    coingecko_id VARCHAR(100) UNIQUE NOT NULL,
    symbol VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(500),
    current_price DOUBLE PRECISION,
    market_cap DOUBLE PRECISION,
    market_cap_rank INTEGER,
    fully_diluted_valuation DOUBLE PRECISION,
    total_volume DOUBLE PRECISION,
    high_24h DOUBLE PRECISION,
    low_24h DOUBLE PRECISION,
    price_change_24h DOUBLE PRECISION,
    price_change_percentage_24h DOUBLE PRECISION,
    market_cap_change_24h DOUBLE PRECISION,
    market_cap_change_percentage_24h DOUBLE PRECISION,
    circulating_supply DOUBLE PRECISION,
    total_supply DOUBLE PRECISION,
    max_supply DOUBLE PRECISION,
    ath DOUBLE PRECISION,
    ath_change_percentage DOUBLE PRECISION,
    ath_date TIMESTAMP WITH TIME ZONE,
    atl DOUBLE PRECISION,
    atl_change_percentage DOUBLE PRECISION,
    atl_date TIMESTAMP WITH TIME ZONE,
    last_updated TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes
CREATE UNIQUE INDEX idx_coins_coingecko_id ON coins(coingecko_id);
CREATE INDEX idx_coins_symbol ON coins(symbol);
CREATE INDEX idx_coins_name ON coins(name);
CREATE INDEX idx_coins_market_cap_rank ON coins(market_cap_rank);
```

### Coin Market Data Table

```sql
CREATE TABLE coin_market_data (
    id SERIAL PRIMARY KEY,
    coin_id INTEGER NOT NULL REFERENCES coins(id),
    exchange_id INTEGER NOT NULL REFERENCES exchanges(id),
    price DOUBLE PRECISION,
    volume_24h DOUBLE PRECISION,
    volume_percentage DOUBLE PRECISION,
    last_updated TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes
CREATE INDEX idx_coin_market_data_coin_id ON coin_market_data(coin_id);
CREATE INDEX idx_coin_market_data_exchange_id ON coin_market_data(exchange_id);
CREATE UNIQUE INDEX idx_coin_market_data_coin_exchange ON coin_market_data(coin_id, exchange_id);
```

## API Integration

The application integrates with the CoinGecko API to fetch data from five endpoints:

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

### Exchanges
- **Endpoint**: `https://api.coingecko.com/api/v3/exchanges`
- **Method**: GET
- **Response**: Array of exchange objects
- **Data**: Cryptocurrency exchanges with trading volumes, trust scores, and metadata

### Coins
- **Endpoint**: `https://api.coingecko.com/api/v3/coins/markets`
- **Method**: GET
- **Response**: Array of coin objects with market data
- **Data**: Cryptocurrency coins with prices, market caps, volumes, and market rankings

### Coin Market Data
- **Endpoint**: `https://api.coingecko.com/api/v3/coins/{coin_id}/tickers`
- **Method**: GET
- **Response**: Array of ticker objects
- **Data**: Market data for specific coins across different exchanges

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
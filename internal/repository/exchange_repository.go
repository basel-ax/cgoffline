package repository

import (
	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// ExchangeRepository defines the interface for exchange data operations
type ExchangeRepository interface {
	GetAll() ([]domain.Exchange, error)
	Upsert(exchange domain.Exchange) error
	UpsertBatch(exchanges []domain.Exchange) error
}

type exchangeRepository struct {
	db *gorm.DB
}

// NewExchangeRepository creates a new instance of ExchangeRepository
func NewExchangeRepository(db *gorm.DB) ExchangeRepository {
	return &exchangeRepository{db: db}
}

// GetAll retrieves all exchanges from the database
func (r *exchangeRepository) GetAll() ([]domain.Exchange, error) {
	var exchanges []domain.Exchange
	if err := r.db.Find(&exchanges).Error; err != nil {
		return nil, fmt.Errorf("failed to get all exchanges: %w", err)
	}
	return exchanges, nil
}

// Upsert creates a new exchange or updates an existing one
func (r *exchangeRepository) Upsert(exchange domain.Exchange) error {
	// Set CreatedAt and UpdatedAt for new records or update UpdatedAt for existing
	if exchange.CreatedAt.IsZero() {
		exchange.CreatedAt = time.Now()
	}
	exchange.UpdatedAt = time.Now()

	if err := r.db.
		Where(domain.Exchange{CoingeckoID: exchange.CoingeckoID}).
		Assign(exchange).
		FirstOrCreate(&exchange).Error; err != nil {
		return fmt.Errorf("failed to upsert exchange: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple exchanges in a single transaction
func (r *exchangeRepository) UpsertBatch(exchanges []domain.Exchange) error {
	if len(exchanges) == 0 {
		return nil
	}

	// Filter out exchanges with empty coingecko_id
	validExchanges := make([]domain.Exchange, 0, len(exchanges))
	for _, exchange := range exchanges {
		if exchange.CoingeckoID != "" {
			validExchanges = append(validExchanges, exchange)
		} else {
			logger.GetLogger().WithField("name", exchange.Name).Warn("Skipping exchange with empty coingecko_id")
		}
	}

	if len(validExchanges) == 0 {
		logger.GetLogger().Warn("No valid exchanges to upsert")
		return nil
	}

	// Use a transaction for batch upsert
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, exchange := range validExchanges {
			// Use ON CONFLICT for proper upsert
			if err := tx.Exec(`
				INSERT INTO exchanges (
					coingecko_id, name, year_established, country, description, url, image,
					has_trading_incentive, trust_score, trust_score_rank, trade_volume_24h_btc,
					trade_volume_24h_btc_normalized, created_at, updated_at, deleted_at
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
				ON CONFLICT (coingecko_id)
				DO UPDATE SET
					name = EXCLUDED.name,
					year_established = EXCLUDED.year_established,
					country = EXCLUDED.country,
					description = EXCLUDED.description,
					url = EXCLUDED.url,
					image = EXCLUDED.image,
					has_trading_incentive = EXCLUDED.has_trading_incentive,
					trust_score = EXCLUDED.trust_score,
					trust_score_rank = EXCLUDED.trust_score_rank,
					trade_volume_24h_btc = EXCLUDED.trade_volume_24h_btc,
					trade_volume_24h_btc_normalized = EXCLUDED.trade_volume_24h_btc_normalized,
					updated_at = EXCLUDED.updated_at,
					deleted_at = EXCLUDED.deleted_at
			`,
				exchange.CoingeckoID, exchange.Name, exchange.YearEstablished, exchange.Country,
				exchange.Description, exchange.URL, exchange.Image, exchange.HasTradingIncentive,
				exchange.TrustScore, exchange.TrustScoreRank, exchange.TradeVolume24hBTC,
				exchange.TradeVolume24hBTCNormalized, exchange.CreatedAt, exchange.UpdatedAt, exchange.DeletedAt).Error; err != nil {
				logger.GetLogger().WithError(err).WithField("exchange_id", exchange.CoingeckoID).Error("Failed to upsert exchange in batch")
				return fmt.Errorf("failed to upsert exchange %s: %w", exchange.CoingeckoID, err)
			}
		}
		logger.GetLogger().WithField("count", len(validExchanges)).Info("Successfully upserted exchanges batch")
		return nil
	})
}

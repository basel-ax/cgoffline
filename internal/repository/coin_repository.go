package repository

import (
	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// CoinRepository defines the interface for coin data operations
type CoinRepository interface {
	GetAll() ([]domain.Coin, error)
	GetByCoingeckoID(coingeckoID string) (*domain.Coin, error)
	Upsert(coin domain.Coin) error
	UpsertBatch(coins []domain.Coin) error
}

type coinRepository struct {
	db *gorm.DB
}

// NewCoinRepository creates a new instance of CoinRepository
func NewCoinRepository(db *gorm.DB) CoinRepository {
	return &coinRepository{db: db}
}

// GetAll retrieves all coins from the database
func (r *coinRepository) GetAll() ([]domain.Coin, error) {
	var coins []domain.Coin
	if err := r.db.Find(&coins).Error; err != nil {
		return nil, fmt.Errorf("failed to get all coins: %w", err)
	}
	return coins, nil
}

// GetByCoingeckoID retrieves a coin by its CoinGecko ID
func (r *coinRepository) GetByCoingeckoID(coingeckoID string) (*domain.Coin, error) {
	var coin domain.Coin
	if err := r.db.Where("coingecko_id = ?", coingeckoID).First(&coin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get coin by coingecko_id: %w", err)
	}
	return &coin, nil
}

// Upsert creates a new coin or updates an existing one
func (r *coinRepository) Upsert(coin domain.Coin) error {
	// Set CreatedAt and UpdatedAt for new records or update UpdatedAt for existing
	if coin.CreatedAt.IsZero() {
		coin.CreatedAt = time.Now()
	}
	coin.UpdatedAt = time.Now()

	if err := r.db.
		Where(domain.Coin{CoingeckoID: coin.CoingeckoID}).
		Assign(coin).
		FirstOrCreate(&coin).Error; err != nil {
		return fmt.Errorf("failed to upsert coin: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple coins in a single transaction
func (r *coinRepository) UpsertBatch(coins []domain.Coin) error {
	if len(coins) == 0 {
		return nil
	}

	// Filter out coins with empty coingecko_id
	validCoins := make([]domain.Coin, 0, len(coins))
	for _, coin := range coins {
		if coin.CoingeckoID != "" {
			validCoins = append(validCoins, coin)
		} else {
			logger.GetLogger().WithField("symbol", coin.Symbol).Warn("Skipping coin with empty coingecko_id")
		}
	}

	if len(validCoins) == 0 {
		logger.GetLogger().Warn("No valid coins to upsert")
		return nil
	}

	// Use a transaction for batch upsert
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, coin := range validCoins {
			// Use ON CONFLICT for proper upsert
			if err := tx.Exec(`
				INSERT INTO coins (
					coingecko_id, symbol, name, image, current_price, market_cap, market_cap_rank,
					fully_diluted_valuation, total_volume, high_24h, low_24h, price_change_24h,
					price_change_percentage_24h, market_cap_change_24h, market_cap_change_percentage_24h,
					circulating_supply, total_supply, max_supply, ath, ath_change_percentage,
					ath_date, atl, atl_change_percentage, atl_date, last_updated,
					created_at, updated_at, deleted_at
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)
				ON CONFLICT (coingecko_id)
				DO UPDATE SET
					symbol = EXCLUDED.symbol,
					name = EXCLUDED.name,
					image = EXCLUDED.image,
					current_price = EXCLUDED.current_price,
					market_cap = EXCLUDED.market_cap,
					market_cap_rank = EXCLUDED.market_cap_rank,
					fully_diluted_valuation = EXCLUDED.fully_diluted_valuation,
					total_volume = EXCLUDED.total_volume,
					high_24h = EXCLUDED.high_24h,
					low_24h = EXCLUDED.low_24h,
					price_change_24h = EXCLUDED.price_change_24h,
					price_change_percentage_24h = EXCLUDED.price_change_percentage_24h,
					market_cap_change_24h = EXCLUDED.market_cap_change_24h,
					market_cap_change_percentage_24h = EXCLUDED.market_cap_change_percentage_24h,
					circulating_supply = EXCLUDED.circulating_supply,
					total_supply = EXCLUDED.total_supply,
					max_supply = EXCLUDED.max_supply,
					ath = EXCLUDED.ath,
					ath_change_percentage = EXCLUDED.ath_change_percentage,
					ath_date = EXCLUDED.ath_date,
					atl = EXCLUDED.atl,
					atl_change_percentage = EXCLUDED.atl_change_percentage,
					atl_date = EXCLUDED.atl_date,
					last_updated = EXCLUDED.last_updated,
					updated_at = EXCLUDED.updated_at,
					deleted_at = EXCLUDED.deleted_at
			`,
				coin.CoingeckoID, coin.Symbol, coin.Name, coin.Image, coin.CurrentPrice, coin.MarketCap, coin.MarketCapRank,
				coin.FullyDilutedValuation, coin.TotalVolume, coin.High24h, coin.Low24h, coin.PriceChange24h,
				coin.PriceChangePercentage24h, coin.MarketCapChange24h, coin.MarketCapChangePercentage24h,
				coin.CirculatingSupply, coin.TotalSupply, coin.MaxSupply, coin.Ath, coin.AthChangePercentage,
				coin.AthDate, coin.Atl, coin.AtlChangePercentage, coin.AtlDate, coin.LastUpdated,
				coin.CreatedAt, coin.UpdatedAt, coin.DeletedAt).Error; err != nil {
				logger.GetLogger().WithError(err).WithField("coin_id", coin.CoingeckoID).Error("Failed to upsert coin in batch")
				return fmt.Errorf("failed to upsert coin %s: %w", coin.CoingeckoID, err)
			}
		}
		logger.GetLogger().WithField("count", len(validCoins)).Info("Successfully upserted coins batch")
		return nil
	})
}

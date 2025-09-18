package repository

import (
	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// CoinMarketDataRepository defines the interface for coin market data operations
type CoinMarketDataRepository interface {
	GetAll() ([]domain.CoinMarketData, error)
	GetByCoinID(coinID uint) ([]domain.CoinMarketData, error)
	GetByExchangeID(exchangeID uint) ([]domain.CoinMarketData, error)
	Upsert(marketData domain.CoinMarketData) error
	UpsertBatch(marketData []domain.CoinMarketData) error
	DeleteByCoinID(coinID uint) error
}

type coinMarketDataRepository struct {
	db *gorm.DB
}

// NewCoinMarketDataRepository creates a new instance of CoinMarketDataRepository
func NewCoinMarketDataRepository(db *gorm.DB) CoinMarketDataRepository {
	return &coinMarketDataRepository{db: db}
}

// GetAll retrieves all coin market data from the database
func (r *coinMarketDataRepository) GetAll() ([]domain.CoinMarketData, error) {
	var marketData []domain.CoinMarketData
	if err := r.db.Preload("Coin").Preload("Exchange").Find(&marketData).Error; err != nil {
		return nil, fmt.Errorf("failed to get all coin market data: %w", err)
	}
	return marketData, nil
}

// GetByCoinID retrieves market data for a specific coin
func (r *coinMarketDataRepository) GetByCoinID(coinID uint) ([]domain.CoinMarketData, error) {
	var marketData []domain.CoinMarketData
	if err := r.db.Preload("Exchange").Where("coin_id = ?", coinID).Find(&marketData).Error; err != nil {
		return nil, fmt.Errorf("failed to get market data by coin_id: %w", err)
	}
	return marketData, nil
}

// GetByExchangeID retrieves market data for a specific exchange
func (r *coinMarketDataRepository) GetByExchangeID(exchangeID uint) ([]domain.CoinMarketData, error) {
	var marketData []domain.CoinMarketData
	if err := r.db.Preload("Coin").Where("exchange_id = ?", exchangeID).Find(&marketData).Error; err != nil {
		return nil, fmt.Errorf("failed to get market data by exchange_id: %w", err)
	}
	return marketData, nil
}

// Upsert creates a new market data record or updates an existing one
func (r *coinMarketDataRepository) Upsert(marketData domain.CoinMarketData) error {
	// Set CreatedAt and UpdatedAt for new records or update UpdatedAt for existing
	if marketData.CreatedAt.IsZero() {
		marketData.CreatedAt = time.Now()
	}
	marketData.UpdatedAt = time.Now()

	if err := r.db.
		Where("coin_id = ? AND exchange_id = ?", marketData.CoinID, marketData.ExchangeID).
		Assign(marketData).
		FirstOrCreate(&marketData).Error; err != nil {
		return fmt.Errorf("failed to upsert coin market data: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple market data records in a single transaction
func (r *coinMarketDataRepository) UpsertBatch(marketData []domain.CoinMarketData) error {
	if len(marketData) == 0 {
		return nil
	}

	// Filter out invalid records
	validMarketData := make([]domain.CoinMarketData, 0, len(marketData))
	for _, data := range marketData {
		if data.CoinID != 0 && data.ExchangeID != 0 && data.Price != nil {
			validMarketData = append(validMarketData, data)
		} else {
			logger.GetLogger().WithFields(map[string]interface{}{
				"coin_id":     data.CoinID,
				"exchange_id": data.ExchangeID,
				"price":       data.Price,
			}).Warn("Skipping invalid market data record")
		}
	}

	if len(validMarketData) == 0 {
		logger.GetLogger().Warn("No valid market data records to upsert")
		return nil
	}

	// Use a transaction for batch upsert
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, data := range validMarketData {
			// Use ON CONFLICT for proper upsert
			if err := tx.Exec(`
				INSERT INTO coin_market_data (
					coin_id, exchange_id, price, volume_24h, volume_percentage,
					last_updated, created_at, updated_at, deleted_at
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				ON CONFLICT (coin_id, exchange_id)
				DO UPDATE SET
					price = EXCLUDED.price,
					volume_24h = EXCLUDED.volume_24h,
					volume_percentage = EXCLUDED.volume_percentage,
					last_updated = EXCLUDED.last_updated,
					updated_at = EXCLUDED.updated_at,
					deleted_at = EXCLUDED.deleted_at
			`,
				data.CoinID, data.ExchangeID, data.Price, data.Volume24h, data.VolumePercentage,
				data.LastUpdated, data.CreatedAt, data.UpdatedAt, data.DeletedAt).Error; err != nil {
				logger.GetLogger().WithError(err).WithFields(map[string]interface{}{
					"coin_id":     data.CoinID,
					"exchange_id": data.ExchangeID,
				}).Error("Failed to upsert market data in batch")
				return fmt.Errorf("failed to upsert market data for coin %d on exchange %d: %w", data.CoinID, data.ExchangeID, err)
			}
		}
		logger.GetLogger().WithField("count", len(validMarketData)).Info("Successfully upserted coin market data batch")
		return nil
	})
}

// DeleteByCoinID deletes all market data for a specific coin
func (r *coinMarketDataRepository) DeleteByCoinID(coinID uint) error {
	if err := r.db.Where("coin_id = ?", coinID).Delete(&domain.CoinMarketData{}).Error; err != nil {
		return fmt.Errorf("failed to delete market data by coin_id: %w", err)
	}
	return nil
}

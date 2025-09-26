package repository

import (
	"cgoffline/internal/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CoinTickerRepository interface {
	Upsert(t domain.CoinTicker) error
}

type coinTickerRepository struct {
	db *gorm.DB
}

func NewCoinTickerRepository(db *gorm.DB) CoinTickerRepository {
	return &coinTickerRepository{db: db}
}

func (r *coinTickerRepository) Upsert(t domain.CoinTicker) error {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
	// Uniqueness by coin_id + page
	if err := r.db.Where("coin_id = ? AND page = ?", t.CoinID, t.Page).Assign(t).FirstOrCreate(&t).Error; err != nil {
		return fmt.Errorf("failed to upsert coin ticker: %w", err)
	}
	return nil
}

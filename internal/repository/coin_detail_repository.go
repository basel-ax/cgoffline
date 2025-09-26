package repository

import (
	"cgoffline/internal/domain"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CoinDetailRepository interface {
	Upsert(detail domain.CoinDetail) error
	GetByCoinID(coinID uint) (*domain.CoinDetail, error)
}

type coinDetailRepository struct {
	db *gorm.DB
}

func NewCoinDetailRepository(db *gorm.DB) CoinDetailRepository {
	return &coinDetailRepository{db: db}
}

func (r *coinDetailRepository) Upsert(detail domain.CoinDetail) error {
	if detail.CreatedAt.IsZero() {
		detail.CreatedAt = time.Now()
	}
	detail.UpdatedAt = time.Now()
	if err := r.db.Where(domain.CoinDetail{CoingeckoID: detail.CoingeckoID}).Assign(detail).FirstOrCreate(&detail).Error; err != nil {
		return fmt.Errorf("failed to upsert coin detail: %w", err)
	}
	return nil
}

func (r *coinDetailRepository) GetByCoinID(coinID uint) (*domain.CoinDetail, error) {
	var d domain.CoinDetail
	if err := r.db.Where("coin_id = ?", coinID).First(&d).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get coin detail by coin id: %w", err)
	}
	return &d, nil
}

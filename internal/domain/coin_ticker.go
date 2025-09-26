package domain

import (
	"time"

	"gorm.io/gorm"
)

// CoinTicker stores raw tickers payload per coin and page
type CoinTicker struct {
	ID        uint           `gorm:"primaryKey"`
	CoinID    uint           `gorm:"not null;index"`
	Page      int            `gorm:"not null"`
	RawJSON   []byte         `gorm:"type:jsonb"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package domain

import (
	"time"

	"gorm.io/gorm"
)

// CoinDetail stores detailed coin information fetched from CoinGecko /coins/{id}
type CoinDetail struct {
	ID          uint   `gorm:"primaryKey"`
	CoinID      uint   `gorm:"not null;index"` // FK to coins(id)
	CoingeckoID string `gorm:"uniqueIndex;size:100;not null"`
	RawJSON     []byte `gorm:"type:jsonb"`

	// Selected denormalized fields for quick access
	GenesisDate   *time.Time `gorm:"type:timestamptz"`
	HashingAlgo   *string    `gorm:"type:text"`
	Categories    []byte     `gorm:"type:jsonb"`
	Homepage      []byte     `gorm:"type:jsonb"`
	LastUpdatedAt *time.Time `gorm:"type:timestamptz"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

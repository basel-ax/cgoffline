package domain

import (
	"time"

	"gorm.io/gorm"
)

// CoinCategory represents a coin category from CoinGecko API
type CoinCategory struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CoingeckoID string       `json:"coingecko_id" gorm:"type:varchar(100);uniqueIndex;not null"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName returns the table name for the CoinCategory model
func (CoinCategory) TableName() string {
	return "coin_categories"
}

// CoinCategoryRepository defines the interface for coin category data operations
type CoinCategoryRepository interface {
	Create(category *CoinCategory) error
	CreateBatch(categories []CoinCategory) error
	GetByID(id uint) (*CoinCategory, error)
	GetByCoingeckoID(coingeckoID string) (*CoinCategory, error)
	GetAll() ([]CoinCategory, error)
	Update(category *CoinCategory) error
	Delete(id uint) error
	Upsert(category *CoinCategory) error
	UpsertBatch(categories []CoinCategory) error
}

// CoinCategoryService defines the interface for coin category business logic
type CoinCategoryService interface {
	FetchAndStoreCoinCategories() error
	GetAllCoinCategories() ([]CoinCategory, error)
	GetCoinCategoryByID(id uint) (*CoinCategory, error)
	GetCoinCategoryByCoingeckoID(coingeckoID string) (*CoinCategory, error)
	SyncCoinCategories() error
}

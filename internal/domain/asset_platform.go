package domain

import (
	"time"

	"gorm.io/gorm"
)

// AssetPlatform represents a blockchain platform from CoinGecko API
type AssetPlatform struct {
	ID              string         `json:"id" gorm:"primaryKey;type:varchar(50)"`
	ChainIdentifier *int64         `json:"chain_identifier" gorm:"type:bigint"`
	Name            string         `json:"name" gorm:"type:varchar(255);not null"`
	ShortName       *string        `json:"short_name" gorm:"type:varchar(100)"`
	NativeCoinID    *string        `json:"native_coin_id" gorm:"type:varchar(50)"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName returns the table name for the AssetPlatform model
func (AssetPlatform) TableName() string {
	return "asset_platforms"
}

// AssetPlatformRepository defines the interface for asset platform data operations
type AssetPlatformRepository interface {
	Create(platform *AssetPlatform) error
	CreateBatch(platforms []AssetPlatform) error
	GetByID(id string) (*AssetPlatform, error)
	GetAll() ([]AssetPlatform, error)
	Update(platform *AssetPlatform) error
	Delete(id string) error
	Upsert(platform *AssetPlatform) error
	UpsertBatch(platforms []AssetPlatform) error
}

// AssetPlatformService defines the interface for asset platform business logic
type AssetPlatformService interface {
	FetchAndStoreAssetPlatforms() error
	GetAllAssetPlatforms() ([]AssetPlatform, error)
	GetAssetPlatformByID(id string) (*AssetPlatform, error)
	SyncAssetPlatforms() error
}

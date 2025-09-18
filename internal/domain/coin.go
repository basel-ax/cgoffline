package domain

import (
	"time"

	"gorm.io/gorm"
)

// Coin represents a cryptocurrency from CoinGecko
type Coin struct {
	ID                           uint           `gorm:"primaryKey"`
	CoingeckoID                  string         `gorm:"uniqueIndex;size:100;not null"`
	Symbol                       string         `gorm:"size:20;not null"`
	Name                         string         `gorm:"size:255;not null"`
	Image                        *string        `gorm:"size:500"`
	CurrentPrice                 *float64       `gorm:"column:current_price"`
	MarketCap                    *float64       `gorm:"column:market_cap"`
	MarketCapRank                *int           `gorm:"column:market_cap_rank"`
	FullyDilutedValuation        *float64       `gorm:"column:fully_diluted_valuation"`
	TotalVolume                  *float64       `gorm:"column:total_volume"`
	High24h                      *float64       `gorm:"column:high_24h"`
	Low24h                       *float64       `gorm:"column:low_24h"`
	PriceChange24h               *float64       `gorm:"column:price_change_24h"`
	PriceChangePercentage24h     *float64       `gorm:"column:price_change_percentage_24h"`
	MarketCapChange24h           *float64       `gorm:"column:market_cap_change_24h"`
	MarketCapChangePercentage24h *float64       `gorm:"column:market_cap_change_percentage_24h"`
	CirculatingSupply            *float64       `gorm:"column:circulating_supply"`
	TotalSupply                  *float64       `gorm:"column:total_supply"`
	MaxSupply                    *float64       `gorm:"column:max_supply"`
	Ath                          *float64       `gorm:"column:ath"`
	AthChangePercentage          *float64       `gorm:"column:ath_change_percentage"`
	AthDate                      *time.Time     `gorm:"column:ath_date"`
	Atl                          *float64       `gorm:"column:atl"`
	AtlChangePercentage          *float64       `gorm:"column:atl_change_percentage"`
	AtlDate                      *time.Time     `gorm:"column:atl_date"`
	LastUpdated                  *time.Time     `gorm:"column:last_updated"`
	CreatedAt                    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt                    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt                    gorm.DeletedAt `gorm:"index"`
}

// CoinMarketData represents market data for a coin on a specific exchange
type CoinMarketData struct {
	ID               uint           `gorm:"primaryKey"`
	CoinID           uint           `gorm:"not null;index"`
	ExchangeID       uint           `gorm:"not null;index"`
	Coin             Coin           `gorm:"foreignKey:CoinID"`
	Exchange         Exchange       `gorm:"foreignKey:ExchangeID"`
	Price            *float64       `gorm:"not null"`
	Volume24h        *float64       `gorm:"column:volume_24h"`
	VolumePercentage *float64       `gorm:"column:volume_percentage"`
	LastUpdated      *time.Time     `gorm:"column:last_updated"`
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

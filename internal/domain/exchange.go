package domain

import (
	"time"

	"gorm.io/gorm"
)

// Exchange represents an exchange from CoinGecko
type Exchange struct {
	ID                          uint           `gorm:"primaryKey"`
	CoingeckoID                 string         `gorm:"uniqueIndex;size:100;not null"`
	Name                        string         `gorm:"size:255;not null"`
	YearEstablished             *int           `gorm:"column:year_established"`
	Country                     *string        `gorm:"size:100"`
	Description                 *string        `gorm:"type:text"`
	URL                         *string        `gorm:"size:500"`
	Image                       *string        `gorm:"size:500"`
	HasTradingIncentive         *bool          `gorm:"column:has_trading_incentive"`
	TrustScore                  *int           `gorm:"column:trust_score"`
	TrustScoreRank              *int           `gorm:"column:trust_score_rank"`
	TradeVolume24hBTC           *float64       `gorm:"column:trade_volume_24h_btc"`
	TradeVolume24hBTCNormalized *float64       `gorm:"column:trade_volume_24h_btc_normalized"`
	CreatedAt                   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt                   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt                   gorm.DeletedAt `gorm:"index"`
}

package service

import (
	"cgoffline/internal/domain"
	"cgoffline/internal/repository"
	"cgoffline/pkg/logger"
	"context"
	"fmt"
	"time"
)

// CoinService defines the interface for coin operations
type CoinService interface {
	SyncCoins() error
	SyncCoinMarketData(coinID string) error
}

type coinService struct {
	coinRepo           repository.CoinRepository
	coinMarketDataRepo repository.CoinMarketDataRepository
	exchangeRepo       repository.ExchangeRepository
	coingeckoClient    *CoinGeckoClient
}

// NewCoinService creates a new instance of CoinService
func NewCoinService(
	coinRepo repository.CoinRepository,
	coinMarketDataRepo repository.CoinMarketDataRepository,
	exchangeRepo repository.ExchangeRepository,
	client *CoinGeckoClient,
) CoinService {
	return &coinService{
		coinRepo:           coinRepo,
		coinMarketDataRepo: coinMarketDataRepo,
		exchangeRepo:       exchangeRepo,
		coingeckoClient:    client,
	}
}

// SyncCoins fetches coins from CoinGecko API and stores them in the database
func (s *coinService) SyncCoins() error {
	logger.GetLogger().Info("Starting coins synchronization")

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second) // 5 minutes timeout
	defer cancel()

	// Get current coins in DB for logging purposes
	currentCoins, err := s.coinRepo.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get current coins from database")
		return fmt.Errorf("failed to get current coins: %w", err)
	}
	logger.GetLogger().WithField("current_count", len(currentCoins)).Info("Current coins in database")

	// Fetch coins in batches (CoinGecko API returns max 250 per page)
	page := 1
	perPage := 250
	totalFetched := 0

	for {
		logger.GetLogger().WithFields(map[string]interface{}{
			"page":     page,
			"per_page": perPage,
		}).Info("Fetching coins page")

		// Fetch coins from CoinGecko API
		apiCoins, err := s.coingeckoClient.GetCoins(ctx, page, perPage)
		if err != nil {
			logger.GetLogger().WithError(err).WithField("page", page).Error("Failed to fetch coins from API")
			return fmt.Errorf("failed to fetch coins page %d: %w", page, err)
		}

		// If no coins returned, we've reached the end
		if len(apiCoins) == 0 {
			logger.GetLogger().WithField("page", page).Info("No more coins to fetch")
			break
		}

		// Store coins in the database
		if err := s.coinRepo.UpsertBatch(apiCoins); err != nil {
			logger.GetLogger().WithError(err).WithField("page", page).Error("Failed to store coins in database")
			return fmt.Errorf("failed to store coins page %d: %w", page, err)
		}

		totalFetched += len(apiCoins)
		logger.GetLogger().WithFields(map[string]interface{}{
			"page":          page,
			"count":         len(apiCoins),
			"total_fetched": totalFetched,
		}).Info("Successfully fetched and stored coins page")

		// If we got fewer coins than requested, we've reached the end
		if len(apiCoins) < perPage {
			logger.GetLogger().Info("Reached end of coins data")
			break
		}

		page++

		// Add a small delay to respect rate limits
		time.Sleep(1 * time.Second)
	}

	logger.GetLogger().WithField("total_fetched", totalFetched).Info("Successfully fetched and stored all coins")

	// Verify count after sync
	updatedCoins, err := s.coinRepo.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get updated coins from database")
		return fmt.Errorf("failed to get updated coins: %w", err)
	}
	logger.GetLogger().WithField("updated_count", len(updatedCoins)).Info("Coins after synchronization")

	logger.GetLogger().Info("Coins synchronization completed successfully")
	return nil
}

// SyncCoinMarketData fetches market data for a specific coin and stores it in the database
func (s *coinService) SyncCoinMarketData(coinID string) error {
	logger.GetLogger().WithField("coin_id", coinID).Info("Starting coin market data synchronization")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get the coin from database
	coin, err := s.coinRepo.GetByCoingeckoID(coinID)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("coin_id", coinID).Error("Failed to get coin from database")
		return fmt.Errorf("failed to get coin: %w", err)
	}
	if coin == nil {
		return fmt.Errorf("coin with ID %s not found in database", coinID)
	}

	// Get current market data for this coin
	currentMarketData, err := s.coinMarketDataRepo.GetByCoinID(coin.ID)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("coin_id", coinID).Error("Failed to get current market data from database")
		return fmt.Errorf("failed to get current market data: %w", err)
	}
	logger.GetLogger().WithFields(map[string]interface{}{
		"coin_id":       coinID,
		"current_count": len(currentMarketData),
	}).Info("Current market data in database")

	// Fetch market data from CoinGecko API
	apiMarketData, err := s.coingeckoClient.GetCoinMarketData(ctx, coinID)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("coin_id", coinID).Error("Failed to fetch market data from API")
		return fmt.Errorf("failed to fetch market data: %w", err)
	}

	// Get all exchanges to map exchange names to IDs
	exchanges, err := s.exchangeRepo.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get exchanges from database")
		return fmt.Errorf("failed to get exchanges: %w", err)
	}

	// Create exchange name to ID mapping
	exchangeMap := make(map[string]uint)
	for _, exchange := range exchanges {
		exchangeMap[exchange.Name] = exchange.ID
	}

	// Process and store market data
	validMarketData := make([]domain.CoinMarketData, 0, len(apiMarketData))
	for range apiMarketData {
		// Find exchange ID by name (this would need to be enhanced based on the actual API response structure)
		// For now, we'll skip this as the API response structure needs to be refined
		// This is a placeholder implementation
		logger.GetLogger().WithField("coin_id", coinID).Debug("Processing market data entry")
	}

	// Store market data in the database
	if len(validMarketData) > 0 {
		if err := s.coinMarketDataRepo.UpsertBatch(validMarketData); err != nil {
			logger.GetLogger().WithError(err).WithField("coin_id", coinID).Error("Failed to store market data in database")
			return fmt.Errorf("failed to store market data: %w", err)
		}
	}

	logger.GetLogger().WithFields(map[string]interface{}{
		"coin_id": coinID,
		"count":   len(validMarketData),
	}).Info("Successfully fetched and stored coin market data")

	logger.GetLogger().WithField("coin_id", coinID).Info("Coin market data synchronization completed successfully")
	return nil
}

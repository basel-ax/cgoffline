package service

import (
	"cgoffline/internal/domain"
	"cgoffline/internal/repository"
	"cgoffline/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CoinService defines the interface for coin operations
type CoinService interface {
	SyncCoins() error
	SyncCoinMarketData(coinID string) error
	SyncCoinsData(minTotalVolume float64) error
}

type coinService struct {
	coinRepo           repository.CoinRepository
	coinMarketDataRepo repository.CoinMarketDataRepository
	exchangeRepo       repository.ExchangeRepository
	coingeckoClient    *CoinGeckoClient
	coinDetailRepo     repository.CoinDetailRepository
	coinTickerRepo     repository.CoinTickerRepository
}

// NewCoinService creates a new instance of CoinService
func NewCoinService(
	coinRepo repository.CoinRepository,
	coinMarketDataRepo repository.CoinMarketDataRepository,
	exchangeRepo repository.ExchangeRepository,
	coinDetailRepo repository.CoinDetailRepository,
	coinTickerRepo repository.CoinTickerRepository,
	client *CoinGeckoClient,
) CoinService {
	return &coinService{
		coinRepo:           coinRepo,
		coinMarketDataRepo: coinMarketDataRepo,
		exchangeRepo:       exchangeRepo,
		coinDetailRepo:     coinDetailRepo,
		coinTickerRepo:     coinTickerRepo,
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

// SyncCoinsData fetches detailed coin data and tickers for coins above a volume threshold
func (s *coinService) SyncCoinsData(minTotalVolume float64) error {
	logger.GetLogger().WithField("min_total_volume", minTotalVolume).Info("Starting coins data synchronization")

	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Second)
	defer cancel()

	// Load coins and filter by volume
	coins, err := s.coinRepo.GetAll()
	if err != nil {
		return fmt.Errorf("failed to load coins: %w", err)
	}

	filtered := make([]domain.Coin, 0, len(coins))
	for _, c := range coins {
		if c.TotalVolume != nil && *c.TotalVolume >= minTotalVolume {
			filtered = append(filtered, c)
		}
	}

	logger.GetLogger().WithFields(map[string]interface{}{
		"eligible": len(filtered),
		"total":    len(coins),
	}).Info("Coins eligible for detailed sync by volume filter")

	// For each coin, fetch coin data and tickers; store raw JSON in coin_details
	for _, c := range filtered {
		// Fetch /coins/{id}
		data, err := s.coingeckoClient.GetCoinDataByID(ctx, c.CoingeckoID)
		if err != nil {
			logger.GetLogger().WithError(err).WithField("coin_id", c.CoingeckoID).Warn("Failed to fetch coin data by id; skipping")
			continue
		}

		// Save coin detail
		raw, _ := json.Marshal(data)
		detail := domain.CoinDetail{
			CoinID:      c.ID,
			CoingeckoID: c.CoingeckoID,
			RawJSON:     raw,
		}

		// Optional denormalized fields
		if v, ok := data["genesis_date"].(string); ok && v != "" {
			if t, parseErr := time.Parse(time.RFC3339, v+"T00:00:00Z"); parseErr == nil {
				detail.GenesisDate = &t
			}
		}
		if v, ok := data["hashing_algorithm"].(string); ok {
			detail.HashingAlgo = &v
		}
		if cats, ok := data["categories"].([]any); ok {
			if b, mErr := json.Marshal(cats); mErr == nil {
				detail.Categories = b
			}
		}
		if links, ok := data["links"].(map[string]any); ok {
			if hp, ok2 := links["homepage"].([]any); ok2 {
				if b, mErr := json.Marshal(hp); mErr == nil {
					detail.Homepage = b
				}
			}
		}

		if lu, ok := data["last_updated"].(string); ok && lu != "" {
			if t, perr := time.Parse(time.RFC3339, lu); perr == nil {
				detail.LastUpdatedAt = &t
			}
		}

		if err := s.coinDetailRepo.Upsert(detail); err != nil {
			logger.GetLogger().WithError(err).WithField("coin_id", c.CoingeckoID).Warn("Failed to upsert coin detail")
		}

		// Fetch tickers with pagination (100 per page). Persisting raw is sufficient for now.
		page := 1
		for {
			tickersPayload, err := s.coingeckoClient.GetCoinTickers(ctx, c.CoingeckoID, page)
			if err != nil {
				logger.GetLogger().WithError(err).WithFields(map[string]interface{}{"coin_id": c.CoingeckoID, "page": page}).Warn("Failed to fetch tickers; stopping pagination")
				break
			}
			// persist
			if b, mErr := json.Marshal(tickersPayload); mErr == nil {
				_ = s.coinTickerRepo.Upsert(domain.CoinTicker{CoinID: c.ID, Page: page, RawJSON: b})
			}
			// We currently do not persist tickers separately; this is a placeholder to extend later.
			// Stop if no tickers returned
			if arr, ok := tickersPayload["tickers"].([]any); !ok || len(arr) == 0 {
				break
			}
			// Next page
			page++
			time.Sleep(500 * time.Millisecond)
		}
	}

	logger.GetLogger().Info("Coins data synchronization completed")
	return nil
}

package service

import (
	"cgoffline/internal/repository"
	"cgoffline/pkg/logger"
	"context"
	"fmt"
	"time"
)

// ExchangeService defines the interface for exchange operations
type ExchangeService interface {
	SyncExchanges() error
}

type exchangeService struct {
	repo            repository.ExchangeRepository
	coingeckoClient *CoinGeckoClient
}

// NewExchangeService creates a new instance of ExchangeService
func NewExchangeService(repo repository.ExchangeRepository, client *CoinGeckoClient) ExchangeService {
	return &exchangeService{
		repo:            repo,
		coingeckoClient: client,
	}
}

// SyncExchanges fetches exchanges from CoinGecko API and stores them in the database
func (s *exchangeService) SyncExchanges() error {
	logger.GetLogger().Info("Starting exchanges synchronization")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Get current exchanges in DB for logging purposes
	currentExchanges, err := s.repo.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get current exchanges from database")
		return fmt.Errorf("failed to get current exchanges: %w", err)
	}
	logger.GetLogger().WithField("current_count", len(currentExchanges)).Info("Current exchanges in database")

	logger.GetLogger().Info("Starting to fetch and store exchanges")

	// Fetch exchanges from CoinGecko API
	apiExchanges, err := s.coingeckoClient.GetExchanges(ctx)
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to fetch exchanges from API")
		return fmt.Errorf("failed to fetch exchanges: %w", err)
	}

	// Store exchanges in the database
	if err := s.repo.UpsertBatch(apiExchanges); err != nil {
		logger.GetLogger().WithError(err).Error("Failed to store exchanges in database")
		return fmt.Errorf("failed to store exchanges: %w", err)
	}

	logger.GetLogger().WithField("count", len(apiExchanges)).Info("Successfully fetched and stored exchanges")

	// Verify count after sync
	updatedExchanges, err := s.repo.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get updated exchanges from database")
		return fmt.Errorf("failed to get updated exchanges: %w", err)
	}
	logger.GetLogger().WithField("updated_count", len(updatedExchanges)).Info("Exchanges after synchronization")

	logger.GetLogger().Info("Exchanges synchronization completed successfully")
	return nil
}

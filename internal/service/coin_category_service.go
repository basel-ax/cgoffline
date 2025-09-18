package service

import (
	"context"
	"fmt"
	"time"

	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"
)

// coinCategoryService implements the CoinCategoryService interface
type coinCategoryService struct {
	repository domain.CoinCategoryRepository
	apiClient  *CoinGeckoClient
}

// NewCoinCategoryService creates a new coin category service
func NewCoinCategoryService(repository domain.CoinCategoryRepository, apiClient *CoinGeckoClient) domain.CoinCategoryService {
	return &coinCategoryService{
		repository: repository,
		apiClient:  apiClient,
	}
}

// FetchAndStoreCoinCategories fetches coin categories from CoinGecko API and stores them in the database
func (s *coinCategoryService) FetchAndStoreCoinCategories() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	logger.GetLogger().Info("Starting to fetch and store coin categories")

	// Check API health first
	if err := s.apiClient.HealthCheck(ctx); err != nil {
		logger.GetLogger().WithError(err).Error("CoinGecko API health check failed")
		return fmt.Errorf("API health check failed: %w", err)
	}

	// Fetch categories from API
	categories, err := s.apiClient.GetCoinCategories(ctx)
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to fetch coin categories from API")
		return fmt.Errorf("failed to fetch coin categories: %w", err)
	}

	if len(categories) == 0 {
		logger.GetLogger().Warn("No coin categories received from API")
		return fmt.Errorf("no coin categories received from API")
	}

	// Store categories in database using upsert to handle updates
	if err := s.repository.UpsertBatch(categories); err != nil {
		logger.GetLogger().WithError(err).Error("Failed to store coin categories in database")
		return fmt.Errorf("failed to store coin categories: %w", err)
	}

	logger.GetLogger().WithField("count", len(categories)).Info("Successfully fetched and stored coin categories")
	return nil
}

// GetAllCoinCategories retrieves all coin categories from the database
func (s *coinCategoryService) GetAllCoinCategories() ([]domain.CoinCategory, error) {
	logger.GetLogger().Info("Retrieving all coin categories from database")

	categories, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to retrieve coin categories from database")
		return nil, fmt.Errorf("failed to retrieve coin categories: %w", err)
	}

	logger.GetLogger().WithField("count", len(categories)).Info("Successfully retrieved coin categories from database")
	return categories, nil
}

// GetCoinCategoryByID retrieves a specific coin category by ID
func (s *coinCategoryService) GetCoinCategoryByID(id uint) (*domain.CoinCategory, error) {
	logger.GetLogger().WithField("category_id", id).Info("Retrieving coin category by ID")

	category, err := s.repository.GetByID(id)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("category_id", id).Error("Failed to retrieve coin category by ID")
		return nil, fmt.Errorf("failed to retrieve coin category: %w", err)
	}

	logger.GetLogger().WithField("category_id", id).Info("Successfully retrieved coin category by ID")
	return category, nil
}

// GetCoinCategoryByCoingeckoID retrieves a specific coin category by CoinGecko ID
func (s *coinCategoryService) GetCoinCategoryByCoingeckoID(coingeckoID string) (*domain.CoinCategory, error) {
	logger.GetLogger().WithField("coingecko_id", coingeckoID).Info("Retrieving coin category by CoinGecko ID")

	category, err := s.repository.GetByCoingeckoID(coingeckoID)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("coingecko_id", coingeckoID).Error("Failed to retrieve coin category by CoinGecko ID")
		return nil, fmt.Errorf("failed to retrieve coin category: %w", err)
	}

	logger.GetLogger().WithField("coingecko_id", coingeckoID).Info("Successfully retrieved coin category by CoinGecko ID")
	return category, nil
}

// SyncCoinCategories synchronizes coin categories with the CoinGecko API
// This method fetches fresh data and updates the database
func (s *coinCategoryService) SyncCoinCategories() error {
	logger.GetLogger().Info("Starting coin categories synchronization")

	// Get current count from database
	currentCategories, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Warn("Failed to get current category count, proceeding with sync")
	} else {
		logger.GetLogger().WithField("current_count", len(currentCategories)).Info("Current coin categories in database")
	}

	// Fetch and store fresh data
	if err := s.FetchAndStoreCoinCategories(); err != nil {
		return fmt.Errorf("failed to sync coin categories: %w", err)
	}

	// Get updated count
	updatedCategories, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Warn("Failed to get updated category count")
	} else {
		logger.GetLogger().WithField("updated_count", len(updatedCategories)).Info("Coin categories after synchronization")
	}

	logger.GetLogger().Info("Coin categories synchronization completed successfully")
	return nil
}

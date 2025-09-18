package service

import (
	"context"
	"fmt"
	"time"

	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"
)

// assetPlatformService implements the AssetPlatformService interface
type assetPlatformService struct {
	repository domain.AssetPlatformRepository
	apiClient  *CoinGeckoClient
}

// NewAssetPlatformService creates a new asset platform service
func NewAssetPlatformService(repository domain.AssetPlatformRepository, apiClient *CoinGeckoClient) domain.AssetPlatformService {
	return &assetPlatformService{
		repository: repository,
		apiClient:  apiClient,
	}
}

// FetchAndStoreAssetPlatforms fetches asset platforms from CoinGecko API and stores them in the database
func (s *assetPlatformService) FetchAndStoreAssetPlatforms() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	logger.GetLogger().Info("Starting to fetch and store asset platforms")

	// Check API health first
	if err := s.apiClient.HealthCheck(ctx); err != nil {
		logger.GetLogger().WithError(err).Error("CoinGecko API health check failed")
		return fmt.Errorf("API health check failed: %w", err)
	}

	// Fetch platforms from API
	platforms, err := s.apiClient.GetAssetPlatforms(ctx)
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to fetch asset platforms from API")
		return fmt.Errorf("failed to fetch asset platforms: %w", err)
	}

	if len(platforms) == 0 {
		logger.GetLogger().Warn("No asset platforms received from API")
		return fmt.Errorf("no asset platforms received from API")
	}

	// Store platforms in database using upsert to handle updates
	if err := s.repository.UpsertBatch(platforms); err != nil {
		logger.GetLogger().WithError(err).Error("Failed to store asset platforms in database")
		return fmt.Errorf("failed to store asset platforms: %w", err)
	}

	logger.GetLogger().WithField("count", len(platforms)).Info("Successfully fetched and stored asset platforms")
	return nil
}

// GetAllAssetPlatforms retrieves all asset platforms from the database
func (s *assetPlatformService) GetAllAssetPlatforms() ([]domain.AssetPlatform, error) {
	logger.GetLogger().Info("Retrieving all asset platforms from database")

	platforms, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Error("Failed to retrieve asset platforms from database")
		return nil, fmt.Errorf("failed to retrieve asset platforms: %w", err)
	}

	logger.GetLogger().WithField("count", len(platforms)).Info("Successfully retrieved asset platforms from database")
	return platforms, nil
}

// GetAssetPlatformByID retrieves a specific asset platform by ID
func (s *assetPlatformService) GetAssetPlatformByID(id string) (*domain.AssetPlatform, error) {
	logger.GetLogger().WithField("platform_id", id).Info("Retrieving asset platform by ID")

	platform, err := s.repository.GetByID(id)
	if err != nil {
		logger.GetLogger().WithError(err).WithField("platform_id", id).Error("Failed to retrieve asset platform by ID")
		return nil, fmt.Errorf("failed to retrieve asset platform: %w", err)
	}

	logger.GetLogger().WithField("platform_id", id).Info("Successfully retrieved asset platform by ID")
	return platform, nil
}

// SyncAssetPlatforms synchronizes asset platforms with the CoinGecko API
// This method fetches fresh data and updates the database
func (s *assetPlatformService) SyncAssetPlatforms() error {
	logger.GetLogger().Info("Starting asset platforms synchronization")

	// Get current count from database
	currentPlatforms, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Warn("Failed to get current platform count, proceeding with sync")
	} else {
		logger.GetLogger().WithField("current_count", len(currentPlatforms)).Info("Current asset platforms in database")
	}

	// Fetch and store fresh data
	if err := s.FetchAndStoreAssetPlatforms(); err != nil {
		return fmt.Errorf("failed to sync asset platforms: %w", err)
	}

	// Get updated count
	updatedPlatforms, err := s.repository.GetAll()
	if err != nil {
		logger.GetLogger().WithError(err).Warn("Failed to get updated platform count")
	} else {
		logger.GetLogger().WithField("updated_count", len(updatedPlatforms)).Info("Asset platforms after synchronization")
	}

	logger.GetLogger().Info("Asset platforms synchronization completed successfully")
	return nil
}

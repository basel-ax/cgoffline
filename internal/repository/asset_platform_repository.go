package repository

import (
	"fmt"

	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"

	"gorm.io/gorm"
)

// assetPlatformRepository implements the AssetPlatformRepository interface
type assetPlatformRepository struct {
	db *gorm.DB
}

// NewAssetPlatformRepository creates a new asset platform repository
func NewAssetPlatformRepository(db *gorm.DB) domain.AssetPlatformRepository {
	return &assetPlatformRepository{
		db: db,
	}
}

// Create creates a new asset platform
func (r *assetPlatformRepository) Create(platform *domain.AssetPlatform) error {
	if err := r.db.Create(platform).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("platform_id", platform.ID).Error("Failed to create asset platform")
		return fmt.Errorf("failed to create asset platform: %w", err)
	}
	return nil
}

// CreateBatch creates multiple asset platforms in a single transaction
func (r *assetPlatformRepository) CreateBatch(platforms []domain.AssetPlatform) error {
	if len(platforms) == 0 {
		return nil
	}

	if err := r.db.CreateInBatches(platforms, 100).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("count", len(platforms)).Error("Failed to create asset platforms batch")
		return fmt.Errorf("failed to create asset platforms batch: %w", err)
	}

	logger.GetLogger().WithField("count", len(platforms)).Info("Successfully created asset platforms batch")
	return nil
}

// GetByID retrieves an asset platform by ID
func (r *assetPlatformRepository) GetByID(id string) (*domain.AssetPlatform, error) {
	var platform domain.AssetPlatform
	if err := r.db.First(&platform, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("asset platform with ID %s not found", id)
		}
		logger.GetLogger().WithError(err).WithField("platform_id", id).Error("Failed to get asset platform by ID")
		return nil, fmt.Errorf("failed to get asset platform by ID: %w", err)
	}
	return &platform, nil
}

// GetAll retrieves all asset platforms
func (r *assetPlatformRepository) GetAll() ([]domain.AssetPlatform, error) {
	var platforms []domain.AssetPlatform
	if err := r.db.Find(&platforms).Error; err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get all asset platforms")
		return nil, fmt.Errorf("failed to get all asset platforms: %w", err)
	}
	return platforms, nil
}

// Update updates an existing asset platform
func (r *assetPlatformRepository) Update(platform *domain.AssetPlatform) error {
	if err := r.db.Save(platform).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("platform_id", platform.ID).Error("Failed to update asset platform")
		return fmt.Errorf("failed to update asset platform: %w", err)
	}
	return nil
}

// Delete soft deletes an asset platform
func (r *assetPlatformRepository) Delete(id string) error {
	if err := r.db.Delete(&domain.AssetPlatform{}, "id = ?", id).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("platform_id", id).Error("Failed to delete asset platform")
		return fmt.Errorf("failed to delete asset platform: %w", err)
	}
	return nil
}

// Upsert creates or updates an asset platform
func (r *assetPlatformRepository) Upsert(platform *domain.AssetPlatform) error {
	if err := r.db.Save(platform).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("platform_id", platform.ID).Error("Failed to upsert asset platform")
		return fmt.Errorf("failed to upsert asset platform: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple asset platforms in a single transaction
func (r *assetPlatformRepository) UpsertBatch(platforms []domain.AssetPlatform) error {
	if len(platforms) == 0 {
		return nil
	}

	// Use a transaction for batch upsert
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, platform := range platforms {
			if err := tx.Save(&platform).Error; err != nil {
				logger.GetLogger().WithError(err).WithField("platform_id", platform.ID).Error("Failed to upsert asset platform in batch")
				return fmt.Errorf("failed to upsert asset platform %s: %w", platform.ID, err)
			}
		}
		logger.GetLogger().WithField("count", len(platforms)).Info("Successfully upserted asset platforms batch")
		return nil
	})
}

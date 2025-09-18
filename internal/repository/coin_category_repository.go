package repository

import (
	"fmt"

	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"

	"gorm.io/gorm"
)

// coinCategoryRepository implements the CoinCategoryRepository interface
type coinCategoryRepository struct {
	db *gorm.DB
}

// NewCoinCategoryRepository creates a new coin category repository
func NewCoinCategoryRepository(db *gorm.DB) domain.CoinCategoryRepository {
	return &coinCategoryRepository{
		db: db,
	}
}

// Create creates a new coin category
func (r *coinCategoryRepository) Create(category *domain.CoinCategory) error {
	if err := r.db.Create(category).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("category_id", category.CoingeckoID).Error("Failed to create coin category")
		return fmt.Errorf("failed to create coin category: %w", err)
	}
	return nil
}

// CreateBatch creates multiple coin categories in a single transaction
func (r *coinCategoryRepository) CreateBatch(categories []domain.CoinCategory) error {
	if len(categories) == 0 {
		return nil
	}

	if err := r.db.CreateInBatches(categories, 100).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("count", len(categories)).Error("Failed to create coin categories batch")
		return fmt.Errorf("failed to create coin categories batch: %w", err)
	}

	logger.GetLogger().WithField("count", len(categories)).Info("Successfully created coin categories batch")
	return nil
}

// GetByID retrieves a coin category by ID
func (r *coinCategoryRepository) GetByID(id uint) (*domain.CoinCategory, error) {
	var category domain.CoinCategory
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("coin category with ID %d not found", id)
		}
		logger.GetLogger().WithError(err).WithField("category_id", id).Error("Failed to get coin category by ID")
		return nil, fmt.Errorf("failed to get coin category by ID: %w", err)
	}
	return &category, nil
}

// GetByCoingeckoID retrieves a coin category by CoinGecko ID
func (r *coinCategoryRepository) GetByCoingeckoID(coingeckoID string) (*domain.CoinCategory, error) {
	var category domain.CoinCategory
	if err := r.db.First(&category, "coingecko_id = ?", coingeckoID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("coin category with CoinGecko ID %s not found", coingeckoID)
		}
		logger.GetLogger().WithError(err).WithField("coingecko_id", coingeckoID).Error("Failed to get coin category by CoinGecko ID")
		return nil, fmt.Errorf("failed to get coin category by CoinGecko ID: %w", err)
	}
	return &category, nil
}

// GetAll retrieves all coin categories
func (r *coinCategoryRepository) GetAll() ([]domain.CoinCategory, error) {
	var categories []domain.CoinCategory
	if err := r.db.Find(&categories).Error; err != nil {
		logger.GetLogger().WithError(err).Error("Failed to get all coin categories")
		return nil, fmt.Errorf("failed to get all coin categories: %w", err)
	}
	return categories, nil
}

// Update updates an existing coin category
func (r *coinCategoryRepository) Update(category *domain.CoinCategory) error {
	if err := r.db.Save(category).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("category_id", category.CoingeckoID).Error("Failed to update coin category")
		return fmt.Errorf("failed to update coin category: %w", err)
	}
	return nil
}

// Delete soft deletes a coin category
func (r *coinCategoryRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.CoinCategory{}, "id = ?", id).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("category_id", id).Error("Failed to delete coin category")
		return fmt.Errorf("failed to delete coin category: %w", err)
	}
	return nil
}

// Upsert creates or updates a coin category
func (r *coinCategoryRepository) Upsert(category *domain.CoinCategory) error {
	if err := r.db.Save(category).Error; err != nil {
		logger.GetLogger().WithError(err).WithField("category_id", category.CoingeckoID).Error("Failed to upsert coin category")
		return fmt.Errorf("failed to upsert coin category: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple coin categories in a single transaction
func (r *coinCategoryRepository) UpsertBatch(categories []domain.CoinCategory) error {
	if len(categories) == 0 {
		return nil
	}

	// Filter out categories with empty coingecko_id
	validCategories := make([]domain.CoinCategory, 0, len(categories))
	for _, category := range categories {
		if category.CoingeckoID != "" {
			validCategories = append(validCategories, category)
		} else {
			logger.GetLogger().WithField("name", category.Name).Warn("Skipping category with empty coingecko_id")
		}
	}

	if len(validCategories) == 0 {
		logger.GetLogger().Warn("No valid categories to upsert")
		return nil
	}

	// Use a transaction for batch upsert
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, category := range validCategories {
			// Use ON CONFLICT for proper upsert
			if err := tx.Exec(`
				INSERT INTO coin_categories (coingecko_id, name, created_at, updated_at, deleted_at)
				VALUES ($1, $2, $3, $4, $5)
				ON CONFLICT (coingecko_id) 
				DO UPDATE SET 
					name = EXCLUDED.name,
					updated_at = EXCLUDED.updated_at,
					deleted_at = EXCLUDED.deleted_at
			`, category.CoingeckoID, category.Name, category.CreatedAt, category.UpdatedAt, category.DeletedAt).Error; err != nil {
				logger.GetLogger().WithError(err).WithField("category_id", category.CoingeckoID).Error("Failed to upsert coin category in batch")
				return fmt.Errorf("failed to upsert coin category %s: %w", category.CoingeckoID, err)
			}
		}
		logger.GetLogger().WithField("count", len(validCategories)).Info("Successfully upserted coin categories batch")
		return nil
	})
}

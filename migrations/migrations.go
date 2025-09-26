package migrations

import (
	"cgoffline/internal/domain"
	"cgoffline/pkg/logger"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// GetMigrations returns all database migrations
func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "2024010101",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create asset_platforms table")
				return tx.AutoMigrate(&domain.AssetPlatform{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop asset_platforms table")
				return tx.Migrator().DropTable(&domain.AssetPlatform{})
			},
		},
		{
			ID: "2024010103",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create coin_categories table")
				return tx.AutoMigrate(&domain.CoinCategory{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop coin_categories table")
				return tx.Migrator().DropTable(&domain.CoinCategory{})
			},
		},
		{
			ID: "2024010104",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create exchanges table")
				return tx.AutoMigrate(&domain.Exchange{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop exchanges table")
				return tx.Migrator().DropTable(&domain.Exchange{})
			},
		},
		{
			ID: "2024010105",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create coins table")
				return tx.AutoMigrate(&domain.Coin{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop coins table")
				return tx.Migrator().DropTable(&domain.Coin{})
			},
		},
		{
			ID: "2024010106",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create coin_market_data table")
				return tx.AutoMigrate(&domain.CoinMarketData{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop coin_market_data table")
				return tx.Migrator().DropTable(&domain.CoinMarketData{})
			},
		},
		{
			ID: "2024010102",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Add indexes to asset_platforms table")

				// Add index on chain_identifier for faster lookups
				if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_asset_platforms_chain_identifier ON asset_platforms(chain_identifier)").Error; err != nil {
					return err
				}

				// Add index on name for faster searches
				if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_asset_platforms_name ON asset_platforms(name)").Error; err != nil {
					return err
				}

				// Add index on native_coin_id for faster lookups
				if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_asset_platforms_native_coin_id ON asset_platforms(native_coin_id)").Error; err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop indexes from asset_platforms table")

				// Drop indexes
				tx.Exec("DROP INDEX IF EXISTS idx_asset_platforms_chain_identifier")
				tx.Exec("DROP INDEX IF EXISTS idx_asset_platforms_name")
				tx.Exec("DROP INDEX IF EXISTS idx_asset_platforms_native_coin_id")

				return nil
			},
		},
		{
			ID: "2024010107",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create coin_details table")
				return tx.AutoMigrate(&domain.CoinDetail{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop coin_details table")
				return tx.Migrator().DropTable(&domain.CoinDetail{})
			},
		},
		{
			ID: "2024010108",
			Migrate: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Running migration: Create coin_tickers table")
				return tx.AutoMigrate(&domain.CoinTicker{})
			},
			Rollback: func(tx *gorm.DB) error {
				logger.GetLogger().Info("Rolling back migration: Drop coin_tickers table")
				return tx.Migrator().DropTable(&domain.CoinTicker{})
			},
		},
	}
}

// RunMigrations runs all pending migrations
func RunMigrations(db *gorm.DB) error {
	logger.GetLogger().Info("Starting database migrations")

	options := gormigrate.DefaultOptions
	options.TableName = "gorm_migrations"

	m := gormigrate.New(db, options, GetMigrations())

	if err := m.Migrate(); err != nil {
		logger.GetLogger().WithError(err).Error("Failed to run migrations")
		return err
	}

	logger.GetLogger().Info("Database migrations completed successfully")
	return nil
}

// RollbackLastMigration rolls back the last migration
func RollbackLastMigration(db *gorm.DB) error {
	logger.GetLogger().Info("Rolling back last migration")

	options := gormigrate.DefaultOptions
	options.TableName = "gorm_migrations"

	m := gormigrate.New(db, options, GetMigrations())

	if err := m.RollbackLast(); err != nil {
		logger.GetLogger().WithError(err).Error("Failed to rollback last migration")
		return err
	}

	logger.GetLogger().Info("Last migration rolled back successfully")
	return nil
}

// GetMigrationStatus returns the current migration status
func GetMigrationStatus(db *gorm.DB) error {
	// Check if migrations table exists
	var count int64
	if err := db.Table("gorm_migrations").Count(&count).Error; err != nil {
		logger.GetLogger().WithError(err).Error("Failed to check migration status")
		return err
	}

	logger.GetLogger().WithField("migrations_applied", count).Info("Migration status")
	return nil
}

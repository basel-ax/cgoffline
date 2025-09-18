package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cgoffline/internal/repository"
	"cgoffline/internal/service"
	"cgoffline/migrations"
	"cgoffline/pkg/config"
	"cgoffline/pkg/logger"
)

func main() {
	// Parse command line flags
	var (
		syncOnly = flag.Bool("sync", false, "Only sync asset platforms and exit")
		migrate  = flag.Bool("migrate", false, "Run database migrations and exit")
		rollback = flag.Bool("rollback", false, "Rollback last migration and exit")
		status   = flag.Bool("status", false, "Show migration status and exit")
	)
	flag.Parse()

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger(cfg.Logging)
	log := logger.GetLogger()

	log.Info("Starting cgoffline application")

	// Connect to database
	db, err := repository.NewDatabase(cfg.Database)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}
	defer func() {
		if err := repository.CloseDatabase(db); err != nil {
			log.WithError(err).Error("Failed to close database connection")
		}
	}()

	// Handle migration commands
	if *migrate {
		if err := migrations.RunMigrations(db); err != nil {
			log.WithError(err).Fatal("Failed to run migrations")
		}
		log.Info("Migrations completed successfully")
		return
	}

	if *rollback {
		if err := migrations.RollbackLastMigration(db); err != nil {
			log.WithError(err).Fatal("Failed to rollback migration")
		}
		log.Info("Migration rollback completed successfully")
		return
	}

	if *status {
		if err := migrations.GetMigrationStatus(db); err != nil {
			log.WithError(err).Fatal("Failed to get migration status")
		}
		return
	}

	// Initialize repositories and services
	assetPlatformRepo := repository.NewAssetPlatformRepository(db)
	coinGeckoClient := service.NewCoinGeckoClient(cfg.API)
	assetPlatformService := service.NewAssetPlatformService(assetPlatformRepo, coinGeckoClient)

	// Handle sync-only mode
	if *syncOnly {
		log.Info("Running asset platforms synchronization")
		if err := assetPlatformService.SyncAssetPlatforms(); err != nil {
			log.WithError(err).Fatal("Failed to sync asset platforms")
		}
		log.Info("Asset platforms synchronization completed successfully")
		return
	}

	// Run initial sync
	log.Info("Running initial asset platforms synchronization")
	if err := assetPlatformService.SyncAssetPlatforms(); err != nil {
		log.WithError(err).Error("Failed to sync asset platforms")
		// Don't exit on sync failure, continue running
	}

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Application started successfully. Press Ctrl+C to stop.")

	// Wait for shutdown signal
	<-sigChan
	log.Info("Shutdown signal received, stopping application...")

	// Perform cleanup here if needed
	log.Info("Application stopped gracefully")
}

// printUsage prints usage information
func printUsage() {
	fmt.Println("Usage: cgoffline [options]")
	fmt.Println("Options:")
	fmt.Println("  -sync     Only sync asset platforms and exit")
	fmt.Println("  -migrate  Run database migrations and exit")
	fmt.Println("  -rollback Rollback last migration and exit")
	fmt.Println("  -status   Show migration status and exit")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("  DB_HOST              Database host (default: localhost)")
	fmt.Println("  DB_PORT              Database port (default: 5432)")
	fmt.Println("  DB_USER              Database user (default: postgres)")
	fmt.Println("  DB_PASSWORD          Database password (default: password)")
	fmt.Println("  DB_NAME              Database name (default: cgoffline)")
	fmt.Println("  DB_SSLMODE           Database SSL mode (default: disable)")
	fmt.Println("  DB_TIMEZONE          Database timezone (default: UTC)")
	fmt.Println("  COINGECKO_BASE_URL   CoinGecko API base URL (default: https://api.coingecko.com/api/v3)")
	fmt.Println("  API_TIMEOUT          API timeout (default: 30s)")
	fmt.Println("  API_RETRY_ATTEMPTS   API retry attempts (default: 3)")
	fmt.Println("  API_RETRY_DELAY      API retry delay (default: 1s)")
	fmt.Println("  LOG_LEVEL            Log level (default: info)")
	fmt.Println("  LOG_FORMAT           Log format (default: json)")
}

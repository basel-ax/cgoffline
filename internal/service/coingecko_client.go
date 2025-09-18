package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"cgoffline/internal/domain"
	"cgoffline/pkg/config"
	"cgoffline/pkg/logger"
)

// CoinGeckoClient handles communication with the CoinGecko API
type CoinGeckoClient struct {
	baseURL    string
	httpClient *http.Client
	retryCount int
	retryDelay time.Duration
}

// NewCoinGeckoClient creates a new CoinGecko API client
func NewCoinGeckoClient(cfg config.APIConfig) *CoinGeckoClient {
	return &CoinGeckoClient{
		baseURL: cfg.CoinGeckoBaseURL,
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		retryCount: cfg.RetryAttempts,
		retryDelay: cfg.RetryDelay,
	}
}

// AssetPlatformResponse represents the response structure from CoinGecko API
type AssetPlatformResponse struct {
	ID              string  `json:"id"`
	ChainIdentifier *int64  `json:"chain_identifier"`
	Name            string  `json:"name"`
	ShortName       *string `json:"short_name"`
	NativeCoinID    *string `json:"native_coin_id"`
}

// CoinCategoryResponse represents the response structure for coin categories from CoinGecko API
type CoinCategoryResponse struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// ExchangeResponse represents the response structure for exchanges from CoinGecko API
type ExchangeResponse struct {
	ID                          string   `json:"id"`
	Name                        string   `json:"name"`
	YearEstablished             *int     `json:"year_established"`
	Country                     *string  `json:"country"`
	Description                 *string  `json:"description"`
	URL                         *string  `json:"url"`
	Image                       *string  `json:"image"`
	HasTradingIncentive         *bool    `json:"has_trading_incentive"`
	TrustScore                  *int     `json:"trust_score"`
	TrustScoreRank              *int     `json:"trust_score_rank"`
	TradeVolume24hBTC           *float64 `json:"trade_volume_24h_btc"`
	TradeVolume24hBTCNormalized *float64 `json:"trade_volume_24h_btc_normalized"`
}

// GetAssetPlatforms fetches all asset platforms from CoinGecko API
func (c *CoinGeckoClient) GetAssetPlatforms(ctx context.Context) ([]domain.AssetPlatform, error) {
	url := fmt.Sprintf("%s/asset_platforms", c.baseURL)

	logger.GetLogger().WithField("url", url).Info("Fetching asset platforms from CoinGecko API")

	var platforms []domain.AssetPlatform
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		platforms, lastErr = c.fetchAssetPlatforms(ctx, url)
		if lastErr == nil {
			break
		}

		logger.GetLogger().WithError(lastErr).WithField("attempt", attempt).Warn("API request failed")
	}

	if lastErr != nil {
		logger.GetLogger().WithError(lastErr).Error("Failed to fetch asset platforms after all retry attempts")
		return nil, fmt.Errorf("failed to fetch asset platforms: %w", lastErr)
	}

	logger.GetLogger().WithField("count", len(platforms)).Info("Successfully fetched asset platforms from CoinGecko API")
	return platforms, nil
}

// fetchAssetPlatforms performs the actual HTTP request
func (c *CoinGeckoClient) fetchAssetPlatforms(ctx context.Context, url string) ([]domain.AssetPlatform, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cgoffline/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiPlatforms []AssetPlatformResponse
	if err := json.Unmarshal(body, &apiPlatforms); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to domain models
	platforms := make([]domain.AssetPlatform, len(apiPlatforms))
	for i, apiPlatform := range apiPlatforms {
		platforms[i] = domain.AssetPlatform{
			ID:              apiPlatform.ID,
			ChainIdentifier: apiPlatform.ChainIdentifier,
			Name:            apiPlatform.Name,
			ShortName:       apiPlatform.ShortName,
			NativeCoinID:    apiPlatform.NativeCoinID,
		}
	}

	return platforms, nil
}

// GetCoinCategories fetches all coin categories from CoinGecko API
func (c *CoinGeckoClient) GetCoinCategories(ctx context.Context) ([]domain.CoinCategory, error) {
	url := fmt.Sprintf("%s/coins/categories/list", c.baseURL)

	logger.GetLogger().WithField("url", url).Info("Fetching coin categories from CoinGecko API")

	var categories []domain.CoinCategory
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		categories, lastErr = c.fetchCoinCategories(ctx, url)
		if lastErr == nil {
			break
		}

		logger.GetLogger().WithError(lastErr).WithField("attempt", attempt).Warn("API request failed")
	}

	if lastErr != nil {
		logger.GetLogger().WithError(lastErr).Error("Failed to fetch coin categories after all retry attempts")
		return nil, fmt.Errorf("failed to fetch coin categories: %w", lastErr)
	}

	logger.GetLogger().WithField("count", len(categories)).Info("Successfully fetched coin categories from CoinGecko API")
	return categories, nil
}

// fetchCoinCategories performs the actual HTTP request for categories
func (c *CoinGeckoClient) fetchCoinCategories(ctx context.Context, url string) ([]domain.CoinCategory, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cgoffline/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiCategories []CoinCategoryResponse
	if err := json.Unmarshal(body, &apiCategories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to domain models
	categories := make([]domain.CoinCategory, len(apiCategories))
	for i, apiCategory := range apiCategories {
		categories[i] = domain.CoinCategory{
			CoingeckoID: apiCategory.CategoryID,
			Name:        apiCategory.Name,
		}
	}

	return categories, nil
}

// GetExchanges fetches all exchanges from CoinGecko API
func (c *CoinGeckoClient) GetExchanges(ctx context.Context) ([]domain.Exchange, error) {
	url := fmt.Sprintf("%s/exchanges", c.baseURL)

	logger.GetLogger().WithField("url", url).Info("Fetching exchanges from CoinGecko API")

	var exchanges []domain.Exchange
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		exchanges, lastErr = c.fetchExchanges(ctx, url)
		if lastErr == nil {
			break
		}

		logger.GetLogger().WithError(lastErr).WithField("attempt", attempt).Warn("API request failed")
	}

	if lastErr != nil {
		logger.GetLogger().WithError(lastErr).Error("Failed to fetch exchanges after all retry attempts")
		return nil, fmt.Errorf("failed to fetch exchanges: %w", lastErr)
	}

	logger.GetLogger().WithField("count", len(exchanges)).Info("Successfully fetched exchanges from CoinGecko API")
	return exchanges, nil
}

// fetchExchanges performs the actual HTTP request for exchanges
func (c *CoinGeckoClient) fetchExchanges(ctx context.Context, url string) ([]domain.Exchange, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cgoffline/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiExchanges []ExchangeResponse
	if err := json.Unmarshal(body, &apiExchanges); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to domain models
	exchanges := make([]domain.Exchange, len(apiExchanges))
	for i, apiExchange := range apiExchanges {
		exchanges[i] = domain.Exchange{
			CoingeckoID:                 apiExchange.ID,
			Name:                        apiExchange.Name,
			YearEstablished:             apiExchange.YearEstablished,
			Country:                     apiExchange.Country,
			Description:                 apiExchange.Description,
			URL:                         apiExchange.URL,
			Image:                       apiExchange.Image,
			HasTradingIncentive:         apiExchange.HasTradingIncentive,
			TrustScore:                  apiExchange.TrustScore,
			TrustScoreRank:              apiExchange.TrustScoreRank,
			TradeVolume24hBTC:           apiExchange.TradeVolume24hBTC,
			TradeVolume24hBTCNormalized: apiExchange.TradeVolume24hBTCNormalized,
		}
	}

	return exchanges, nil
}

// HealthCheck checks if the CoinGecko API is accessible
func (c *CoinGeckoClient) HealthCheck(ctx context.Context) error {
	url := fmt.Sprintf("%s/ping", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create health check request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check failed with status %d", resp.StatusCode)
	}

	return nil
}

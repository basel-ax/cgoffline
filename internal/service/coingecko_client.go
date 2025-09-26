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

// CoinResponse represents the response structure for coins from CoinGecko API
type CoinResponse struct {
	ID                           string     `json:"id"`
	Symbol                       string     `json:"symbol"`
	Name                         string     `json:"name"`
	Image                        *string    `json:"image"`
	CurrentPrice                 *float64   `json:"current_price"`
	MarketCap                    *float64   `json:"market_cap"`
	MarketCapRank                *int       `json:"market_cap_rank"`
	FullyDilutedValuation        *float64   `json:"fully_diluted_valuation"`
	TotalVolume                  *float64   `json:"total_volume"`
	High24h                      *float64   `json:"high_24h"`
	Low24h                       *float64   `json:"low_24h"`
	PriceChange24h               *float64   `json:"price_change_24h"`
	PriceChangePercentage24h     *float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           *float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h *float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            *float64   `json:"circulating_supply"`
	TotalSupply                  *float64   `json:"total_supply"`
	MaxSupply                    *float64   `json:"max_supply"`
	Ath                          *float64   `json:"ath"`
	AthChangePercentage          *float64   `json:"ath_change_percentage"`
	AthDate                      *time.Time `json:"ath_date"`
	Atl                          *float64   `json:"atl"`
	AtlChangePercentage          *float64   `json:"atl_change_percentage"`
	AtlDate                      *time.Time `json:"atl_date"`
	LastUpdated                  *time.Time `json:"last_updated"`
}

// CoinMarketDataResponse represents the response structure for coin market data from CoinGecko API
type CoinMarketDataResponse struct {
	ID                           string     `json:"id"`
	Symbol                       string     `json:"symbol"`
	Name                         string     `json:"name"`
	Image                        *string    `json:"image"`
	CurrentPrice                 *float64   `json:"current_price"`
	MarketCap                    *float64   `json:"market_cap"`
	MarketCapRank                *int       `json:"market_cap_rank"`
	FullyDilutedValuation        *float64   `json:"fully_diluted_valuation"`
	TotalVolume                  *float64   `json:"total_volume"`
	High24h                      *float64   `json:"high_24h"`
	Low24h                       *float64   `json:"low_24h"`
	PriceChange24h               *float64   `json:"price_change_24h"`
	PriceChangePercentage24h     *float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           *float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h *float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            *float64   `json:"circulating_supply"`
	TotalSupply                  *float64   `json:"total_supply"`
	MaxSupply                    *float64   `json:"max_supply"`
	Ath                          *float64   `json:"ath"`
	AthChangePercentage          *float64   `json:"ath_change_percentage"`
	AthDate                      *time.Time `json:"ath_date"`
	Atl                          *float64   `json:"atl"`
	AtlChangePercentage          *float64   `json:"atl_change_percentage"`
	AtlDate                      *time.Time `json:"atl_date"`
	LastUpdated                  *time.Time `json:"last_updated"`
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

// GetCoins fetches coins with market data from CoinGecko API
func (c *CoinGeckoClient) GetCoins(ctx context.Context, page int, perPage int) ([]domain.Coin, error) {
	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=%d&page=%d&sparkline=false",
		c.baseURL, perPage, page)

	logger.GetLogger().WithFields(map[string]interface{}{
		"url":      url,
		"page":     page,
		"per_page": perPage,
	}).Info("Fetching coins from CoinGecko API")

	var coins []domain.Coin
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		coins, lastErr = c.fetchCoins(ctx, url)
		if lastErr == nil {
			break
		}

		logger.GetLogger().WithError(lastErr).WithField("attempt", attempt).Warn("API request failed")
	}

	if lastErr != nil {
		logger.GetLogger().WithError(lastErr).Error("Failed to fetch coins after all retry attempts")
		return nil, fmt.Errorf("failed to fetch coins: %w", lastErr)
	}

	logger.GetLogger().WithField("count", len(coins)).Info("Successfully fetched coins from CoinGecko API")
	return coins, nil
}

// fetchCoins performs the actual HTTP request for coins
func (c *CoinGeckoClient) fetchCoins(ctx context.Context, url string) ([]domain.Coin, error) {
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

	var apiCoins []CoinResponse
	if err := json.Unmarshal(body, &apiCoins); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to domain models
	coins := make([]domain.Coin, len(apiCoins))
	for i, apiCoin := range apiCoins {
		coins[i] = domain.Coin{
			CoingeckoID:                  apiCoin.ID,
			Symbol:                       apiCoin.Symbol,
			Name:                         apiCoin.Name,
			Image:                        apiCoin.Image,
			CurrentPrice:                 apiCoin.CurrentPrice,
			MarketCap:                    apiCoin.MarketCap,
			MarketCapRank:                apiCoin.MarketCapRank,
			FullyDilutedValuation:        apiCoin.FullyDilutedValuation,
			TotalVolume:                  apiCoin.TotalVolume,
			High24h:                      apiCoin.High24h,
			Low24h:                       apiCoin.Low24h,
			PriceChange24h:               apiCoin.PriceChange24h,
			PriceChangePercentage24h:     apiCoin.PriceChangePercentage24h,
			MarketCapChange24h:           apiCoin.MarketCapChange24h,
			MarketCapChangePercentage24h: apiCoin.MarketCapChangePercentage24h,
			CirculatingSupply:            apiCoin.CirculatingSupply,
			TotalSupply:                  apiCoin.TotalSupply,
			MaxSupply:                    apiCoin.MaxSupply,
			Ath:                          apiCoin.Ath,
			AthChangePercentage:          apiCoin.AthChangePercentage,
			AthDate:                      apiCoin.AthDate,
			Atl:                          apiCoin.Atl,
			AtlChangePercentage:          apiCoin.AtlChangePercentage,
			AtlDate:                      apiCoin.AtlDate,
			LastUpdated:                  apiCoin.LastUpdated,
		}
	}

	return coins, nil
}

// GetCoinMarketData fetches market data for a specific coin from CoinGecko API
func (c *CoinGeckoClient) GetCoinMarketData(ctx context.Context, coinID string) ([]domain.CoinMarketData, error) {
	url := fmt.Sprintf("%s/coins/%s/tickers", c.baseURL, coinID)

	logger.GetLogger().WithFields(map[string]interface{}{
		"url":     url,
		"coin_id": coinID,
	}).Info("Fetching coin market data from CoinGecko API")

	var marketData []domain.CoinMarketData
	var lastErr error

	// Retry logic
	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		marketData, lastErr = c.fetchCoinMarketData(ctx, url, coinID)
		if lastErr == nil {
			break
		}

		logger.GetLogger().WithError(lastErr).WithField("attempt", attempt).Warn("API request failed")
	}

	if lastErr != nil {
		logger.GetLogger().WithError(lastErr).Error("Failed to fetch coin market data after all retry attempts")
		return nil, fmt.Errorf("failed to fetch coin market data: %w", lastErr)
	}

	logger.GetLogger().WithField("count", len(marketData)).Info("Successfully fetched coin market data from CoinGecko API")
	return marketData, nil
}

// GetCoinDataByID fetches full coin data by ID from CoinGecko API (/coins/{id})
func (c *CoinGeckoClient) GetCoinDataByID(ctx context.Context, coinID string) (map[string]any, error) {
	url := fmt.Sprintf("%s/coins/%s", c.baseURL, coinID)

	logger.GetLogger().WithFields(map[string]interface{}{
		"url":     url,
		"coin_id": coinID,
	}).Info("Fetching coin data by ID from CoinGecko API")

	var lastErr error
	var payload map[string]any

	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "cgoffline/1.0")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to execute request: %w", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			lastErr = fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("failed to read response body: %w", err)
			continue
		}

		if err := json.Unmarshal(body, &payload); err != nil {
			lastErr = fmt.Errorf("failed to unmarshal response: %w", err)
			continue
		}

		return payload, nil
	}

	return nil, fmt.Errorf("failed to fetch coin data by id: %w", lastErr)
}

// GetCoinTickers fetches coin tickers by ID from CoinGecko API (/coins/{id}/tickers)
// Reference: https://docs.coingecko.com/v3.0.1/reference/coins-id-tickers
func (c *CoinGeckoClient) GetCoinTickers(ctx context.Context, coinID string, page int) (map[string]any, error) {
	url := fmt.Sprintf("%s/coins/%s/tickers?page=%d", c.baseURL, coinID, page)

	logger.GetLogger().WithFields(map[string]interface{}{
		"url":     url,
		"coin_id": coinID,
		"page":    page,
	}).Info("Fetching coin tickers by ID from CoinGecko API")

	var lastErr error
	var payload map[string]any

	for attempt := 0; attempt <= c.retryCount; attempt++ {
		if attempt > 0 {
			logger.GetLogger().WithField("attempt", attempt).Info("Retrying API request")
			time.Sleep(c.retryDelay)
		}

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", "cgoffline/1.0")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to execute request: %w", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			lastErr = fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("failed to read response body: %w", err)
			continue
		}

		if err := json.Unmarshal(body, &payload); err != nil {
			lastErr = fmt.Errorf("failed to unmarshal response: %w", err)
			continue
		}

		return payload, nil
	}

	return nil, fmt.Errorf("failed to fetch coin tickers: %w", lastErr)
}

// fetchCoinMarketData performs the actual HTTP request for coin market data
func (c *CoinGeckoClient) fetchCoinMarketData(ctx context.Context, url string, coinID string) ([]domain.CoinMarketData, error) {
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

	var response struct {
		Tickers []struct {
			Base   string `json:"base"`
			Target string `json:"target"`
			Market struct {
				Name       string `json:"name"`
				Identifier string `json:"identifier"`
			} `json:"market"`
			Last          *float64 `json:"last"`
			Volume        *float64 `json:"volume"`
			ConvertedLast struct {
				USD *float64 `json:"usd"`
			} `json:"converted_last"`
			ConvertedVolume struct {
				USD *float64 `json:"usd"`
			} `json:"converted_volume"`
			TrustScore             string     `json:"trust_score"`
			BidAskSpreadPercentage *float64   `json:"bid_ask_spread_percentage"`
			Timestamp              *time.Time `json:"timestamp"`
			LastTradedAt           *time.Time `json:"last_traded_at"`
			LastFetchAt            *time.Time `json:"last_fetch_at"`
			IsAnomaly              bool       `json:"is_anomaly"`
			IsStale                bool       `json:"is_stale"`
			TradeURL               *string    `json:"trade_url"`
			TokenInfoURL           *string    `json:"token_info_url"`
			CoinID                 string     `json:"coin_id"`
		} `json:"tickers"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Convert API response to domain models
	marketData := make([]domain.CoinMarketData, 0, len(response.Tickers))
	for _, ticker := range response.Tickers {
		// Skip if no price data
		if ticker.Last == nil {
			continue
		}

		marketData = append(marketData, domain.CoinMarketData{
			// Note: CoinID and ExchangeID will be set by the service layer
			// after resolving the coin and exchange IDs
			Price:            ticker.Last,
			Volume24h:        ticker.Volume,
			VolumePercentage: nil, // Not available in this API response
			LastUpdated:      ticker.LastTradedAt,
		})
	}

	return marketData, nil
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

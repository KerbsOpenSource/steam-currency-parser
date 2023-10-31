// The steam package implements requests to the site to obtain information
package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MarketPrice struct {
	Success     bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	Volume      string `json:"volume"`
	MedianPrice string `json:"median_price"`
}

func GetMarketPrice(appID string, currency string, marketHashName string) (*MarketPrice, error) {
	url := fmt.Sprintf("https://steamcommunity.com/market/priceoverview/?appid=%s&currency=%s&market_hash_name=%s", appID, currency, marketHashName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var marketPrice MarketPrice
	err = json.Unmarshal(body, &marketPrice)
	if err != nil {
		return nil, err
	}

	if !marketPrice.Success {
		return nil, err
	}

	return &marketPrice, nil
}

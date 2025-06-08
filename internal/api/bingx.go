package api

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"bingx-bot/internal/config"

	"github.com/go-resty/resty/v2"
)

type BingXClient struct {
	client  *resty.Client
	apiKey  string
	secret  string
	baseUrl string
}

func NewBingXClient(cfg *config.Config) *BingXClient {
	return &BingXClient{
		client:  resty.New(),
		apiKey:  cfg.ApiKey,
		secret:  cfg.SecretKey,
		baseUrl: cfg.BaseUrl,
	}
}

type priceData struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
	BidPrice  string `json:"bidPrice"`
	AskPrice  string `json:"askPrice"`
}

type PriceResponse struct {
	Code      int       `json:"code"`
	Timestamp int64     `json:"timestamp"`
	Data      priceData `json:"data"`
}

func (b *BingXClient) GetPrice(symbol string) (float64, error) {
	endpoint := "/openApi/swap/v2/quote/price"

	resp, err := b.client.R().
		SetQueryParams(map[string]string{
			"symbol": symbol, // ejemplo: BTC-USDT
		}).
		SetHeader("X-BX-APIKEY", b.apiKey).
		SetHeader("Content-Type", "application/json").
		Get(b.baseUrl + endpoint)

	if err != nil {
		return 0, err
	}

	log.Println("API Response:", resp.String())

	// Estructura de respuesta esperada
	type priceData struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	}
	type priceResponse struct {
		Code int       `json:"code"`
		Msg  string    `json:"msg"`
		Data priceData `json:"data"`
	}

	var pr priceResponse
	if err := json.Unmarshal(resp.Body(), &pr); err != nil {
		return 0, err
	}

	if pr.Code != 0 {
		return 0, fmt.Errorf("error en respuesta API: %s", pr.Msg)
	}

	price, err := strconv.ParseFloat(pr.Data.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

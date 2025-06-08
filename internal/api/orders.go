package api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"bingx-bot/internal/model"
)

func (b *BingXClient) PlaceOrder(order model.Order) (string, error) {
	endpoint := "/openApi/swap/v2/trade/order"

	params := map[string]string{
		"symbol":   order.Symbol,
		"side":     order.Side, // "BUY" o "SELL"
		"quantity": strconv.FormatFloat(order.Quantity, 'f', -1, 64),
	}

	resp, err := b.client.R().
		SetQueryParams(params).
		SetHeader("X-BX-APIKEY", b.apiKey).
		SetHeader("Content-Type", "application/json").
		Post(b.baseUrl + endpoint)

	if err != nil {
		return "", err
	}

	var result struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			OrderId string `json:"orderId"`
		} `json:"data"`
	}

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", err
	}

	if result.Code != 0 {
		return "", fmt.Errorf("error al colocar orden: %s", result.Msg)
	}

	return result.Data.OrderId, nil
}

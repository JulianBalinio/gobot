package model

type PriceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type Order struct {
	Symbol   string
	Side     string // "BUY" o "SELL"
	Quantity float64
}

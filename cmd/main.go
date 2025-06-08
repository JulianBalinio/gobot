package main

import (
	"log"

	"bingx-bot/internal/api"
	"bingx-bot/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	client := api.NewBingXClient(cfg)

	ticker := "DOT-USDT"

	price, err := client.GetPrice(ticker)
	if err != nil {
		log.Fatal("Error obteniendo precio: ", err)
	}

	log.Printf("Precio %s (futuros): %.2f", ticker, price)
}

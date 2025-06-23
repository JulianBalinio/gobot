package main

import (
	"net/http"

	"bingx-bot/internal/api"
	"bingx-bot/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	client := api.NewBingXClient(cfg)

	router := gin.Default()

	router.GET("/price", func(c *gin.Context) {
		ticker := c.Query("ticker")
		if ticker == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'ticker' query parameter"})
			return
		}

		price, err := client.GetPrice(ticker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ticker": ticker,
			"price":  price,
		})
	})

	router.Run(":8080")
}

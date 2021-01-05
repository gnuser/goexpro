package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex/binance"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

var (
	config      *Config
	binanceSpot *binance.Binance
)

func init() {
	config = ParseConfigFile("config.yml")

	// init binance exchange
	binanceSpot = InitBinance(config.Proxy)
}

func main() {
	log := logrus.New()
	// hooks, config,...

	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.GET("")
	r.Run("127.0.0.1:8080")
}

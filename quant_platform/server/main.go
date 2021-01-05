package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/builder"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

var (
	config *Config
	api    goex.API
	log    *logrus.Logger
)

func init() {
	log := logrus.New()
	config = ParseConfigFile("config.yml")

	// init binance exchange
	api = builder.NewAPIBuilder2(&config.HttpClientConfig).APIKey(config.Binance.ApiKey).APISecretkey(config.Binance.SecretKey).Build(goex.BINANCE)
	log.Printf("[%v] exchange api created", api.GetExchangeName())
}

func main() {
	// hooks, config,...

	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.GET("")
	r.Run("127.0.0.1:8080")
}

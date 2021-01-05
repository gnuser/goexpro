package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex/quant_platform/server/api"
	"github.com/nntaoli-project/goex/quant_platform/server/config"
	"github.com/nntaoli-project/goex/quant_platform/server/handler"
	"github.com/nntaoli-project/goex/quant_platform/server/logger"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

var (
	platformConfig *config.PlatformConfig

	log *logrus.Logger
)

func init() {
	log = logger.InitLogger()

	platformConfig = config.ParseConfigFile("config.yml")
	api.InitApi(platformConfig)

}

func main() {
	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("pong"))
	})

	r.GET("/balance", handler.GetBalance)
	r.Run("127.0.0.1:8080")
}

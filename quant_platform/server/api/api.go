package api

import (
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/builder"
	"github.com/nntaoli-project/goex/quant_platform/server/config"
	"log"
)

var (
	RestApi   goex.API
	WsApi     goex.SpotWsApi
	WalletApi goex.WalletApi
)

func InitApi(config *config.PlatformConfig) {
	// init binance exchange
	httpClientConfig := &builder.HttpClientConfig{
		HttpTimeout:  config.HttpClientConfig.Timeout,
		MaxIdleConns: config.HttpClientConfig.MaxIdleConns}
	httpClientConfig.SetProxyUrl(config.HttpClientConfig.Proxy)
	builder := builder.NewAPIBuilder2(httpClientConfig).APIKey(config.Binance.ApiKey).APISecretkey(config.Binance.SecretKey)
	RestApi = builder.Build(goex.BINANCE)
	log.Printf("[%v] exchange api created", RestApi.GetExchangeName())

	var err error
	WsApi, err = builder.BuildSpotWs(goex.BINANCE)
	if err != nil {
		log.Fatalf("build spot ws api failed: %v", err.Error())
	}

	WalletApi, err = builder.BuildWallet(goex.BINANCE)
	if err != nil {
		log.Fatalf("build wallet api failed: %v", err.Error())
	}
}

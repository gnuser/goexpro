package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nntaoli-project/goex/quant_platform/server/api"
	"strings"
)

func GetBalance(c *gin.Context) {
	assets := make(map[string]float64)
	account, err := api.WalletApi.GetAccount()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	for currency, subAccount := range account.SubAccounts {
		if subAccount.Amount > 0 {
			assets[strings.ToLower(currency.Symbol)] = subAccount.Amount
		}
	}

	c.JSON(200, gin.H{
		"data": assets,
	})
}

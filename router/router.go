package router

import (
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/service"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	rootGroup := r.Group(global.Config.Server.ContextPath)
	{
		rootGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		priceGroup := rootGroup.Group("/priceSync")
		{
			priceGroup.POST("/restart", service.RestartPriceServer)
		}

		erc20Group := rootGroup.Group("/erc20")
		{
			erc20Group.POST("/balanceOf", service.BalanceOf)
			erc20Group.POST("/approve", service.Approve)
		}

		dexGroup := rootGroup.Group("/dex")
		{
			utilGroup := dexGroup.Group("/util")
			utilGroup.POST("/getPair", service.GetPair)
			utilGroup.GET("/getPairInfo", service.GetPairInfo)
			dexGroup.POST("/createPair", service.CreatePair)

			priceGroup := dexGroup.Group("/price")
			priceGroup.POST("/quote", service.Quote)
			priceGroup.POST("/getAmountOut", service.GetAmountOut)
			priceGroup.POST("/getAmountIn", service.GetAmountIn)

			txGroup := dexGroup.Group("/tx")
			txGroup.POST("/addLiquidity", service.AddLiquidity)
			txGroup.POST("/swapExactTokenForTokens", service.SwapExactTokenForTokens)
			txGroup.POST("/removeLiquidity", service.RemoveLiquidity)
		}

	}

}

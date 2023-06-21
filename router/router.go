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

		priceGroup := rootGroup.Group("/price")
		{
			priceGroup.POST("/restart", service.RestartPriceServer)
		}

		erc20Group := rootGroup.Group("/erc20")
		{
			erc20Group.GET("/")
		}

		dexGroup := rootGroup.Group("/dex")
		{
			dexGroup.POST("/getPair", service.GetPair)
		}

	}

}

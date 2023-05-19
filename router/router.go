package router

import (
	"github.com/gin-gonic/gin"
	"interest-arbitrage/global"
	"interest-arbitrage/service"
	"net/http"
)

func RegisterRouter(r *gin.Engine) {
	routerGroup := r.Group(global.Config.Server.ContextPath)
	{
		routerGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		priceGroup := routerGroup.Group("/price")
		priceGroup.POST("/restart", service.RestartPriceServer)

	}

}

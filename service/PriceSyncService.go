package service

import (
	"github.com/gin-gonic/gin"
	"interest-arbitrage/server"
	"log"
	"net/http"
)

func RestartPriceServer(c *gin.Context) {
	server.UsablePriceServer.Restart()
	log.Println("priceServer restart")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

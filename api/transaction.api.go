package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {

	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/transaction", getTransaction)
		productAPI.POST("/transaction", createTransaction)

	}
}

func getTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "get Transaction"})
}

func createTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "create Transaction"})

}

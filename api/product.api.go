package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProduct(router *gin.Engine) {

	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.POST("/product", createProduct)

	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "get product"})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "create product"})

}

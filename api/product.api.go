package api

import (
	"fmt"
	"jwt/database"
	"jwt/interceptor"
	"jwt/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)



func SetupProduct(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.JWTVerify, getProduct)
		productAPI.POST("/product", createProduct)
		productAPI.PUT("/product", editProduct)

	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "get product",
		"username": c.GetString("jwt_username"),
	})
}

// ----------check file already exist?
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func saveImage(image *multipart.FileHeader, product *model.Product, c *gin.Context) {
	if(image != nil){
		wd, _ := os.Getwd()
		product.Image = image.Filename
		extension := filepath.Ext(image.Filename)
		fileName := fmt.Sprintf("%d%s", product.ID, extension)
		filePath := fmt.Sprintf("%s/uploaded/images/%s", wd, fileName)

		if fileExists(filePath){
			os.Remove(filePath)
		}

		c.SaveUploadedFile(image, filePath)
		database.GetDB().Model(&product).Update("image", fileName)
	}
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.CreatedAt = time.Now()
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	database.GetDB().Create(&product)

	image, _ := c.FormFile(("image"))
	saveImage(image, &product, c)
	
	c.JSON(http.StatusOK, gin.H{
		"status": "product created",
		"result": product,
	})
}

func editProduct(c *gin.Context) {
	var product model.Product
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)
	product.ID = uint(id)
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

	database.GetDB().Save(&product)

	image, _ := c.FormFile("image")
	saveImage(image, &product, c)

	c.JSON(http.StatusOK, gin.H{"status": product})
}

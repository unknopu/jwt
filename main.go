package main

import (
	"jwt/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Static("/images", "./uploaded/images")

	api.Setup(router)
	router.Run()
}
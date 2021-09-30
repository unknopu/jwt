package api

import (
	"jwt/database"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	database.SetUpDB()

	SetupAuthenAPI(router)
	SetupProduct(router)
	SetupTransactionAPI(router)

}

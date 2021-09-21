package api

import (
	"github.com/gin-gonic/gin"
	"jwt/database"
)


func Setup(router *gin.Engine){
	database.SetUpDB()

	SetupAuthenAPI(router)
	SetupProduct(router)
	SetupTransactionAPI(router)
	
}

package api

import "github.com/gin-gonic/gin"


func Setup(router *gin.Engine){
	// db.SetupDB()

	SetupAuthenAPI(router)
	// SetupProduct(router)
	// SetupTransactionAPI(router)
	
}

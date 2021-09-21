package api

import (
	"jwt/database"
	"jwt/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {

	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "login"})
}

func register(c *gin.Context) {
	var user model.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()

		if err := database.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "not ok", "error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "ok", "data": user})
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind data"})
	}

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

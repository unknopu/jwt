package interceptor

import (
	"fmt"
	"jwt/model"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = "qazwsxedcrfvtgb"

func JWTSign(payload model.User) string {
	atClaims := jwt.MapClaims{}

	// ------------ Payload
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secret))
	return token
}

func JWTVerify(c *gin.Context) {

	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		c.Set("jwt_username", claims["username"])
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "not ok", "error": err, "message": "invalid token"})
		c.Abort()

	}
}

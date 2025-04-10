package middleware

import (
	"api/models"
	"api/store"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTAuth(handlerFunc gin.HandlerFunc, s store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-jwt-token")

		token, err := ValidateJWTtoken(tokenString)

		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if !token.Valid || claims["expires_at"] == nil || claims["user_id"] == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			return
		}

		if time.Now().Unix() > int64(claims["expires_at"].(float64)) {
			c.JSON(http.StatusForbidden, gin.H{"error": "token expired"})
			return
		}

		id, err := strconv.Atoi(c.Param("user_id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id needs to be an int"})
			return
		}

		if id != int(claims["user_id"].(float64)) {

			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}

		handlerFunc(c)
	}

}

func CreateJWTtoken(user *models.User) (string, error) {

	claims := jwt.MapClaims{
		"expires_at": time.Now().Add(time.Hour * 1).Unix(),
		"user_id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJWTtoken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}

	return token, err
}

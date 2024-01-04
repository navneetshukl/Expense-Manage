package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// ! Authenticate will work as middleware to validate the JWT token
func Authenticate(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	tokenString, err := c.Cookie("Authorization")
	secret := os.Getenv("SECRET")

	if err != nil {
		log.Println("Error in Getting the Tokenstring from cookie ",err)
		c.Redirect(http.StatusSeeOther, "/user/login")
		return

	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Redirect(http.StatusSeeOther, "/user/login")
			return

		}
		email := claims["sub"].(string)
		if email == "" {
			c.Redirect(http.StatusSeeOther, "/user/login")
			return

		}
		c.Set("user", email)
		c.Next()

	} else {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return

	}

}

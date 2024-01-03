package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// ! Authenticate will work as middleware to validate the JWT token
func Authenticate(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		log.Println("Line number 19")
		//c.Redirect(http.StatusMovedPermanently, "/user/login")
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Line number 35")
			//c.Redirect(http.StatusMovedPermanently, "/user/login")
			c.AbortWithStatus(http.StatusUnauthorized)

			return

		}
		email := claims["sub"]
		if email == "" {
			log.Println("Line number 42")
			//c.Redirect(http.StatusMovedPermanently, "/user/login")
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}
		c.Set("user", email)
		c.Next()

	} else {
		//c.Redirect(http.StatusMovedPermanently, "/user/login")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	fmt.Println("Inside the Validate route")
	email, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "This is validated page",
		"Email":   email,
	})
}

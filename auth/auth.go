package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// !Register will create a user in our database
func Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	limit := c.PostForm("limit")
	password := c.PostForm("password")

	mp := map[string]interface{}{
		"name": name, "email": email, "limit": limit, "password": password,
	}

	c.JSON(http.StatusOK, mp)

}

func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "register.page.tmpl", nil)

}

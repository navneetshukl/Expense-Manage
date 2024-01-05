package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	categories:=[]string{"grocerry"}
	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"category": categories,
	})
}

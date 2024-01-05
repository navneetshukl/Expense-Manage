package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	categories := []string{"grocerry", "medicine"}
	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"category": categories,
	})
}

// ! Add function will show the page for adding the expense
func Add(c *gin.Context) {
	param := c.Param("param")

	c.HTML(http.StatusOK, "addexpense.page.tmpl", gin.H{
		"param": param,
	})
}

// ! AddPrice function will enter the expense for particular category
func AddPrice(c *gin.Context) {
	price := c.PostForm("price")

	c.JSON(http.StatusOK, gin.H{
		"Price": price,
	})
}

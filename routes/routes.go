package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/helpers"
)

func Home(c *gin.Context) {

	//? Get the expense for today for every category

	email, ok := c.Get("user")

	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}

	expenses, err := helpers.GetExpenses(email.(string))

	if err != nil {
		log.Println("Error in getting the expenses of all the category ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please retry again",
		})
	}
	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"Grocerry":          expenses[0],
		"Transportation":    expenses[1],
		"HouseMaintanance": expenses[2],
		"Medicine":          expenses[3],
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
func AddExpenseForToday(c *gin.Context) {
	param := c.Param("param")

	fmt.Println("Param from AddExpenseForToday is ", param)
	price := c.PostForm("price")
	email, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}
	err := helpers.AddExpenseForCategory(param, email.(string), price)
	if err != nil {
		log.Println("Error in Inserting to database ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some Error Occur.Please retry again",
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/expense")

}

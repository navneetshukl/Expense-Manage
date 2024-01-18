package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/services"
)

func Home(c *gin.Context) {

	//? Get the expense for today for every category

	email, ok := c.Get("user")

	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}

	expenses, err := helpers.GetExpenses(email.(string))
	total := 0

	for _, val := range expenses {
		total += val
	}

	if err != nil {
		log.Println("Error in getting the expenses of all the category ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please retry again",
		})
	}

	limit, err := helpers.GetMaxLimit(email.(string))
	if err != nil {
		log.Println("Error in getting the maximum limit of particular user ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please try again",
		})
		return
	}
	c.HTML(http.StatusOK, "home.page.tmpl", gin.H{
		"Grocerry":         expenses[0],
		"Transportation":   expenses[1],
		"HouseMaintanance": expenses[2],
		"Medicine":         expenses[3],
		"Limit":            limit,
		"Total":            total,
	})

	expLimit := (limit * 90) / 100

	if total >= expLimit {
		_ = services.SendMail(email.(string))

	}

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

func ExtraInformationHTMLPage(c *gin.Context) {
	c.HTML(http.StatusOK, "prevexpense.page.tmpl", nil)
}

func GetPreviousExpense(c *gin.Context) {
	email, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}

	month := c.PostForm("month")
	category := c.PostForm("category")

	data, err := helpers.GetExpenseForAnyMonth(month, category, email.(string))
	if err != nil {
		log.Println("Error in getting the given month data ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please retry again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"expense": data,
	})
}

package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/models"
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

	for _, val := range expenses {
		fmt.Println(val)
	}

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
func AddExpenseForToday(c *gin.Context) {
	param := c.Param("param")

	fmt.Println("Param from AddExpenseForToday is ", param)
	price := c.PostForm("price")
	email, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in connecting to database ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please retry again",
		})
	}
	grocData := models.Grocery{
		Email:   email.(string),
		Date:    time.Now(),
		Expense: price,
	}

	result := db.Create(&grocData)

	if result.Error != nil {
		log.Println("Error in inserting to database ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some Error Occur.Please retry after sometime",
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/expense")

}

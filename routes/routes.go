package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/models"
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
func AddExpenseForToday(c *gin.Context) {
	price := c.PostForm("price")
	email, ok := c.Get("user")
	if !ok {
		c.Redirect(http.StatusSeeOther, "/user/login")
		return
	}

	groc := models.Grocery{}
	db, err := database.ConnectToDatabase()

	if err != nil {
		log.Println("Error in connecting to Database in AddExpenseForToday ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please try again",
		})

	} else {
		createdAt := time.Now()
		result := db.Where("email = ? AND created_at = ?", email, createdAt).First(&groc)

		if result.Error != nil {
			log.Println("Unable to get the data from database ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Some error occured.Please retry after Again",
			})
		}

		//? Get the price for day

		todayExpense, err := helpers.StringToInt(groc.Expense)
		currExpense, err := helpers.StringToInt(price)

		if err != nil {
			log.Println("Error in conversion of string to int ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Some error occured.Please retry",
			})
		}

		//? Add this price
		totalExpense := (todayExpense) + (currExpense)
		newExpense := helpers.IntToString(totalExpense)

		//? Add the updated price for the day
		conditions := map[string]interface{}{
			"email":      email,
			"created_at": createdAt,
		}

		result = db.Model(&groc).Where(conditions).Updates(models.Grocery{Expense: newExpense})

		if result.Error != nil {
			log.Println("Error in updating the Expense ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "SOme Error occured.Please retry again",
			})
		}
	}
}

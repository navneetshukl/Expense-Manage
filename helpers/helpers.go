package helpers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
)

// ! StringToInt this will convert string to int
func StringToInt(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	val, err := strconv.Atoi(str)

	if err != nil {
		return 0, err
	} else {
		return val, nil
	}
}

// ! IntToString convert int to string
func IntToString(val int) string {
	str := strconv.Itoa(val)
	return str
}

// ! GetExpenses function will get the expense for every category for particular month
func GetExpenses(email string) ([]int, error) {
	expenses := []int{}

	db, err := database.ConnectToDatabase()
	if err != nil {

		return expenses, err
	}

	todayDate := time.Now()
	yesterdayData := todayDate.AddDate(0, 0, -1)

	var grocData []models.Grocery
	var medData []models.Medicine
	var homeData []models.HomeMaintanance
	var transData []models.Transportation

	grocExpense, medExpense, homeExpense, transExpense := 0, 0, 0, 0

	res1 := db.Where("email=? and date>=? and date<=?", email, yesterdayData, todayDate).Find(&grocData)
	res2 := db.Where("email=? and date>=? and date<=?", email, yesterdayData, todayDate).Find(&medData)
	res3 := db.Where("email=? and date>=? and date<=?", email, yesterdayData, todayDate).Find(&homeData)
	res4 := db.Where("email=? and date>=? and date<=?", email, yesterdayData, todayDate).Find(&transData)

	if res1.Error != nil || res2.Error != nil || res3.Error != nil || res4.Error != nil {
		log.Println("Error from Groccery table is : ", res1.Error)
		log.Println("Error from Medicine table is : ", res2.Error)
		log.Println("Error from HomeMaintanance table is : ", res3.Error)
		log.Println("Error from Transportation table is : ", res4.Error)

		return expenses, fmt.Errorf("error in Getting the data from database ")
	}

	if len(grocData) > 0 {
		for _, val := range grocData {
			exp, _ := StringToInt(val.Expense)
			grocExpense += exp
		}
		expenses = append(expenses, grocExpense)
	}
	if len(medData) > 0 {
		for _, val := range medData {
			exp, _ := StringToInt(val.Expense)
			medExpense += exp
		}
		expenses = append(expenses, medExpense)
	}
	if len(homeData) > 0 {
		for _, val := range homeData {
			exp, _ := StringToInt(val.Expense)
			homeExpense += exp
		}
		expenses = append(expenses, homeExpense)
	}
	if len(transData) > 0 {
		for _, val := range transData {
			exp, _ := StringToInt(val.Expense)
			transExpense += exp
		}
		expenses = append(expenses, transExpense)
	}
	return expenses, nil

}

// ! AddExpenseForCategory function will enter the expense for particular category to database
func AddExpenseForCategory(param, email, price string) error {
	db, err := database.ConnectToDatabase()
	if err != nil {
		return err
	}

	if param == "grocerry" {

		groc := models.Grocery{
			Email:   email,
			Expense: price,
			Date:    time.Now(),
		}

		res := db.Create(&groc)
		if res.Error != nil {
			return res.Error
		}

	} else if param == "medicine" {

		med := models.Medicine{
			Email:   email,
			Expense: price,
			Date:    time.Now(),
		}
		res := db.Create(&med)
		if res.Error != nil {
			return res.Error
		}

	} else if param == "transportation" {

		trans := models.Transportation{
			Email:   email,
			Expense: price,
			Date:    time.Now(),
		}
		res := db.Create(&trans)
		if res.Error != nil {
			return res.Error
		}

	} else if param == "house-maintainance" {

		home := models.HomeMaintanance{
			Email:   email,
			Expense: price,
			Date:    time.Now(),
		}

		res := db.Create(&home)
		if res.Error != nil {
			return res.Error
		}

	}
	return nil
}

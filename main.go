package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/auth"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/middleware"
	"github.com/navneetshukl/routes"
)

func init() {
	database.MigrateDatabase()
	/* data,_:=redis.GetUserDetailsFromRedis()
	fmt.Println("Data is ",data["email"]) */

}
func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.POST("/", auth.Register)
	router.GET("/", auth.Home)

	router.GET("/user/login", auth.LoginPage)
	router.POST("/user/login", auth.Login)
	router.GET("/user/signup", auth.Signup)

	router.GET("/expense", middleware.Authenticate, routes.Home)
	router.GET("/:param/add", middleware.Authenticate, routes.Add)
	router.POST("/:param/add", middleware.Authenticate, routes.AddExpenseForToday)

	router.GET("/more", middleware.Authenticate, routes.ExtraInformationHTMLPage)
	router.POST("/expense/history", middleware.Authenticate, routes.GetPreviousExpense)

	router.GET("/expense/pdf", middleware.Authenticate, routes.ShowPdf)

	router.Run()

}

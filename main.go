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
	//router.POST("/:param/add", middleware.Authenticate, routes.AddPrice)

	router.Run()

}

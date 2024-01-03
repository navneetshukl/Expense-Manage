package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/auth"
	"github.com/navneetshukl/database"
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
	router.GET("/valid",auth.IsValid)

	router.Run()

}

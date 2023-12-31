package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/routes"
)

func init() {
	database.MigrateDatabase()
}
func main() {
	router := gin.Default()

	router.GET("/",routes.Home)

	router.Run()

}

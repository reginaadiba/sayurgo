package main

import (
	"sayurgo/config"
	"sayurgo/models"
	"sayurgo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8090")
}

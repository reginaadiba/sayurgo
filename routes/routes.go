package routes

import (
	"sayurgo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/products", controllers.GetProducts)
		api.POST("/products", controllers.CreateProduct)
	}

}

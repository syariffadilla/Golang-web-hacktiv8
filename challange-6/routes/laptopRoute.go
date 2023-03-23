package routes

import (
	"challange-4/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/laptop", controllers.CreateLaptop)

	router.PUT("/laptop/:LaptopID", controllers.UpdateLaptop)

	router.GET("/laptop/:LaptopID", controllers.GetLaptop)

	router.DELETE("/laptop/:LaptopID", controllers.DelateLaptop)

	return router

}

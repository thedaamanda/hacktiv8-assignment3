package router

import (
	"assignment3/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.GET("/api/get-weather", controllers.GetData)
	router.GET("/weather", controllers.WebPage)

	return router
}

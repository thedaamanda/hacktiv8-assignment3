package controllers

import (
	"assignment3/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {
	data, _ := services.ReadJSON(services.GetFiles())

	ctx.JSON(http.StatusOK, data)
}

func WebPage(ctx *gin.Context) {
	ctx.File("web/index.html")
}

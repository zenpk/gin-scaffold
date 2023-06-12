package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/gin-scaffold/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter())         // CORS

	router.GET("/", ginHandler.Temp.TempHandler)

	// login required URL
	auth := router.Group("/")
	auth.Use(middleware.RequireLogin())
	{
	}
}

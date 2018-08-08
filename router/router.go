package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine  {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"message": "pong",
		})
	});

	return router
}

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine,mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())

	g.Use(mw...)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"code": 404,
			"message": "not fond",
		})
	})

	//u := g.Group("v1")

	return g
}

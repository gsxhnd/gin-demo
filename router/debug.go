package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/gin-demo/controller"
)

func DebugRoute(g *gin.Engine) *gin.Engine {
	debug := g.Group("/debug")
	debug.GET("", controller.Debug)
	return g
}

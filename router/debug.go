package router

import (
	"GinDemo/controller/debugHandler"
	"github.com/gin-gonic/gin"
)

func DebugRoute(g *gin.Engine) *gin.Engine {
	debug := g.Group("/debug")
	debug.GET("", debugHandler.Debug)
	return g
}

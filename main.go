package main

import (
	"gin-demo/config"
	"gin-demo/middleware"
	"gin-demo/model"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title Swagger API
// @version 1.0
// @description This is a  server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	// init flag
	_ = config.FlagInit()

	// init database
	model.DB.Init()
	defer model.DB.Close()

	// run a  gin server without default middleware
	gin.SetMode(viper.GetString("runMode"))
	g := gin.New()

	// load route and middleware
	router.Load(
		g,
		middleware.RequestLogger(),
		middleware.Cors(),
	)

	// listen and serve on 0.0.0.0:8080
	_ = g.Run()
}

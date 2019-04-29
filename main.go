package main

import (
	"APIServer/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// init flag
	_ = config.FlagInit()

	// Create the Gin engine.
	//g := gin.New()

	// Routes.
	//router.Load(g)

	//fmt.Print(viper.GetStringMap("log")["writers"])

	//run a default gin server for test
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

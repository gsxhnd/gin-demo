package main

import (
	"APIServer/config"
)

func main() {
	// init flag
	_ = config.FlagInit()

	// init config
	//if err := config.Init(*cfg); err != nil {
	//	panic(err)
	//}

	// Create the Gin engine.
	//g := gin.New()

	// Routes.
	//router.Load(g)
}

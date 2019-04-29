package config

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "conf/config.yaml", "api server config file path")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func FlagInit() error {
	pflag.Parse()
	if *version {
		fmt.Println("APIServer version ", 1)
	}

	fmt.Println(*cfg)

	return nil
}

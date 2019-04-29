package config

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "conf/config.yaml", "api server config file path")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

var (
	_version_   string
	_branch_    string
	_commitId_  string
	_buildTime_ string
)

// init flag
func FlagInit() error {
	pflag.Parse()
	if *version {
		fmt.Printf("Version: %s, Branch: %s, Build: %s, Build time: %s\n",
			_version_, _branch_, _commitId_, _buildTime_)
	}

	// init config
	if err := ConfInit(*cfg); err != nil {
		panic(err)
	}

	return nil
}

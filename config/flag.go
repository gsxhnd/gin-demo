package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"runtime"
)

var (
	cfg     = pflag.StringP("config", "c", "conf/config.yaml", "api server config file path")
	version = pflag.BoolP("version", "v", false, "show version info.")
	help    = pflag.BoolP("help", "h", false, "show help info.")
)

var (
	gitTag       string = ""
	gitCommit    string
	gitTreeState string
	buildDate    string
)

type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// init flag
func FlagInit() error {
	pflag.Parse()
	if *version {
		v := Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return errors.New("stop")
	}

	if *help {
		pflag.PrintDefaults()
		return errors.New("stop")
	}

	// init config
	if err := ConfInit(*cfg); err != nil {
		panic(err)
	}
	return nil
}

func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

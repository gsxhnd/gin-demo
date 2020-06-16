package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/gin-demo/logger"
	"github.com/gsxhnd/gin-demo/middleware"
	"github.com/gsxhnd/gin-demo/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	cobra.OnInitialize(initConfig)
	runCmd.PersistentFlags().String("name", "bangumi", "server name")
	runCmd.PersistentFlags().String("run_mode", "debug", "run mode")
	runCmd.PersistentFlags().String("addr", ":8080", "listen addr")
	runCmd.PersistentFlags().StringVar(&cfgFile, "", "listen addr", "")
	_ = viper.BindPFlag("name", runCmd.PersistentFlags().Lookup("name"))
	_ = viper.BindPFlag("run_mode", runCmd.PersistentFlags().Lookup("runMode"))
	_ = viper.BindPFlag("addr", runCmd.PersistentFlags().Lookup("addr"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")                  // 设置配置文件格式为YAML
		if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
			panic(err)
		}
	}
}

var cfgFile string
var runCmd = &cobra.Command{
	Use:   "run",
	Short: `run server`,
	Long:  `run server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run mode: ", viper.GetString("run_mode"))
		fmt.Println("name: ", viper.GetString("name"))
		fmt.Println("addr: ", viper.GetString("addr"))
		logger.Init(viper.GetString("run_mode"), viper.GetString("name"))

		// run a  gin server without default middleware
		gin.SetMode(viper.GetString("run_mode"))
		g := gin.New()

		// load route and middleware
		router.Load(
			g,
			//middleware.RequestLogger(),
			middleware.ReqZapLogger(viper.GetString("name")),
			middleware.Cors(),
			gin.Recovery(),
		)

		// listen and serve on 0.0.0.0:8080
		err := g.Run(viper.GetString("addr"))
		if err != nil {
			logger.Error("server run error", zap.Error(err))
			panic(err)
		}
	},
}

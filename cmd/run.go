package cmd

import (
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
	runCmd.PersistentFlags().String("runMode", "debug", "run mode")
	runCmd.PersistentFlags().String("addr", ":8080", "listen addr")
	_ = viper.BindPFlag("name", runCmd.PersistentFlags().Lookup("name"))
	_ = viper.BindPFlag("runMode", runCmd.PersistentFlags().Lookup("runMode"))
	_ = viper.BindPFlag("addr", runCmd.PersistentFlags().Lookup("addr"))
}

func initConfig() {
	//if cfgFile != "" {
	//	viper.SetConfigFile(cfgFile)
	//	viper.SetConfigType("yaml")                  // 设置配置文件格式为YAML
	//	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
	//		panic(err)
	//	}
	//}
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: `run server`,
	Long:  `run server`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Init(viper.GetString("runMode"), viper.GetString("name"))

		// run a  gin server without default middleware
		gin.SetMode(viper.GetString("runMode"))
		g := gin.New()

		// load route and middleware
		router.Load(
			g,
			//middleware.RequestLogger(),
			middleware.ReqZapLogger("gin-demo"),
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

func init() {
	rootCmd.AddCommand(runCmd)
}

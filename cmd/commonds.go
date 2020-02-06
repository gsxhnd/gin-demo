package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	//cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "A generator for Cobra based Applications",
		Long:  `server is a CLI  applications.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Println(cmd)
			fmt.Println(args)
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"jeanfo_mix/cmd/subcmd"
	"jeanfo_mix/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

//	@title			JEANFO_MIX_API
//	@version		1.0
//	@description	This is a WEB server for JEANFO_MIX_API.
//	@termsOfService	jeanfo.cn

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @contact.name	Jeanfo Peng
// @contact.url	http://jeanfo.cn
// @contact.email	jeanf@qq.com
func main() {
	rootCmd := &cobra.Command{
		Use:   "jeanfo mix app",
		Short: "jeanfo mix app -- short msg",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if configPath != "" {
				config.SetConfigPath(configPath)
			}
		},
	}
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to config file")

	rootCmd.AddCommand(subcmd.WebCmd)
	rootCmd.AddCommand(subcmd.GetKickUserCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

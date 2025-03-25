package subcmd

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"

	"github.com/spf13/cobra"
)

func RunWeb() {
	cfg := config.GetConfig()
	listen_on := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	db := model.GetDB()

	r := router.SetupRouter(db)
	r.Run(listen_on)
}

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		RunWeb()
	},
}

package subcmd

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"
	"jeanfo_mix/util/log_util"

	"github.com/spf13/cobra"
)

func RunWeb() {
	cfg := config.GetConfig()
	listen_on := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	db := model.GetDB()

	log_util.Info("start web server on %s...", listen_on)
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

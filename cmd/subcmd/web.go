package subcmd

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/model"
	"jeanfo_mix/internal/router"
)

func RunWeb() {
	cfg := config.GetConfig()
	listen_on := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	db := model.GetDB()

	r := router.SetupRouter(db)
	r.Run(listen_on)
}

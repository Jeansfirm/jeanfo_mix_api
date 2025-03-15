package subcmd

import (
	"fmt"
	"jeanfo_mix/config"
	"jeanfo_mix/internal/router"

	"gorm.io/gorm"
)

func RunWeb(db *gorm.DB) {
	cfg := config.GetConfig()
	listen_on := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	r := router.SetupRouter(db)
	r.Run(listen_on)
}

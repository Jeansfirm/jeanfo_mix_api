package main

import (
	"flag"
	"jeanfo_mix/cmd/subcmd"
	"jeanfo_mix/config"
	"jeanfo_mix/util/log_util"
)

var (
	configPath string
	mode       string
	execCmd    string
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
	flag.StringVar(&configPath, "c", "", "指定配置文件路径")
	flag.StringVar(&configPath, "config", "", "指定配置文件路径")
	flag.StringVar(&mode, "m", "web", "启动后台服务类型: web")
	flag.StringVar(&execCmd, "e", "", "执行子命令")

	flag.Parse()

	if configPath != "" {
		config.SetConfigPath(configPath)
	}

	err := log_util.Init()
	if err != nil {
		panic("Log init fail: " + err.Error())
	}

	//执行子命令
	if execCmd != "" {
		switch execCmd {

		}
		return
	}

	//后台服务模式
	switch mode {
	case "web":
		subcmd.RunWeb()
	default:
		panic("必须指定合法的启动模式或者单独执行子命令")
	}

}

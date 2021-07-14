package main

import (
	"fmt"

	"github.com/zhs007/block7"
	block7serv "github.com/zhs007/block7/httpserv"
	"go.uber.org/zap"
)

func main() {
	cfg, err := block7serv.LoadConfig("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig error! %v", err)

		return
	}

	block7.InitLogger("block7.serv", block7.Version,
		cfg.LogLevel, true, cfg.LogPath)

	service, err := block7serv.NewBasicServ(cfg)
	if err != nil {
		block7.Info("NewBasicServ",
			zap.String("addr", cfg.BindAddr),
			zap.Error(err))

		return
	}

	serv := block7serv.NewServ(service)

	block7.Info("init serv ...",
		zap.String("addr", cfg.BindAddr))

	serv.Start()

	block7.Info("serv end.")
}

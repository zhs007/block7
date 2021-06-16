package main

import (
	"github.com/zhs007/block7"
	block7serv "github.com/zhs007/block7/httpserv"
	"go.uber.org/zap"
)

func main() {
	block7.InitLogger("block7.serv", block7.Version,
		"debug", true, "./logs")

	cfg := &block7serv.Config{
		BindAddr:    "0.0.0.0:3723",
		IsDebugMode: false,
	}

	service := block7serv.NewBasicServ()
	serv := block7serv.NewServ(service, cfg)

	block7.Info("init serv ...",
		zap.String("addr", cfg.BindAddr))

	serv.Start()

	block7.Info("serv end.")
}

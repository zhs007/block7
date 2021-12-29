package main

import (
	"fmt"

	"github.com/zhs007/block7"
	block7serv "github.com/zhs007/block7/httpserv"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func main() {
	cfg, err := block7serv.LoadConfig("./cfg/config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig error! %v", err)

		return
	}

	goutils.InitLogger("block7.serv", block7.Version,
		cfg.LogLevel, true, cfg.LogPath)

	service, err := block7serv.NewBasicServ(cfg)
	if err != nil {
		goutils.Info("NewBasicServ",
			zap.String("addr", cfg.BindAddr),
			zap.Error(err))

		return
	}

	retLogin, err := service.Login(&block7serv.LoginParams{
		Game:      "test002",
		Platform:  "test",
		ABVersion: "b",
	})
	if err != nil {
		goutils.Info("Login",
			zap.Error(err))

		return
	}

	retMission, err := service.Mission(&block7serv.MissionParams{
		UserHash:  retLogin.UserHash,
		MissionID: 8,
	})
	if err != nil {
		goutils.Info("Mission",
			zap.Error(err))

		return
	}

	goutils.Info("serv end.",
		goutils.JSON("ret", retMission))
}

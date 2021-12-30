package main

import (
	"fmt"
	"io/ioutil"

	"github.com/zhs007/block7"
	block7serv "github.com/zhs007/block7/httpserv"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config - configuration
type Config struct {
	ABVersion  string `yaml:"abVersion"`
	StartStage int    `yaml:"startStage"`
	EndStage   int    `yaml:"endStage"`
}

// LoadConfig - load config
func LoadConfig(fn string) (*Config, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func foreachStages(service *block7serv.BasicServ, abVersion string, startMissionID int, endMissionID int) {
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

	for curMissionID := startMissionID; curMissionID <= endMissionID; curMissionID++ {
		retMission, err := service.Mission(&block7serv.MissionParams{
			UserHash:  retLogin.UserHash,
			MissionID: curMissionID,
		})
		if err != nil {
			goutils.Info("Mission",
				zap.Int("missionid", curMissionID),
				zap.Error(err))

			return
		}

		goutils.Info("serv end.",
			zap.Int("missionid", curMissionID),
			goutils.JSON("ret", retMission))
	}
}

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

	cfg2, err := LoadConfig("./cfg/checkmaps.yaml")
	if err != nil {
		fmt.Printf("LoadConfig checkmaps error! %v", err)

		return
	}

	foreachStages(service, cfg2.ABVersion, cfg2.StartStage, cfg2.EndStage)

	// retLogin, err := service.Login(&block7serv.LoginParams{
	// 	Game:      "test002",
	// 	Platform:  "test",
	// 	ABVersion: "b",
	// })
	// if err != nil {
	// 	goutils.Info("Login",
	// 		zap.Error(err))

	// 	return
	// }

	// retMission, err := service.Mission(&block7serv.MissionParams{
	// 	UserHash:  retLogin.UserHash,
	// 	MissionID: 8,
	// })
	// if err != nil {
	// 	goutils.Info("Mission",
	// 		zap.Error(err))

	// 	return
	// }

	// goutils.Info("serv end.",
	// 	goutils.JSON("ret", retMission))
}

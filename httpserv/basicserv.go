package block7serv

import (
	"context"

	"github.com/zhs007/block7"
	"go.uber.org/zap"
)

// BasicServ - basic server
type BasicServ struct {
	UserDB  *block7.UserDB
	StageDB *block7.StageDB
	cfg     *Config
}

func NewBasicServ(cfg *Config) (*BasicServ, error) {
	userdb, err := block7.NewUserDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		block7.Error("NewBasicServ:NewUserDB",
			zap.Error(err))

		return nil, err
	}

	stagedb, err := block7.NewStageDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		block7.Error("NewBasicServ:NewStageDB",
			zap.Error(err))

		return nil, err
	}

	return &BasicServ{
		UserDB:  userdb,
		StageDB: stagedb,
		cfg:     cfg,
	}, nil
}

// GetConfig - get configuation
func (serv *BasicServ) GetConfig() *Config {
	return serv.cfg
}

// Login - login
func (serv *BasicServ) Login(params *LoginParams) (*LoginResult, error) {
	udi := LoginParams2PB(params)
	if udi.UserHash == "" {
		ui, err := serv.UserDB.NewUser(context.Background(), udi)
		if err != nil {
			block7.Error("BasicServ.Login:NewUser",
				zap.Error(err))

			return nil, err
		}

		return &LoginResult{
			UserID:   ui.UserID,
			UserHash: udi.UserHash,
		}, nil
	}

	ui, err := serv.UserDB.UpdUserDeviceInfo(context.Background(), udi)
	if err != nil {
		block7.Error("BasicServ.Login:UpdUserDeviceInfo",
			zap.Error(err))

		return nil, err
	}

	return &LoginResult{
		UserID:   ui.UserID,
		UserHash: udi.UserHash,
	}, nil
}

// Mission - get mission
func (serv *BasicServ) Mission(params *MissionParams) (*MissionResult, error) {
	stage, err := block7.LoadStage("./cfg/level_0100.json")
	if err != nil {
		block7.Error("LoadStage",
			zap.Error(err))

		return nil, err
	}

	rng := block7.NewRngNormal()

	scene, err := block7.NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, block7.DefaultMaxBlockNums)
	if err != nil {
		block7.Error("NewScene",
			zap.Error(err))

		return nil, err
	}

	scene.IsOutputScene = true

	return &MissionResult{
		Scene:       scene,
		MissionHash: block7.GenHashCode(16),
	}, nil
}

// MissionData - upload mission data
func (serv *BasicServ) MissionData(params *MissionDataParams) (*MissionDataResult, error) {
	return &MissionDataResult{UserLevel: 100}, nil
}

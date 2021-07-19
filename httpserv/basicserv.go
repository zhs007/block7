package block7serv

import (
	"context"

	"github.com/zhs007/block7"
	"go.uber.org/zap"
)

// BasicServ - basic server
type BasicServ struct {
	UserDB    *block7.UserDB
	StageDB   *block7.StageDB
	HistoryDB *block7.HistoryDB
	cfg       *Config
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

	historydb, err := block7.NewHistoryDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		block7.Error("NewBasicServ:NewHistoryDB",
			zap.Error(err))

		return nil, err
	}

	return &BasicServ{
		UserDB:    userdb,
		StageDB:   stagedb,
		HistoryDB: historydb,
		cfg:       cfg,
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
	if params.UserHash == "" {
		block7.Error("BasicServ.Mission",
			zap.Error(ErrInvalidUserHash))

		return nil, ErrInvalidUserHash
	}

	uid, err := serv.UserDB.GetUserID(context.Background(), params.UserHash)
	if err != nil {
		block7.Error("BasicServ.Mission:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		block7.Error("BasicServ.Mission:GetUserID",
			zap.Int64("uid", uid),
			zap.String("userhash", params.UserHash),
			zap.Error(ErrInvalidUserHash))

		return nil, err
	}

	stage, err := block7.LoadStage("./cfg/level_0100.json")
	if err != nil {
		block7.Error("BasicServ.Mission:LoadStage",
			zap.Error(err))

		return nil, err
	}

	rng := block7.NewRngNormal()

	scene, err := block7.NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, block7.DefaultMaxBlockNums)
	if err != nil {
		block7.Error("BasicServ.Mission:NewScene",
			zap.Error(err))

		return nil, err
	}

	scene.MapID = "0100"
	scene.IsOutputScene = true

	pbScene, err := serv.StageDB.SaveStage(context.Background(), scene)
	if err != nil {
		block7.Error("BasicServ.Mission:SaveStage",
			zap.Error(err))

		return nil, err
	}
	// mhash :=

	return &MissionResult{
		Scene:   scene,
		SceneID: pbScene.SceneID,
	}, nil
}

// MissionData - upload mission data
func (serv *BasicServ) MissionData(params *MissionDataParams) (*MissionDataResult, error) {
	if params.UserHash == "" {
		block7.Error("BasicServ.MissionData",
			zap.Error(ErrInvalidUserHash))

		return nil, ErrInvalidUserHash
	}

	uid, err := serv.UserDB.GetUserID(context.Background(), params.UserHash)
	if err != nil {
		block7.Error("BasicServ.MissionData:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		block7.Error("BasicServ.MissionData:GetUserID",
			zap.Int64("uid", uid),
			zap.String("userhash", params.UserHash),
			zap.Error(ErrInvalidUserHash))

		return nil, err
	}

	pbscene, err := serv.StageDB.GetStage(context.Background(), params.SceneID)
	if err != nil {
		block7.Error("BasicServ.MissionData:GetStage",
			zap.Error(err))

		return nil, err
	}

	scene := block7.NewSceneFromPB(pbscene)
	MissionDataParams2Scene(scene, params)

	pbscene1, err := serv.HistoryDB.SaveHistory(context.Background(), scene)
	if err != nil {
		block7.Error("BasicServ.MissionData:SaveHistory",
			zap.Error(err))

		return nil, err
	}

	if serv.cfg.IsDebugMode {
		block7.Debug("BasicServ.MissionData",
			block7.JSON("history", pbscene1))
	}

	return &MissionDataResult{UserLevel: 100}, nil
}

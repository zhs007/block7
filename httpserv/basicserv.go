package block7serv

import (
	"context"
	"fmt"

	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// BasicServ - basic server
type BasicServ struct {
	UserDB    *block7.UserDB
	StageDB   *block7.StageDB
	HistoryDB *block7.HistoryDB
	LevelMgr  *block7game.LevelMgr
	cfg       *Config
}

func NewBasicServ(cfg *Config) (*BasicServ, error) {
	userdb, err := block7.NewUserDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		goutils.Error("NewBasicServ:NewUserDB",
			zap.Error(err))

		return nil, err
	}

	stagedb, err := block7.NewStageDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		goutils.Error("NewBasicServ:NewStageDB",
			zap.Error(err))

		return nil, err
	}

	historydb, err := block7.NewHistoryDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		goutils.Error("NewBasicServ:NewHistoryDB",
			zap.Error(err))

		return nil, err
	}

	levelmgr := block7game.NewLevelMgr()
	err = levelmgr.LoadLevel("./gamedata/level.json")
	if err != nil {
		goutils.Error("NewBasicServ:LoadLevel",
			zap.Error(err))

		return nil, err
	}

	return &BasicServ{
		UserDB:    userdb,
		StageDB:   stagedb,
		HistoryDB: historydb,
		LevelMgr:  levelmgr,
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
			goutils.Error("BasicServ.Login:NewUser",
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
		goutils.Error("BasicServ.Login:UpdUserDeviceInfo",
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
		goutils.Error("BasicServ.Mission",
			zap.Error(ErrInvalidUserHash))

		return nil, ErrInvalidUserHash
	}

	uid, err := serv.UserDB.GetUserID(context.Background(), params.UserHash)
	if err != nil {
		goutils.Error("BasicServ.Mission:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		goutils.Error("BasicServ.Mission:GetUserID",
			zap.Int64("uid", uid),
			zap.String("userhash", params.UserHash),
			zap.Error(ErrInvalidUserHash))

		return nil, ErrInvalidUserHash
	}

	if params.HistoryID > 0 {
		pbscene, err := serv.HistoryDB.GetHistory(context.Background(), params.HistoryID)
		if err != nil {
			goutils.Error("BasicServ.Mission:StageDB.GetHistory",
				zap.Error(err))

			return nil, err
		}

		if pbscene != nil {
			pbscene1, err := serv.StageDB.GetStage(context.Background(), pbscene.SceneID)
			if err != nil {
				goutils.Error("BasicServ.Mission:StageDB.GetStage",
					zap.Error(err))

				return nil, err
			}

			if pbscene1 != nil {
				pbscene1.History2 = pbscene.History2

				scene, err := block7game.NewSceneFromPB(pbscene1)
				if err != nil {
					goutils.Error("BasicServ.Mission:NewSceneFromPB",
						goutils.JSON("pbscene", pbscene1),
						zap.Error(err))

					return nil, err
				}

				scene.IsOutputScene = true

				scene.ReadyToClient()

				return &MissionResult{
					Scene:   scene,
					SceneID: pbscene1.SceneID,
				}, nil
			}
		}
	}

	if params.SceneID > 0 {
		pbscene, err := serv.StageDB.GetStage(context.Background(), params.SceneID)
		if err != nil {
			goutils.Error("BasicServ.Mission:StageDB.GetStage",
				zap.Error(err))

			return nil, err
		}

		if pbscene != nil {
			scene, err := block7game.NewSceneFromPB(pbscene)
			if err != nil {
				goutils.Error("BasicServ.Mission:NewSceneFromPB",
					goutils.JSON("pbscene", pbscene),
					zap.Error(err))

				return nil, err
			}

			scene.IsOutputScene = true

			scene.ReadyToClient()

			return &MissionResult{
				Scene:   scene,
				SceneID: pbscene.SceneID,
			}, nil
		}
	}

	ld2, isok := serv.LevelMgr.MapLevel[params.MissionID+30000]
	if !isok {
		goutils.Error("BasicServ.Mission:GetUserID",
			zap.Int64("uid", uid),
			zap.Int("missionid", params.MissionID),
			zap.Error(ErrInvalidMissionID))

		return nil, ErrInvalidMissionID
	}

	stage, err := block7game.LoadStage(fmt.Sprintf("./gamedata/map/level_%04d.json", ld2.MapID))
	if err != nil {
		goutils.Error("BasicServ.Mission:LoadStage",
			zap.Error(err))

		return nil, err
	}

	rng := block7game.NewRngNormal()

	scene, err := block7game.NewScene(rng, stage, ld2.GenSymbols(), block7game.DefaultMaxBlockNums, ld2)
	if err != nil {
		goutils.Error("BasicServ.Mission:NewScene",
			zap.Error(err))

		return nil, err
	}

	scene.MapID = ld2.MapID
	scene.IsOutputScene = true

	pbScene, err := serv.StageDB.SaveStage(context.Background(), scene)
	if err != nil {
		goutils.Error("BasicServ.Mission:SaveStage",
			zap.Error(err))

		return nil, err
	}

	scene.ReadyToClient()
	// mhash :=

	return &MissionResult{
		Scene:   scene,
		SceneID: pbScene.SceneID,
	}, nil
}

// MissionData - upload mission data
func (serv *BasicServ) MissionData(params *MissionDataParams) (*MissionDataResult, error) {
	if params.UserHash == "" {
		goutils.Error("BasicServ.MissionData",
			zap.Error(ErrInvalidUserHash))

		return nil, ErrInvalidUserHash
	}

	uid, err := serv.UserDB.GetUserID(context.Background(), params.UserHash)
	if err != nil {
		goutils.Error("BasicServ.MissionData:GetUserID",
			zap.Error(err))

		return nil, err
	}

	if uid <= 0 {
		goutils.Error("BasicServ.MissionData:GetUserID",
			zap.Int64("uid", uid),
			zap.String("userhash", params.UserHash),
			zap.Error(ErrInvalidUserHash))

		return nil, err
	}

	if params.HistoryID > 0 {
		pbscene2, err := serv.HistoryDB.GetHistory(context.Background(), params.HistoryID)
		if err != nil {
			goutils.Error("BasicServ.MissionData:GetHistory",
				zap.Error(err))

			return nil, err
		}

		if pbscene2 != nil {
			arr2, err := goutils.Int32ArrToIntArr2(pbscene2.History2, 4, len(pbscene2.History2)/4)
			if err != nil {
				goutils.Error("BasicServ.MissionData:Int32ArrToIntArr2",
					zap.Error(err))

				return nil, err
			}

			if pbscene2.SceneID == params.SceneID && goutils.IsSameIntArr2Ex2(params.History, arr2, 3) {
				if serv.cfg.IsDebugMode {
					goutils.Debug("BasicServ.MissionData",
						goutils.JSON("history", pbscene2))
				}

				return &MissionDataResult{
					UserLevel: 100,
					HistoryID: pbscene2.HistoryID}, nil
			}

			if serv.cfg.IsDebugMode {
				goutils.Debug("BasicServ.MissionData:cmp",
					zap.Int64("sceneID", pbscene2.SceneID),
					goutils.JSON("arr2", arr2))
			}
		}
	}

	pbscene, err := serv.StageDB.GetStage(context.Background(), params.SceneID)
	if err != nil {
		goutils.Error("BasicServ.MissionData:GetStage",
			zap.Error(err))

		return nil, err
	}

	scene, err := block7game.NewSceneFromPB(pbscene)
	if err != nil {
		goutils.Error("BasicServ.MissionData:NewSceneFromPB",
			zap.Error(err))

		return nil, err
	}

	MissionDataParams2Scene(scene, params)

	pbscene1, err := serv.HistoryDB.SaveHistory(context.Background(), scene)
	if err != nil {
		goutils.Error("BasicServ.MissionData:SaveHistory",
			zap.Error(err))

		return nil, err
	}

	if serv.cfg.IsDebugMode {
		goutils.Debug("BasicServ.MissionData",
			goutils.JSON("history", pbscene1))
	}

	return &MissionDataResult{
		UserLevel: 100,
		HistoryID: pbscene1.HistoryID}, nil
}

// GetUserData - get UserData
func (serv *BasicServ) GetUserData(params *UserDataParams) (*UserDataResult, error) {
	ud, err := serv.UserDB.GetUserData(context.Background(), params.Name, params.Platform)
	if err != nil {
		goutils.Error("BasicServ.GetUserData:GetUserData",
			zap.Error(err))

		return nil, err
	}

	if ud == nil {
		return &UserDataResult{
			Name:     params.Name,
			Platform: params.Platform,
			Version:  0,
		}, nil
	}

	return PB2UserDataResult(ud), nil
}

// UpdUserData - update UserData
func (serv *BasicServ) UpdUserData(ud *UpdUserDataParams, uds *block7.UpdUserDataStatus) (*UpdUserDataResult, error) {
	udpb := UpdUserDataParams2PB(ud, uds)

	oldversion, err := serv.UserDB.UpdUserData(context.Background(), udpb, uds)
	if err != nil {
		goutils.Error("BasicServ.UpdUserData:UpdUserData",
			zap.Error(err))

		return nil, err
	}

	return &UpdUserDataResult{
		OldVersion: oldversion,
		NewVersion: ud.Version,
	}, nil
}

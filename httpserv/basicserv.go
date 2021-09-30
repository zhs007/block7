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
	StatsDB   *block7.StatsDB
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

	statsdb, err := block7.NewStatsDB(cfg.DBPath, "", cfg.DBEngine, userdb, stagedb)
	if err != nil {
		goutils.Error("NewBasicServ:NewStatsDB",
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
		StatsDB:   statsdb,
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
			} else {
				if len(pbscene.InitArr2) > 0 {
					scene, err := block7game.NewSceneFromPB(pbscene)
					if err != nil {
						goutils.Error("BasicServ.Mission:NewSceneFromPB",
							goutils.JSON("pbscene", pbscene),
							zap.Error(err))

						return nil, err
					}

					scene.IsOutputScene = true
					scene.IsFullHistoryData = true

					scene.ReadyToClient()

					return &MissionResult{
						Scene: scene,
					}, nil
				} else {
					goutils.Warn("BasicServ.Mission:StageDB.GetStage:nil",
						zap.Int64("sceneid", pbscene.SceneID))
				}
			}
		} else {
			goutils.Warn("BasicServ.Mission:StageDB.GetHistory:nil",
				zap.Int64("HistoryID", params.HistoryID))
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
		goutils.Error("BasicServ.Mission:MapLevel",
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

	scene.StageID = params.MissionID
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

	var curScene *block7game.Scene

	if params.SceneID > 0 {
		pbscene, err := serv.StageDB.GetStage(context.Background(), params.SceneID)
		if err != nil {
			goutils.Error("BasicServ.MissionData:GetStage",
				zap.Error(err))

			return nil, err
		}

		if pbscene == nil {
			goutils.Error("BasicServ.MissionData:GetStage",
				zap.Int64("sceneID", params.SceneID),
				zap.Error(ErrInvalidSceneID))

			return nil, ErrInvalidSceneID
		}

		scene, err := block7game.NewSceneFromPB(pbscene)
		if err != nil {
			goutils.Error("BasicServ.MissionData:NewSceneFromPB",
				zap.Error(err))

			return nil, err
		}

		curScene = MissionDataParams2Scene(scene, params)
	} else {
		curScene = MissionDataParams2SceneEx(params)
	}

	pbscene1, err := serv.HistoryDB.SaveHistory(context.Background(), curScene)
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
	if ud.UserHash != "" {
		uid, err := serv.UserDB.GetUserID(context.Background(), ud.UserHash)
		if err != nil {
			goutils.Error("BasicServ.UpdUserData:GetUserID",
				zap.Error(err))

			return nil, err
		}

		udpb.UserID = uid
	}

	// goutils.Debug("BasicServ.UpdUserData",
	// 	goutils.JSON("udpb", udpb),
	// 	goutils.JSON("ud", ud),
	// 	goutils.JSON("uds", uds))

	oldversion, newversion, err := serv.UserDB.UpdUserData(context.Background(), udpb, uds)
	if err != nil {
		goutils.Error("BasicServ.UpdUserData:UpdUserData",
			zap.Error(err))

		return nil, err
	}

	return &UpdUserDataResult{
		OldVersion: oldversion,
		NewVersion: newversion,
		Version:    ud.Version,
	}, nil
}

// Stats - statistics
func (serv *BasicServ) Stats(params *StatsParams) (*StatsResult, error) {
	if params.Token != serv.cfg.StatsToken {
		goutils.Error("BasicServ.Stats:checkToken",
			zap.String("params.Token", params.Token),
			zap.String("cfg.StatsToken", serv.cfg.StatsToken),
			zap.Error(ErrInvalidToken))

		return nil, ErrInvalidToken
	}

	stage, err := serv.StageDB.Stats(context.Background())
	if err != nil {
		goutils.Error("BasicServ.Stats:StageDB.Stats",
			zap.Error(err))

		return nil, err
	}

	history, err := serv.HistoryDB.Stats(context.Background())
	if err != nil {
		goutils.Error("BasicServ.Stats:HistoryDB.Stats",
			zap.Error(err))

		return nil, err
	}

	latestUserID, userNums, userDataNums, err := serv.UserDB.Stats(context.Background())
	if err != nil {
		goutils.Error("BasicServ.Stats:UserDB.Stats",
			zap.Error(err))

		return nil, err
	}

	stats, err := serv.StatsDB.Stats(context.Background())
	if err != nil {
		goutils.Error("BasicServ.Stats:StatsDB.Stats",
			zap.Error(err))

		return nil, err
	}

	return &StatsResult{
		LatestUserID: latestUserID,
		UserNums:     userNums,
		UserDataNums: userDataNums,
		Stage:        stage,
		History:      history,
		Stats:        stats,
	}, nil
}

// Start - start
func (serv *BasicServ) Start() {
	serv.StatsDB.Start()
}

// Stop - stop
func (serv *BasicServ) Stop() {
	serv.StatsDB.Stop()
}

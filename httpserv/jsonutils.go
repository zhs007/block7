package block7serv

import (
	"strings"

	"github.com/buger/jsonparser"
	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func parseMissionDataParams(data []byte) (*MissionDataParams, error) {
	userHash, _, err := goutils.GetJsonString(data, "userHash")
	if err != nil {
		goutils.Error("parseMissionDataParams:userHash",
			zap.Error(err))

		return nil, err
	}

	sceneID, _, err := goutils.GetJsonInt(data, "mission")
	if err != nil {
		goutils.Error("parseMissionDataParams:mission",
			zap.Error(err))

		return nil, err
	}

	// if sceneID <= 0 {
	// 	goutils.Error("parseMissionDataParams",
	// 		zap.Error(ErrInvalidSceneID))

	// 	return nil, ErrInvalidSceneID
	// }

	history, err := goutils.GetJsonIntArr2(data, "history")
	if err != nil {
		goutils.Error("parseMissionDataParams:GetJsonIntArr2:history",
			zap.Error(err))

		return nil, err
	}

	historyID, _, err := goutils.GetJsonInt(data, "srcHistory")
	if err != nil {
		goutils.Error("parseMissionDataParams:srcHistory",
			zap.Error(err))

		return nil, err
	}

	rngdata, err := goutils.GetJsonInt64Arr(data, "rngdata")
	if err != nil {
		goutils.Error("parseMissionDataParams:GetJsonInt64Arr:rngdata",
			zap.Error(err))

		return nil, err
	}

	gameState, _, err := goutils.GetJsonInt(data, "gamestate")
	if err != nil {
		goutils.Error("parseMissionDataParams:gamestate",
			zap.Error(err))

		return nil, err
	}

	initArr, err := goutils.GetJsonIntArr3(data, "initArr")
	if err != nil {
		goutils.Error("parseMissionDataParams:GetJsonIntArr3:initArr",
			zap.Error(err))

		return nil, err
	}

	blockNums, _, err := goutils.GetJsonInt(data, "blockNums")
	if err != nil {
		goutils.Error("parseMissionDataParams:blockNums",
			zap.Error(err))

		return nil, err
	}

	stageType, _, err := goutils.GetJsonInt(data, "stageType")
	if err != nil {
		goutils.Error("parseMissionDataParams:stageType",
			zap.Error(err))

		return nil, err
	}

	specialLayers := []*block7game.SpecialLayer{}

	err = goutils.GetJsonObjectArr(data,
		func(value1 []byte, dataType jsonparser.ValueType, offset1 int, err1 error) {
			if err1 != nil {
				goutils.Error("parseMissionDataSpecialLayers:ArrayEach:specialLayers:func",
					zap.Int("offset", offset1),
					zap.Error(err1))

				return
			}

			if dataType == jsonparser.Object {
				sl := &block7game.SpecialLayer{}

				layer, _, err := goutils.GetJsonInt(value1, "layer")
				if err != nil {
					goutils.Error("parseMissionDataSpecialLayers:layer",
						zap.Int("offset", offset1),
						zap.Error(err))

					return
				}

				sl.Layer = int(layer)

				layerType, _, err := goutils.GetJsonInt(value1, "layerType")
				if err != nil {
					goutils.Error("parseMissionDataSpecialLayers:layerType",
						zap.Int("offset", offset1),
						zap.Error(err))
				}

				sl.LayerType = int(layerType)

				pos, err := goutils.GetJsonIntArr2(value1, "pos")
				if err != nil {
					goutils.Error("parseMissionDataSpecialLayers:parseMissionDataSpecialLayersPos",
						zap.Int("offset", offset1),
						zap.Error(err))
				}

				sl.Pos = pos

				specialLayers = append(specialLayers, sl)

				return
			}

			goutils.Error("parseMissionDataSpecialLayers:ArrayEach:specialLayers:func:dataType",
				zap.Int("offset", offset1),
				zap.String("dataType", dataType.String()))
		}, "specialLayers")
	if err != nil {
		goutils.Error("parseMissionDataParams:parseMissionDataSpecialLayers",
			zap.Error(err))

		return nil, err
	}

	firstItem, _, err := goutils.GetJsonInt(data, "firstItem")
	if err != nil {
		goutils.Error("parseMissionDataParams:firstItem",
			zap.Error(err))

		return nil, err
	}

	return &MissionDataParams{
		UserHash:      userHash,
		SceneID:       sceneID,
		History:       history,
		HistoryID:     historyID,
		RngData:       rngdata,
		GameState:     int32(gameState),
		InitArr:       initArr,
		BlockNums:     int(blockNums),
		StageType:     int(stageType),
		SpecialLayers: specialLayers,
		FirstItem:     int(firstItem),
	}, nil
}

func parseUpdUserDataParams(data []byte) (*UpdUserDataParams, *block7.UpdUserDataStatus, error) {
	ud := &UpdUserDataParams{}
	uds := &block7.UpdUserDataStatus{}

	name, _, err := goutils.GetJsonString(data, "name")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:name",
			zap.Error(err))

		return nil, nil, err
	}

	name = strings.TrimSpace(name)

	if name == "" {
		goutils.Error("parseUpdUserDataParams",
			zap.Error(ErrInvalidUserDataName))

		return nil, nil, ErrInvalidUserDataName
	}

	ud.Name = name

	coin, isok, err := goutils.GetJsonInt(data, "coin")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:coin",
			zap.Error(err))

		return nil, nil, err
	}

	uds.HasCoin = isok

	if uds.HasCoin {
		ud.Coin = coin
	}

	level, isok, err := goutils.GetJsonInt(data, "level")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:level",
			zap.Error(err))

		return nil, nil, err
	}

	uds.HasLevel = isok

	if uds.HasLevel {
		ud.Level = int(level)
	}

	hasLevelArr := false
	mapLevelArr := make(map[string]int)
	err = goutils.GetJsonObject(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Number {
			// k, err := goutils.String2Int64(string(key))
			// if err != nil {
			// 	goutils.Error("parseUpdUserDataParams:levelarr:GetJsonObject",
			// 		zap.String("key", string(key)),
			// 		zap.Error(err))

			// 	return err
			// }

			v, err := goutils.String2Int64(string(value))
			if err != nil {
				goutils.Error("parseUpdUserDataParams:levelarr:GetJsonObject",
					zap.String("value", string(value)),
					zap.Error(err))

				return err
			}

			mapLevelArr[string(key)] = int(v)
			hasLevelArr = true
		}

		return nil
	}, "levelarr")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:levelarr",
			zap.Error(err))

		return nil, nil, err
	}

	uds.HasLevelArr = hasLevelArr

	if uds.HasLevelArr {
		ud.LevelArr = mapLevelArr
	}

	hasToolsArr := false
	mapToolsArr := make(map[string]int)
	err = goutils.GetJsonObject(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Number {
			// k, err := goutils.String2Int64(string(key))
			// if err != nil {
			// 	goutils.Error("parseUpdUserDataParams:levelarr:GetJsonObject",
			// 		zap.String("key", string(key)),
			// 		zap.Error(err))

			// 	return err
			// }

			v, err := goutils.String2Int64(string(value))
			if err != nil {
				goutils.Error("parseUpdUserDataParams:levelarr:GetJsonObject",
					zap.String("value", string(value)),
					zap.Error(err))

				return err
			}

			mapToolsArr[string(key)] = int(v)
			hasToolsArr = true
		}

		return nil
	}, "toolsarr")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:toolsarr",
			zap.Error(err))

		return nil, nil, err
	}

	uds.HasToolsArr = hasToolsArr

	if uds.HasToolsArr {
		ud.ToolsArr = mapToolsArr
	}

	homeScene, err := goutils.GetJsonIntArr(data, "homeScene")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:GetJsonIntArr:homeScene",
			zap.Error(err))

		return nil, nil, err
	}

	if len(homeScene) > 0 {
		uds.HasHomeScene = true
		ud.HomeScene = homeScene
	}

	cookings := []*Cooking{}

	err = goutils.GetJsonObjectArr(data,
		func(value1 []byte, dataType jsonparser.ValueType, offset1 int, err1 error) {
			if err1 != nil {
				goutils.Error("parseUpdUserDataParams:GetJsonObjectArr:Cooking:func",
					zap.Int("offset", offset1),
					zap.Error(err1))

				return
			}

			if dataType == jsonparser.Object {
				// Level    int  `json:"level"`
				// Unlock   bool `json:"unlock"`
				// StarNums int  `json:"starnum"`

				ck := &Cooking{}

				level, _, err := goutils.GetJsonInt(value1, "level")
				if err != nil {
					goutils.Error("parseUpdUserDataParams:level",
						zap.Int("offset", offset1),
						zap.Error(err))

					return
				}

				ck.Level = int(level)

				unlock, _, err := goutils.GetJsonBool(value1, "unlock")
				if err != nil {
					goutils.Error("parseUpdUserDataParams:unlock",
						zap.Int("offset", offset1),
						zap.Error(err))

					return
				}

				ck.Unlock = unlock

				starnum, _, err := goutils.GetJsonInt(value1, "starnum")
				if err != nil {
					goutils.Error("parseUpdUserDataParams:starnum",
						zap.Int("offset", offset1),
						zap.Error(err))

					return
				}

				ck.StarNums = int(starnum)

				cookings = append(cookings, ck)

				return
			}

			goutils.Error("parseUpdUserDataParams:ArrayEach:cookings:func:dataType",
				zap.Int("offset", offset1),
				zap.String("dataType", dataType.String()))
		}, "cooking")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:GetJsonObjectArr",
			zap.Error(err))

		return nil, nil, err
	}

	if len(cookings) > 0 {
		uds.HasCooking = true
		ud.Cooking = cookings
	}

	platform, _, err := goutils.GetJsonString(data, "platform")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:platform",
			zap.Error(err))

		return nil, nil, err
	}

	platform = strings.TrimSpace(platform)

	ud.Platform = platform

	version, _, err := goutils.GetJsonInt(data, "version")
	if err != nil {
		goutils.Error("parseUpdUserDataParams:version",
			zap.Error(err))

		return nil, nil, err
	}

	ud.Version = version

	return ud, uds, nil
}

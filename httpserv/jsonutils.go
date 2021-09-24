package block7serv

import (
	"github.com/buger/jsonparser"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// func parseMissionDataHistory(data []byte) ([][]int, error) {
// 	history := [][]int{}

// 	offset, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 		if err != nil {
// 			goutils.Error("parseMissionDataHistory:ArrayEach:history:func",
// 				zap.Int("offset", offset),
// 				zap.Error(err))

// 			return
// 		}

// 		if dataType == jsonparser.Array {
// 			arr0 := []int{}

// 			offset3, err3 := jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err2 error) {
// 				if err2 != nil {
// 					goutils.Error("parseMissionDataHistory:ArrayEach:history:func2",
// 						zap.Int("offset", offset2),
// 						zap.Error(err2))

// 					return
// 				}

// 				if dataType2 == jsonparser.Number {
// 					cv, err5 := jsonparser.GetInt(value2)
// 					if err != nil {
// 						goutils.Error("parseMissionDataHistory:ArrayEach:history:func2:GetInt",
// 							zap.Int("offset", offset2),
// 							zap.Error(err5))

// 						return
// 					}

// 					arr0 = append(arr0, int(cv))

// 					return
// 				}

// 				goutils.Error("parseMissionDataHistory:ArrayEach:history:func2:dataType",
// 					zap.Int("offset", offset),
// 					zap.String("dataType", dataType.String()))
// 			})
// 			if err3 != nil {
// 				goutils.Error("parseMissionDataHistory:ArrayEach:history:func:ArrayEach",
// 					zap.Int("offset", offset3),
// 					zap.Error(err3))

// 				return
// 			}

// 			history = append(history, arr0)

// 			return
// 		}

// 		goutils.Error("parseMissionDataHistory:ArrayEach:history:func:dataType",
// 			zap.Int("offset", offset),
// 			zap.String("dataType", dataType.String()))
// 	}, "history")
// 	if err != nil {
// 		goutils.Error("parseMissionDataHistory:ArrayEach:history",
// 			zap.Int("offset", offset),
// 			zap.Error(err))

// 		return nil, err
// 	}

// 	return history, nil
// }

// func parseMissionDataRngData(data []byte) ([]int64, error) {
// 	rngdata := []int64{}

// 	offset, err := jsonparser.ArrayEach(data, func(value1 []byte, dataType jsonparser.ValueType, offset1 int, err1 error) {
// 		if err1 != nil {
// 			goutils.Error("parseMissionDataRngData:ArrayEach:rngdata:func",
// 				zap.Int("offset", offset1),
// 				zap.Error(err1))

// 			return
// 		}

// 		if dataType == jsonparser.Number {
// 			cv, err2 := jsonparser.GetInt(value1)
// 			if err2 != nil {
// 				goutils.Error("parseMissionDataRngData:ArrayEach:rngdata:func2:GetInt",
// 					zap.Int("offset", offset1),
// 					zap.Error(err2))

// 				return
// 			}

// 			rngdata = append(rngdata, cv)

// 			return
// 		}

// 		goutils.Error("parseMissionDataRngData:ArrayEach:rngdata:func:dataType",
// 			zap.Int("offset", offset1),
// 			zap.String("dataType", dataType.String()))
// 	}, "rngdata")
// 	if err != nil {
// 		goutils.Error("parseMissionDataRngData:ArrayEach",
// 			zap.Int("offset", offset),
// 			zap.Error(err))

// 		return nil, err
// 	}

// 	return rngdata, nil
// }

// func parseMissionDataInitArr(data []byte) ([][][]int, error) {
// 	initArr := [][][]int{}

// 	offset, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 		if err != nil {
// 			goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func",
// 				zap.Int("offset", offset),
// 				zap.Error(err))

// 			return
// 		}

// 		if dataType == jsonparser.Array {
// 			arr0 := [][]int{}

// 			offset3, err3 := jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err2 error) {
// 				if err2 != nil {
// 					goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func2",
// 						zap.Int("offset", offset2),
// 						zap.Error(err2))

// 					return
// 				}

// 				if dataType == jsonparser.Array {
// 					arr1 := []int{}

// 					offset6, err6 := jsonparser.ArrayEach(value, func(value5 []byte, dataType5 jsonparser.ValueType, offset5 int, err5 error) {
// 						if err5 != nil {
// 							goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func3",
// 								zap.Int("offset", offset5),
// 								zap.Error(err5))

// 							return
// 						}

// 						if dataType5 == jsonparser.Number {
// 							cv, err7 := jsonparser.GetInt(value5)
// 							if err != nil {
// 								goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func3:GetInt",
// 									zap.Int("offset", offset5),
// 									zap.Error(err7))

// 								return
// 							}

// 							arr1 = append(arr1, int(cv))

// 							return
// 						}

// 						goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func3:dataType",
// 							zap.Int("offset", offset5),
// 							zap.String("dataType", dataType5.String()))
// 					})
// 					if err6 != nil {
// 						goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func3:ArrayEach",
// 							zap.Int("offset", offset6),
// 							zap.Error(err6))

// 						return
// 					}

// 					arr0 = append(arr0, arr1)
// 					// history = append(history, arr0)

// 					return
// 				}

// 				goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func2:dataType",
// 					zap.Int("offset", offset2),
// 					zap.String("dataType", dataType2.String()))
// 			})
// 			if err3 != nil {
// 				goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func2:ArrayEach",
// 					zap.Int("offset", offset3),
// 					zap.Error(err3))

// 				return
// 			}

// 			initArr = append(initArr, arr0)

// 			return
// 		}

// 		goutils.Error("parseMissionDataInitArr:ArrayEach:initArr:func:dataType",
// 			zap.Int("offset", offset),
// 			zap.String("dataType", dataType.String()))
// 	}, "initArr")
// 	if err != nil {
// 		goutils.Error("parseMissionDataInitArr:ArrayEach:initArr",
// 			zap.Int("offset", offset),
// 			zap.Error(err))

// 		return nil, err
// 	}

// 	return initArr, nil
// }

// func parseMissionDataSpecialLayersPos(data []byte) ([][]int, error) {
// 	pos := [][]int{}

// 	offset, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
// 		if err != nil {
// 			goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func",
// 				zap.Int("offset", offset),
// 				zap.Error(err))

// 			return
// 		}

// 		if dataType == jsonparser.Array {
// 			arr0 := []int{}

// 			offset3, err3 := jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err2 error) {
// 				if err2 != nil {
// 					goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func2",
// 						zap.Int("offset", offset2),
// 						zap.Error(err2))

// 					return
// 				}

// 				if dataType2 == jsonparser.Number {
// 					cv, err5 := jsonparser.GetInt(value2)
// 					if err != nil {
// 						goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func2:GetInt",
// 							zap.Int("offset", offset2),
// 							zap.Error(err5))

// 						return
// 					}

// 					arr0 = append(arr0, int(cv))

// 					return
// 				}

// 				goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func2:dataType",
// 					zap.Int("offset", offset),
// 					zap.String("dataType", dataType.String()))
// 			})
// 			if err3 != nil {
// 				goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func:ArrayEach",
// 					zap.Int("offset", offset3),
// 					zap.Error(err3))

// 				return
// 			}

// 			pos = append(pos, arr0)

// 			return
// 		}

// 		goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos:func:dataType",
// 			zap.Int("offset", offset),
// 			zap.String("dataType", dataType.String()))
// 	}, "pos")
// 	if err != nil {
// 		goutils.Error("parseMissionDataSpecialLayersPos:ArrayEach:pos",
// 			zap.Int("offset", offset),
// 			zap.Error(err))

// 		return nil, err
// 	}

// 	return pos, nil
// }

// func parseMissionDataSpecialLayers(data []byte) ([]*block7game.SpecialLayer, error) {
// 	specialLayers := []*block7game.SpecialLayer{}

// 	offset, err := jsonparser.ArrayEach(data, func(value1 []byte, dataType jsonparser.ValueType, offset1 int, err1 error) {
// 		if err1 != nil {
// 			goutils.Error("parseMissionDataSpecialLayers:ArrayEach:specialLayers:func",
// 				zap.Int("offset", offset1),
// 				zap.Error(err1))

// 			return
// 		}

// 		if dataType == jsonparser.Object {
// 			sl := &block7game.SpecialLayer{}

// 			layer, err := goutils.GetJsonInt(value1, "layer")
// 			if err != nil {
// 				goutils.Error("parseMissionDataSpecialLayers:layer",
// 					zap.Int("offset", offset1),
// 					zap.Error(err))

// 				return
// 			}

// 			sl.Layer = int(layer)

// 			layerType, err := goutils.GetJsonInt(value1, "layerType")
// 			if err != nil {
// 				goutils.Error("parseMissionDataSpecialLayers:layerType",
// 					zap.Int("offset", offset1),
// 					zap.Error(err))
// 			}

// 			sl.LayerType = int(layerType)

// 			pos, err := parseMissionDataSpecialLayersPos(value1)
// 			if err != nil {
// 				goutils.Error("parseMissionDataSpecialLayers:parseMissionDataSpecialLayersPos",
// 					zap.Int("offset", offset1),
// 					zap.Error(err))
// 			}

// 			sl.Pos = pos

// 			specialLayers = append(specialLayers, sl)

// 			return
// 		}

// 		goutils.Error("parseMissionDataSpecialLayers:ArrayEach:specialLayers:func:dataType",
// 			zap.Int("offset", offset1),
// 			zap.String("dataType", dataType.String()))
// 	}, "specialLayers")
// 	if err != nil {
// 		goutils.Error("parseMissionDataSpecialLayers:ArrayEach",
// 			zap.Int("offset", offset),
// 			zap.Error(err))

// 		return nil, err
// 	}

// 	return specialLayers, nil
// }

func parseMissionDataParams(data []byte) (*MissionDataParams, error) {
	// FirstItem     int                        `json:"firstItem"`

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

	err = goutils.GetJsonObjectArr(data, "specialLayers",
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
		})
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

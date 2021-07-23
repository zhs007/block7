package block7serv

import (
	"github.com/buger/jsonparser"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

func parseMissionDataParams(data []byte) (*MissionDataParams, error) {
	userHash, err := block7utils.GetJsonString(data, "userHash")
	if err != nil {
		block7utils.Error("parseMissionDataParams:userHash",
			zap.Error(err))

		return nil, err
	}

	sceneID, err := block7utils.GetJsonInt(data, "mission")
	if err != nil {
		block7utils.Error("parseMissionDataParams:mission",
			zap.Error(err))

		return nil, err
	}

	if sceneID <= 0 {
		block7utils.Error("parseMissionDataParams",
			zap.Error(ErrInvalidSceneID))

		return nil, ErrInvalidSceneID
	}

	history := [][]int{}

	offset, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			block7utils.Error("parseMissionDataParams:ArrayEach:func",
				zap.Int("offset", offset),
				zap.Error(err))

			return
		}

		if dataType == jsonparser.Array {
			arr0 := []int{}

			offset3, err3 := jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err2 error) {
				if err2 != nil {
					block7utils.Error("parseMissionDataParams:ArrayEach:func2",
						zap.Int("offset", offset2),
						zap.Error(err2))

					return
				}

				if dataType == jsonparser.Number {
					cv, err5 := jsonparser.GetInt(value)
					if err != nil {
						block7utils.Error("parseMissionDataParams:ArrayEach:func2:GetInt",
							zap.Int("offset", offset2),
							zap.Error(err5))

						return
					}

					arr0 = append(arr0, int(cv))

					return
				}

				block7utils.Error("parseMissionDataParams:ArrayEach:func2:dataType",
					zap.Int("offset", offset),
					zap.String("dataType", dataType.String()))
			})
			if err3 != nil {
				block7utils.Error("parseMissionDataParams:ArrayEach:func:ArrayEach",
					zap.Int("offset", offset3),
					zap.Error(err3))

				return
			}

			history = append(history, arr0)

			return
		}

		block7utils.Error("parseMissionDataParams:ArrayEach:func:dataType",
			zap.Int("offset", offset),
			zap.String("dataType", dataType.String()))
	}, "history")
	if err != nil {
		block7utils.Error("parseMissionDataParams:ArrayEach",
			zap.Int("offset", offset),
			zap.Error(err))

		return nil, err
	}

	historyID, err := block7utils.GetJsonInt(data, "srcHistory")
	if err != nil {
		block7utils.Error("parseMissionDataParams:srcHistory",
			zap.Error(err))

		return nil, err
	}

	return &MissionDataParams{
		UserHash:  userHash,
		SceneID:   sceneID,
		History:   history,
		HistoryID: historyID,
	}, nil
}

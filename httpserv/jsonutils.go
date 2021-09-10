package block7serv

import (
	"github.com/buger/jsonparser"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func parseMissionDataParams(data []byte) (*MissionDataParams, error) {
	userHash, err := goutils.GetJsonString(data, "userHash")
	if err != nil {
		goutils.Error("parseMissionDataParams:userHash",
			zap.Error(err))

		return nil, err
	}

	sceneID, err := goutils.GetJsonInt(data, "mission")
	if err != nil {
		goutils.Error("parseMissionDataParams:mission",
			zap.Error(err))

		return nil, err
	}

	if sceneID <= 0 {
		goutils.Error("parseMissionDataParams",
			zap.Error(ErrInvalidSceneID))

		return nil, ErrInvalidSceneID
	}

	history := [][]int{}

	offset, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			goutils.Error("parseMissionDataParams:ArrayEach:func",
				zap.Int("offset", offset),
				zap.Error(err))

			return
		}

		if dataType == jsonparser.Array {
			arr0 := []int{}

			offset3, err3 := jsonparser.ArrayEach(value, func(value2 []byte, dataType2 jsonparser.ValueType, offset2 int, err2 error) {
				if err2 != nil {
					goutils.Error("parseMissionDataParams:ArrayEach:func2",
						zap.Int("offset", offset2),
						zap.Error(err2))

					return
				}

				if dataType2 == jsonparser.Number {
					cv, err5 := jsonparser.GetInt(value2)
					if err != nil {
						goutils.Error("parseMissionDataParams:ArrayEach:func2:GetInt",
							zap.Int("offset", offset2),
							zap.Error(err5))

						return
					}

					arr0 = append(arr0, int(cv))

					return
				}

				goutils.Error("parseMissionDataParams:ArrayEach:func2:dataType",
					zap.Int("offset", offset),
					zap.String("dataType", dataType.String()))
			})
			if err3 != nil {
				goutils.Error("parseMissionDataParams:ArrayEach:func:ArrayEach",
					zap.Int("offset", offset3),
					zap.Error(err3))

				return
			}

			history = append(history, arr0)

			return
		}

		goutils.Error("parseMissionDataParams:ArrayEach:func:dataType",
			zap.Int("offset", offset),
			zap.String("dataType", dataType.String()))
	}, "history")
	if err != nil {
		goutils.Error("parseMissionDataParams:ArrayEach",
			zap.Int("offset", offset),
			zap.Error(err))

		return nil, err
	}

	historyID, err := goutils.GetJsonInt(data, "srcHistory")
	if err != nil {
		goutils.Error("parseMissionDataParams:srcHistory",
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

package block7serv

import (
	block7game "github.com/zhs007/block7/game"
	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func MissionDataParams2Scene(scene *block7game.Scene, params *MissionDataParams) *block7game.Scene {
	scene.RngData = params.RngData[0:]
	scene.GameState = params.GameState
	scene.BlockNums = params.BlockNums
	scene.ClientMissionID = params.MissionID
	scene.ClientStageType = params.StageType
	scene.FirstItem = params.FirstItem

	for _, arr := range params.History {
		scene.History = append(scene.History, append([]int{}, arr...))
	}

	return scene
}

func MissionDataParams2SceneEx(params *MissionDataParams) *block7game.Scene {
	scene, err := block7game.NewSceneFromData(params.InitArr, params.SpecialLayers)
	if err != nil {
		goutils.Warn("MissionDataParams2SceneEx:NewSceneFromData",
			goutils.JSON("params", params),
			zap.Error(err))

		return nil
	}

	scene.RngData = params.RngData[0:]
	scene.GameState = params.GameState
	scene.BlockNums = params.BlockNums
	scene.ClientMissionID = params.MissionID
	scene.ClientStageType = params.StageType
	scene.FirstItem = params.FirstItem

	for _, arr := range params.History {
		scene.History = append(scene.History, append([]int{}, arr...))
	}

	return scene
}

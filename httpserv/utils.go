package block7serv

import block7game "github.com/zhs007/block7/game"

func MissionDataParams2Scene(scene *block7game.Scene, params *MissionDataParams) {
	for _, arr := range params.History {
		scene.History = append(scene.History, append([]int{}, arr...))
	}
}

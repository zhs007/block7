package block7serv

import (
	"github.com/zhs007/block7"
)

func MissionDataParams2Scene(scene *block7.Scene, params *MissionDataParams) {
	for _, arr := range params.History {
		scene.History = append(scene.History, append([]int{}, arr...))
	}
}

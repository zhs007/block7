package main

import (
	"fmt"

	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

func main() {
	block7utils.InitLogger("block7.ai", block7.Version,
		"debug", true, "./logs")

	stage, err := block7game.LoadStage("./cfg/level_0100.json")
	if err != nil {
		block7utils.Error("LoadStage",
			zap.Error(err))

		return
	}

	rng := block7game.NewRngNormal()

	for i := 0; i < 100; i++ {
		scene, err := block7game.NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, block7game.DefaultMaxBlockNums, nil)
		if err != nil {
			block7utils.Error("NewScene",
				zap.Error(err))

			return
		}

		// block7.AI1(scene, fmt.Sprintf("%v", i))
		block7game.AI4(rng, scene, fmt.Sprintf("%v", i), 300)
	}
	// mapBI := scene.Analysis()

	// mapBI.OutputLog("first")
}

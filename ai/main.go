package main

import (
	"fmt"

	"github.com/zhs007/block7"
	"go.uber.org/zap"
)

func main() {
	block7.InitLogger("block7.ai", block7.Version,
		"debug", true, "./logs")

	stage, err := block7.LoadStage("./cfg/level_0100.json")
	if err != nil {
		block7.Error("LoadStage",
			zap.Error(err))

		return
	}

	rng := block7.NewRngNormal()

	for i := 0; i < 100; i++ {
		scene, err := block7.NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, block7.DefaultMaxBlockNums)
		if err != nil {
			block7.Error("NewScene",
				zap.Error(err))

			return
		}

		// block7.AI1(scene, fmt.Sprintf("%v", i))
		block7.AI4(rng, scene, fmt.Sprintf("%v", i), 100)
	}
	// mapBI := scene.Analysis()

	// mapBI.OutputLog("first")
}

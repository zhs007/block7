package main

import (
	"fmt"

	"github.com/zhs007/block7"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func main() {
	goutils.InitLogger("block7.ai", block7.Version,
		"info", true, "./logs")

	rng := block7game.NewRngNormal()

	scene, err := block7game.LoadScene(rng, "./cfg/0-15.json", block7game.DefaultMaxBlockNums)
	if err != nil {
		goutils.Error("LoadStage",
			zap.Error(err))

		return
	}

	for i := 0; i < 1; i++ {
		// block7.AI1(scene, fmt.Sprintf("%v", i))
		// block7.AI2(rng, scene, fmt.Sprintf("%v", i), 1)
		// block7.AI3(scene, fmt.Sprintf("%v", i))
		block7game.AI4(rng, scene, fmt.Sprintf("%v", i), 1)
	}
	// mapBI := scene.Analysis()

	// mapBI.OutputLog("first")
}

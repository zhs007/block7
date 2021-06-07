package main

import (
	"fmt"

	"github.com/zhs007/block7"
	"go.uber.org/zap"
)

func main() {
	block7.InitLogger("block7.ai", block7.Version,
		"debug", true, "./logs")

	rng := block7.NewRngNormal()

	scene, err := block7.LoadScene(rng, "./cfg/0-15.json", block7.DefaultMaxBlockNums)
	if err != nil {
		block7.Error("LoadStage",
			zap.Error(err))

		return
	}

	for i := 0; i < 1; i++ {
		// block7.AI1(scene, fmt.Sprintf("%v", i))
		// block7.AI2(rng, scene, fmt.Sprintf("%v", i), 1)
		// block7.AI3(scene, fmt.Sprintf("%v", i))
		block7.AI4(rng, scene, fmt.Sprintf("%v", i), 1)
	}
	// mapBI := scene.Analysis()

	// mapBI.OutputLog("first")
}

package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// getBlockArea - ABCCCDDD -> A
func getBlockArea(block int) int {
	if block < 10000000 {
		return 0
	}

	return block / 10000000
}

// getBlockType - ABCCCDDD -> B
func getBlockType(block int) int {
	if block < 10000000 {
		return 1
	}

	return (block % 10000000) / 1000000
}

// getBlockSpecialLayer - ABCCCDDD -> C
func getBlockSpecialLayer(block int) int {
	if block < 10000000 {
		return 0
	}

	return ((block % 1000000) - (block % 1000)) / 1000
}

// getBlockSpecialBlock - ABCCCDDD -> D
func getBlockSpecialBlock(block int) int {
	if block < 10000000 {
		return 0
	}

	return (block % 1000)
}

// genSymbolList - 生成 nums 个数，每个数都在 initSymbols 里，且尽量平均，每个 symbol 都有3个倍数个
func genSymbolList(rng IRng, initSymbols []int, nums int) ([]int, error) {
	arr := []int{}

	if nums%BlockNums != 0 {
		goutils.Warn("genSymbolList",
			zap.Int("nums", nums),
			zap.Error(ErrInvalidMap2BlockNums))

		return nil, ErrInvalidMap2BlockNums
	}

	// 如果可以刚好平均分
	if (nums/BlockNums)%len(initSymbols) == 0 {
		n := nums / len(initSymbols)
		for _, v := range initSymbols {
			for i := 0; i < n; i++ {
				arr = append(arr, v)
			}
		}
	} else {
		n := nums / (len(initSymbols) * BlockNums)
		for _, v := range initSymbols {
			for i := 0; i < n; i++ {
				for j := 0; j < BlockNums; j++ {
					arr = append(arr, v)
				}
			}
		}

		ln := (nums - n*len(initSymbols)*BlockNums) / BlockNums
		lastsymbols := make([]int, len(initSymbols))
		copy(lastsymbols, initSymbols)

		for i := 0; i < ln; i++ {
			si, err := rng.Rand(len(lastsymbols))
			if err != nil {
				goutils.Warn("genSymbolList:Rand",
					zap.Int("nums", len(lastsymbols)),
					zap.Error(err))

				return nil, err
			}

			for j := 0; j < BlockNums; j++ {
				arr = append(arr, lastsymbols[si])
			}

			lastsymbols = append(lastsymbols[:si], lastsymbols[si+1:]...)
		}
	}

	return arr, nil
}

// NewScene2 - new a scene
func NewScene2(rng IRng, stage *Stage, symbols []int, blockNums int, ld2 *LevelData2) (*Scene, error) {
	if stage.MapType == 0 {
		return NewScene(rng, stage, symbols, blockNums, ld2)
	}

	if ld2.IconType2Ex == nil {
		goutils.Warn("NewScene2:IconType2Ex",
			goutils.JSON("ld2", ld2),
			zap.Error(ErrInvalidIconTypes2))

		return nil, ErrInvalidIconTypes2
	}

	// ss, err := MgrSpecial.GenSymbols(ld2)
	// if err != nil {
	// 	goutils.Warn("NewScene:MgrSpecial.GenSymbols",
	// 		zap.Error(err))

	// 	return nil, err
	// }

	// if len(ss) > stage.IconNums {
	// 	goutils.Warn("NewScene:IconNums",
	// 		zap.Error(ErrInvalidSpecialNums))

	// 	return nil, ErrInvalidSpecialNums
	// }

	// if len(ss) < stage.IconNums {
	// 	ss1, err := genSymbols(rng, symbols, stage.IconNums-len(ss))
	// 	if err != nil {
	// 		goutils.Warn("NewScene:genSymbols",
	// 			zap.Error(err))

	// 		return nil, err
	// 	}

	// 	ss = append(ss, ss1...)
	// }

	scene := &Scene{
		Width:        stage.Width,
		Height:       stage.Height,
		Layers:       len(stage.Layer),
		XOff:         stage.XOff,
		YOff:         stage.YOff,
		MaxBlockNums: blockNums,
		Offset:       stage.Offset,
		LayerLevel:   stage.LayerLevel,
	}

	if ld2 != nil {
		scene.SpecialType = ld2.SpecialTypeStr
	}

	nums := 0
	mapLayerPos := NewLayerPosMap()
	for z, arrlayer := range stage.Layer {
		arrslayer := [][]int{}

		for y, arrrow := range arrlayer {
			arrsrow := []int{}

			for x, v := range arrrow {
				arrsrow = append(arrsrow, 0)

				if v > 0 {
					nums++

					area := getBlockArea(v)
					d := getBlockSpecialBlock(v)

					if d == 0 {
						mapLayerPos.AddPos(x, y, z, area)
					}
				}
			}

			arrslayer = append(arrslayer, arrsrow)
		}

		scene.Arr = append(scene.Arr, arrslayer)
	}

	for area, lpl := range mapLayerPos.MapLayerPos {
		if len(lpl.Pos)%3 != 0 {
			goutils.Warn("NewScene2:MapLayerPos",
				zap.Int("area", area),
				zap.Int("length", len(lpl.Pos)),
				zap.Error(ErrInvalidMap2BlockNums))

			return nil, ErrInvalidMap2BlockNums
		}

		if area <= 0 || area > len(ld2.IconType2Ex) {
			goutils.Warn("NewScene2:MapLayerPos",
				goutils.JSON("ld2", ld2),
				zap.Int("area", area),
				zap.Error(ErrInvalidIconTypes2))

			return nil, ErrInvalidIconTypes2
		}

		symbolarr := ld2.IconType2Ex[area-1]

		arr2, err := genSymbolList(rng, symbolarr, len(lpl.Pos))
		if err != nil {
			goutils.Warn("NewScene2:genSymbolList",
				zap.Error(err))

			return nil, err
		}

		for _, pos := range lpl.Pos {
			si, err := rng.Rand(len(arr2))
			if err != nil {
				goutils.Warn("NewScene2:Rand",
					zap.Int("nums", len(arr2)),
					zap.Error(err))

				return nil, err
			}

			scene.Arr[pos.Z][pos.Y][pos.X] = arr2[si]

			arr2 = append(arr2[:si], arr2[si+1:]...)
		}

		if len(arr2) != 0 {
			goutils.Warn("NewScene2:check arr2",
				zap.Int("arr2 len", len(arr2)))

			return nil, err
		}
	}

	// nums := 0
	// for _, arrlayer := range stage.Layer {
	// 	arrslayer := [][]int{}
	// 	for _, arrrow := range arrlayer {
	// 		arrsrow := []int{}
	// 		for _, v := range arrrow {
	// 			if v == 0 {
	// 				arrsrow = append(arrsrow, 0)
	// 			} else {
	// 				nss, cs, err := randSymbols(rng, ss)
	// 				if err != nil {
	// 					return nil, err
	// 				}

	// 				arrsrow = append(arrsrow, cs)
	// 				ss = nss

	// 				nums++
	// 			}
	// 		}

	// 		arrslayer = append(arrslayer, arrsrow)
	// 	}

	// 	scene.Arr = append(scene.Arr, arrslayer)
	// }

	scene.InitArr = goutils.CloneArr3(scene.Arr)
	scene.BlockNums = nums

	for z, arrlayer := range stage.Layer {
		for y, arrrow := range arrlayer {
			for x, v := range arrrow {
				if v > 0 {
					nums++

					c := getBlockSpecialLayer(v)
					d := getBlockSpecialBlock(v)

					err := MgrSpecial.Gen2(scene, x, y, z, c, d)
					if err != nil {
						goutils.Warn("NewScene:MgrSpecial.Gen2",
							zap.Error(err))

						return nil, err
					}
				}
			}
		}
	}

	err := scene.resetInitLayerArr(stage)
	if err != nil {
		goutils.Warn("NewScene2:resetInitLayerArr",
			zap.Error(err))

		return nil, err
	}
	// err = MgrSpecial.OnFixScene(rng, ld2, scene)
	// if err != nil {
	// 	goutils.Warn("NewScene:OnFixScene",
	// 		zap.Error(err))

	// 	return nil, err
	// }

	// err = MgrSpecial.GenSymbolLayers(rng, ld2, scene)
	// if err != nil {
	// 	goutils.Warn("NewScene:GenSymbolLayers",
	// 		zap.Error(err))

	// 	return nil, err
	// }

	return scene, nil
}

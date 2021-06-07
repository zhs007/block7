package block7

import (
	"fmt"
	"os"
	"path"

	"go.uber.org/zap"
)

const AI4OutputPath = "./ai4_output"

func ai4L0(rng Rng, scene *Scene, mapBI *BlockInfoMap, aiResult *AIResult, symbol int) bool {
	if aiResult.HasSymbol(symbol) {
		return false
	}

	aiResult.StartSymbol(symbol)

	v, isok := mapBI.MapBlockInfo[symbol]
	if isok {
		lsn := scene.CountBlockSymbols(symbol)
		if lsn > 0 {
			if len(v.LevelList[0]) >= BlockNums-lsn {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[0], BlockNums-lsn)
				if err != nil {
					Error("ai4L0:lsn:RandBlockData",
						zap.Int("lsn", lsn),
						zap.Int("len", len(v.LevelList[0])),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}

				for _, b := range lst1 {
					if !aiResult.ClickEx(symbol, scene, b) {
						Error("ai4L0:lsn:ClickEx",
							zap.Int("lsn", lsn),
							zap.Int("len", len(v.LevelList[0])),
							zap.Error(err))

						aiResult.StopSymbol(symbol, -1)

						return false
					}
				}

				v.LevelList[0] = lst0

				return true
			}
		}

		if len(v.LevelList[0]) >= BlockNums {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[0], len(v.LevelList[0])-len(v.LevelList[0])%BlockNums)
			if err != nil {
				Error("ai4L0:RandBlockData",
					zap.Int("len", len(v.LevelList[0])),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L0:ClickEx",
						zap.Int("len", len(v.LevelList[0])),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			v.LevelList[0] = lst0

			return true
		}
	}

	return false
}

func ai4L1(rng Rng, scene *Scene, mapbi *BlockInfoMap, aiResult *AIResult, symbol int) bool {
	var lstl0 []*BlockData
	var lstl1 []*BlockData

	if aiResult.HasSymbol(symbol) {
		return false
	}

	aiResult.StartSymbol(symbol)

	v, isok := mapbi.MapBlockInfo[symbol]
	if !isok {
		return false
	}

	lsn := scene.CountBlockSymbols(symbol)
	if lsn > 0 {
		if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums-lsn {
			ln0 := len(v.LevelList[0])
			if ln0 >= BlockNums-lsn {
				Error("ai4L1:lsn:ln0",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			if ln0 > 0 {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
				if err != nil {
					Error("ai4L1:lsn:ln0:RandBlockData",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}

				for _, b := range lst1 {
					if !aiResult.ClickEx(symbol, scene, b) {
						Error("ai4L1:lsn:ln0:ClickEx",
							zap.Int("lsn", lsn),
							zap.Int("ln0", ln0),
							zap.Error(err))

						aiResult.StopSymbol(symbol, -1)

						return false
					}
				}

				lstl0 = lst0
			}

			ln1 := BlockNums - lsn - ln0
			if ln1 <= 0 || ln1 > len(v.LevelList[1]) {
				Error("ai4L1:lsn:ln1",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Int("len1", len(v.LevelList[1])))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			lst0, lst1, err := RandBlockData(rng, v.LevelList[1], ln1)
			if err != nil {
				Error("ai4L1:lsn:ln1:RandBlockData",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L1:lsn:ln1:ClickEx",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Int("ln1", ln1),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			lstl1 = lst0

			v.LevelList[0] = lstl0
			v.LevelList[1] = lstl1

			return true
		}

		return false
	}

	if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
		ln0 := len(v.LevelList[0])
		if ln0 >= BlockNums {
			Error("ai4L1:ln0",
				zap.Int("ln0", ln0))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		if ln0 > 0 {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
			if err != nil {
				Error("ai4L1:ln0:RandBlockData",
					zap.Int("ln0", ln0),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L1:ln0:ClickEx",
						zap.Int("ln0", ln0),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			lstl0 = lst0
		}

		ln1 := BlockNums - ln0
		if ln1 <= 0 || ln1 > len(v.LevelList[1]) {
			Error("ai4L1:ln1",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Int("len1", len(v.LevelList[1])))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		lst0, lst1, err := RandBlockData(rng, v.LevelList[1], ln1)
		if err != nil {
			Error("ai4L1:ln1:RandBlockData",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Error(err))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		for _, b := range lst1 {
			if !aiResult.ClickEx(symbol, scene, b) {
				Error("ai4L1:ln1:ClickEx",
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}
		}

		lstl1 = lst0

		v.LevelList[0] = lstl0
		v.LevelList[1] = lstl1

		return true
	}

	return false
}

func ai4L2(rng Rng, scene *Scene, mapbi *BlockInfoMap, aiResult *AIResult, symbol int) bool {
	var lstl0 []*BlockData
	var lstl1 []*BlockData
	var lstl2 []*BlockData

	if aiResult.HasSymbol(symbol) {
		return false
	}

	aiResult.StartSymbol(symbol)

	v, isok := mapbi.MapBlockInfo[symbol]
	if !isok {
		return false
	}

	lsn := scene.CountBlockSymbols(symbol)
	if lsn > 0 {
		if len(v.LevelList[0])+len(v.LevelList[1])+len(v.LevelList[2]) >= BlockNums-lsn {
			ln0 := len(v.LevelList[0])
			if ln0 >= BlockNums-lsn {
				Error("ai4L2:lsn:ln0",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			if ln0 > 0 {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
				if err != nil {
					Error("ai4L2:lsn:ln0:RandBlockData",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}

				for _, b := range lst1 {
					if !aiResult.ClickEx(symbol, scene, b) {
						Error("ai4L2:lsn:ln0:ClickEx",
							zap.Int("lsn", lsn),
							zap.Int("ln0", ln0),
							zap.Error(err))

						aiResult.StopSymbol(symbol, -1)

						return false
					}
				}

				lstl0 = lst0
			}

			ln1 := len(v.LevelList[1])
			if ln1+ln0 >= BlockNums-lsn {
				Error("ai4L2:lsn:ln1",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			if ln1 > 0 {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[1], ln1)
				if err != nil {
					Error("ai4L2:lsn:ln1:RandBlockData",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Int("ln1", ln1),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}

				for _, b := range lst1 {
					if !aiResult.ClickEx(symbol, scene, b) {
						Error("ai4L2:lsn:ln1:ClickEx",
							zap.Int("lsn", lsn),
							zap.Int("ln0", ln0),
							zap.Int("ln1", ln1),
							zap.Error(err))

						aiResult.StopSymbol(symbol, -1)

						return false
					}
				}

				lstl1 = lst0
			}

			ln2 := BlockNums - lsn - ln0 - ln1
			if ln2 <= 0 || ln2 > len(v.LevelList[2]) {
				Error("ai4L2:lsn:ln2",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Int("ln2", ln2),
					zap.Int("len2", len(v.LevelList[2])))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			lst0, lst1, err := RandBlockData(rng, v.LevelList[2], ln2)
			if err != nil {
				Error("ai4L2:lsn:ln2:RandBlockData",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Int("ln2", ln2),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L2:lsn:ln2:ClickEx",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Int("ln1", ln1),
						zap.Int("ln2", ln2),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			lstl2 = lst0

			v.LevelList[0] = lstl0
			v.LevelList[1] = lstl1
			v.LevelList[2] = lstl2

			return true
		}

		return false
	}

	if len(v.LevelList[0])+len(v.LevelList[1])+len(v.LevelList[2]) >= BlockNums {
		ln0 := len(v.LevelList[0])
		if ln0 >= BlockNums {
			Error("ai4L2:ln0",
				zap.Int("ln0", ln0))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		if ln0 > 0 {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
			if err != nil {
				Error("ai4L2:ln0:RandBlockData",
					zap.Int("ln0", ln0),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L2:ln0:ClickEx",
						zap.Int("ln0", ln0),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			lstl0 = lst0
		}

		ln1 := len(v.LevelList[1])
		if ln0+ln1 >= BlockNums {
			Error("ai4L2:ln1",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		if ln1 > 0 {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[1], ln1)
			if err != nil {
				Error("ai4L2:ln1:RandBlockData",
					zap.Int("ln0", ln0),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}

			for _, b := range lst1 {
				if !aiResult.ClickEx(symbol, scene, b) {
					Error("ai4L2:ln0:ClickEx",
						zap.Int("ln0", ln0),
						zap.Int("ln1", ln1),
						zap.Error(err))

					aiResult.StopSymbol(symbol, -1)

					return false
				}
			}

			lstl1 = lst0
		}

		ln2 := BlockNums - ln0 - ln1
		if ln2 <= 0 || ln2 > len(v.LevelList[2]) {
			Error("ai4L2:ln2",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Int("ln2", ln2),
				zap.Int("len2", len(v.LevelList[2])))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		lst0, lst1, err := RandBlockData(rng, v.LevelList[2], ln2)
		if err != nil {
			Error("ai4L2:ln2:RandBlockData",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Int("ln2", ln2),
				zap.Error(err))

			aiResult.StopSymbol(symbol, -1)

			return false
		}

		for _, b := range lst1 {
			if !aiResult.ClickEx(symbol, scene, b) {
				Error("ai4L2:ln2:ClickEx",
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Int("ln2", ln2),
					zap.Error(err))

				aiResult.StopSymbol(symbol, -1)

				return false
			}
		}

		lstl2 = lst0

		v.LevelList[0] = lstl0
		v.LevelList[1] = lstl1
		v.LevelList[2] = lstl2

		return true
	}

	return false
}

func ai4PreProc(rng Rng, scene *Scene, mapBI *BlockInfoMap) (*AIResult, error) {
	aiResult := NewAIResult(scene, mapBI)

	if len(mapBI.BlockSymbols) > 0 {
		for _, v := range mapBI.BlockSymbols {
			ai4L0(rng, scene, mapBI, aiResult, v)
		}

		for _, v := range mapBI.BlockSymbols {
			ai4L1(rng, scene, mapBI, aiResult, v)
		}

		// for _, v := range mapBI.BlockSymbols {
		// 	ai4L2(scene, mapBI, aiResult, v)
		// }
	}

	for k := range mapBI.MapBlockInfo {
		ai4L0(rng, scene, mapBI, aiResult, k)
	}

	for k := range mapBI.MapBlockInfo {
		ai4L1(rng, scene, mapBI, aiResult, k)
	}

	// for k := range mapBI.MapBlockInfo {
	// 	ai4L2(scene, mapBI, aiResult, k)
	// }

	return aiResult, nil
}

func procAI4(rng Rng, scene *Scene, name string) (bool, error) {
	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		aiResult, err := ai4PreProc(rng, scene, mapbi)
		if err != nil {
			Warn("AI4:ai4PreProc",
				zap.Error(err))

			return false, err
		}

		cs := 0
		an := 0
		ln := scene.MaxBlockNums
		for k, v := range aiResult.MapSymbols {
			if v.State < 0 {
				continue
			}

			can := len(v.Arr) - len(v.LastBlocks) + len(scene.Block)
			cln := len(v.LastBlocks)

			if can > 0 {
				if cln < ln {
					if can >= an {
						ln = cln
						an = can
						cs = k
					}
				} else if cln <= ln {
					if can > an {
						ln = cln
						an = can
						cs = k
					}
				}
			}
		}

		if cs > 0 {
			aibr, isok := aiResult.MapSymbols[cs]
			if isok {
				Debug("AI4:Symbol",
					zap.Int("symbol", aibr.Symbol),
					zap.Int("arr", len(aibr.Arr)),
					zap.Int("block", len(aibr.LastBlocks)))

				for _, b := range aibr.Arr {
					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						Warn("AI4:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						Info("AI4:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))
					}
				}
			}
		}

		if scene.CountSymbols() == 0 {
			// fn := fmt.Sprintf("%v.%v.json", "ok", name)
			// scene.Save(path.Join(AI4OutputPath, fn))

			return true, nil
		}

		if clicknums > 0 {
			Info("AI4:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))
		} else {
			Info("AI4:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))

			// fn := fmt.Sprintf("%v.%v.json", "fail", name)
			// scene.Save(path.Join(AI4OutputPath, fn))

			return false, nil
		}
	}

	return false, nil
}

func AI4(rng Rng, scene *Scene, name string, totalnums int) error {
	os.MkdirAll(AI4OutputPath, os.ModePerm)

	if totalnums > 1 {
		finishedNums := 0
		for i := 0; i < totalnums; i++ {
			isok, err := procAI4(rng, scene, name)
			if err != nil {
				Error("AI4:procAI4",
					zap.Error(err))
			}

			if isok {
				finishedNums++
			}
		}

		scene.FinishedPer = float32(finishedNums) / float32(totalnums)

		fn := fmt.Sprintf("%v-%v.json", scene.FinishedPer, name)
		scene.Save(path.Join(AI4OutputPath, fn))

		return nil
	}

	isok, err := procAI4(rng, scene, name)
	if err != nil {
		Error("AI4:procAI4",
			zap.Error(err))
	}

	if isok {
		fn := fmt.Sprintf("%v.%v.json", "ok", name)
		scene.Save(path.Join(AI4OutputPath, fn))
	} else {
		fn := fmt.Sprintf("%v.%v.json", "fail", name)
		scene.Save(path.Join(AI4OutputPath, fn))
	}

	return nil
}

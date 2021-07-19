package block7

import (
	"fmt"
	"os"
	"path"

	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

const AI3OutputPath = "./ai3_output"

func ai3L0(scene *Scene, mapBI *BlockInfoMap, aiResult *AIResult, symbol int) bool {
	if aiResult.HasSymbol(symbol) {
		return false
	}

	aiResult.StartSymbol(symbol)

	v, isok := mapBI.MapBlockInfo[symbol]
	if isok {
		lsn := scene.CountBlockSymbols(symbol)
		if lsn > 0 {
			if len(v.LevelList[0]) >= BlockNums-lsn {
				lst0, lst1, err := GetBlockDataList(v.LevelList[0], BlockNums-lsn)
				if err != nil {
					block7utils.Error("ai3L0:lsn:GetBlockDataList",
						zap.Int("lsn", lsn),
						zap.Int("len", len(v.LevelList[0])),
						zap.Error(err))

					return true
				}

				for _, b := range lst1 {
					if !aiResult.Click(symbol, scene, b) {
						block7utils.Error("ai3L0:lsn:Click",
							zap.Int("lsn", lsn),
							zap.Int("len", len(v.LevelList[0])),
							zap.Error(err))

						return true
					}
				}

				v.LevelList[0] = lst0

				return true
			}
		}

		if len(v.LevelList[0]) >= BlockNums {
			lst0, lst1, err := GetBlockDataList(v.LevelList[0], len(v.LevelList[0])-len(v.LevelList[0])%BlockNums)
			if err != nil {
				block7utils.Error("ai3L0:GetBlockDataList",
					zap.Int("len", len(v.LevelList[0])),
					zap.Error(err))

				return true
			}

			for _, b := range lst1 {
				if !aiResult.Click(symbol, scene, b) {
					block7utils.Error("ai3L0:Click",
						zap.Int("len", len(v.LevelList[0])),
						zap.Error(err))

					return true
				}
			}

			v.LevelList[0] = lst0

			return true
		}
	}

	return false
}

func ai3L1(scene *Scene, mapbi *BlockInfoMap, aiResult *AIResult, symbol int) bool {
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
				block7utils.Error("ai3L1:lsn:ln0",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0))

				return true
			}

			if ln0 > 0 {
				lst0, lst1, err := GetBlockDataList(v.LevelList[0], ln0)
				if err != nil {
					block7utils.Error("ai3L1:lsn:ln0:GetBlockDataList",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Error(err))

					return true
				}

				for _, b := range lst1 {
					if !aiResult.Click(symbol, scene, b) {
						block7utils.Error("ai3L1:lsn:ln0:Click",
							zap.Int("lsn", lsn),
							zap.Int("ln0", ln0),
							zap.Error(err))

						return true
					}
				}

				v.LevelList[0] = lst0
			}

			ln1 := BlockNums - lsn - ln0
			if ln1 <= 0 || ln1 > len(v.LevelList[1]) {
				block7utils.Error("ai3L1:lsn:ln1",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Int("len1", len(v.LevelList[1])))

				return true
			}

			lst0, lst1, err := GetBlockDataList(v.LevelList[1], ln1)
			if err != nil {
				block7utils.Error("ai3L1:lsn:ln1:GetBlockDataList",
					zap.Int("lsn", lsn),
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Error(err))

				return true
			}

			for _, b := range lst1 {
				if len(b.Parent) > 0 {
					for _, bc := range b.Parent {
						if !aiResult.Click(symbol, scene, bc) {
							block7utils.Error("ai3L1:lsn:ln1:parent:Click",
								zap.Int("lsn", lsn),
								zap.Int("ln0", ln0),
								zap.Int("ln1", ln1),
								zap.Error(err))

							return true
						}
					}
				}

				if !aiResult.Click(symbol, scene, b) {
					block7utils.Error("ai3L1:lsn:ln1:Click",
						zap.Int("lsn", lsn),
						zap.Int("ln0", ln0),
						zap.Int("ln1", ln1),
						zap.Error(err))

					return true
				}
			}

			v.LevelList[1] = lst0

			return true
		}

		return false
	}

	if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
		ln0 := len(v.LevelList[0])
		if ln0 >= BlockNums {
			block7utils.Error("ai3L1:ln0",
				zap.Int("ln0", ln0))

			return true
		}

		if ln0 > 0 {
			lst0, lst1, err := GetBlockDataList(v.LevelList[0], ln0)
			if err != nil {
				block7utils.Error("ai3L1:ln0:GetBlockDataList",
					zap.Int("ln0", ln0),
					zap.Error(err))

				return true
			}

			for _, b := range lst1 {
				if !aiResult.Click(symbol, scene, b) {
					block7utils.Error("ai3L1:ln0:Click",
						zap.Int("ln0", ln0),
						zap.Error(err))

					return true
				}
			}

			v.LevelList[0] = lst0
		}

		ln1 := BlockNums - ln0
		if ln1 <= 0 || ln1 > len(v.LevelList[1]) {
			block7utils.Error("ai3L1:ln1",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Int("len1", len(v.LevelList[1])))

			return true
		}

		lst0, lst1, err := GetBlockDataList(v.LevelList[1], ln1)
		if err != nil {
			block7utils.Error("ai3L1:ln1:GetBlockDataList",
				zap.Int("ln0", ln0),
				zap.Int("ln1", ln1),
				zap.Error(err))

			return true
		}

		for _, b := range lst1 {
			if len(b.Parent) > 0 {
				for _, bc := range b.Parent {
					if !aiResult.Click(symbol, scene, bc) {
						block7utils.Error("ai3L1:ln1:parent:Click",
							zap.Int("ln0", ln0),
							zap.Int("ln1", ln1),
							zap.Error(err))

						return true
					}
				}
			}

			if !aiResult.Click(symbol, scene, b) {
				block7utils.Error("ai3L1:ln1:Click",
					zap.Int("ln0", ln0),
					zap.Int("ln1", ln1),
					zap.Error(err))

				return true
			}
		}

		v.LevelList[1] = lst0

		return true
	}

	return false
}

func ai3PreProc(scene *Scene, mapBI *BlockInfoMap) (*AIResult, error) {
	aiResult := NewAIResult(scene, mapBI)

	if len(mapBI.BlockSymbols) > 0 {
		for _, v := range mapBI.BlockSymbols {
			ai3L0(scene, mapBI, aiResult, v)
		}

		for _, v := range mapBI.BlockSymbols {
			ai3L1(scene, mapBI, aiResult, v)
		}

	}

	for k := range mapBI.MapBlockInfo {
		ai3L0(scene, mapBI, aiResult, k)
	}

	for k := range mapBI.MapBlockInfo {
		ai3L1(scene, mapBI, aiResult, k)
	}

	return aiResult, nil
}

func AI3(scene *Scene, name string) error {
	os.MkdirAll(AI3OutputPath, os.ModePerm)

	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		aiResult, err := ai3PreProc(scene, mapbi)
		if err != nil {
			block7utils.Warn("AI3:ai3PreProc",
				zap.Error(err))

			return err
		}

		cs := 0
		an := 0
		ln := scene.MaxBlockNums
		for k, v := range aiResult.MapSymbols {
			can := len(v.Arr)
			cln := len(v.LastBlocks)

			if cln < ln {
				if can > an {
					ln = cln
					an = can
					cs = k
				}
			}
		}

		if cs > 0 {
			aibr, isok := aiResult.MapSymbols[cs]
			if isok {
				for _, b := range aibr.Arr {
					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						block7utils.Warn("AI3:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						block7utils.Info("AI3:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))
					}
				}
			}
		}

		if scene.CountSymbols() == 0 {
			fn := fmt.Sprintf("%v.%v.json", "ok", name)
			scene.Save(path.Join(AI3OutputPath, fn))

			break
		}

		if clicknums > 0 {
			block7utils.Info("AI3:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))
		} else {
			block7utils.Info("AI3:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))

			fn := fmt.Sprintf("%v.%v.json", "fail", name)
			scene.Save(path.Join(AI3OutputPath, fn))

			break
		}
	}

	return nil
}

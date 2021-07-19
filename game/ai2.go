package block7game

import (
	"fmt"
	"os"
	"path"

	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

const AI2OutputPath = "./ai2_output"

func ai2L0(rng IRng, scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
	clicknums := 0
	v, isok := mapbi.MapBlockInfo[symbol]
	if isok {
		lsn := scene.CountBlockSymbols(symbol)
		if lsn > 0 {
			if len(v.LevelList[0]) >= BlockNums-lsn {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[0], BlockNums-lsn)
				if err != nil {
					block7utils.Error("ai2L0:RandBlockData:BlockNums-lsn",
						zap.Error(err))

					return clicknums, true
				}

				for _, b := range lst1 {
					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						block7utils.Warn("ai2L0:Click:BlockNums-lsn",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						block7utils.Info("ai2L0:Click:BlockNums-lsn",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return clicknums, true
					}
				}

				v.LevelList[0] = lst0
			}

			return clicknums, true
		}

		if len(v.LevelList[0]) >= BlockNums {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[0], len(v.LevelList[0])-len(v.LevelList[0])%BlockNums)
			if err != nil {
				block7utils.Error("ai2L0:RandBlockData",
					zap.Error(err))

				return clicknums, true
			}

			for _, b := range lst1 {
				gs, isok := scene.Click(b.X, b.Y, b.Z)
				if !isok {
					block7utils.Warn("ai2L0:Click",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					clicknums++
				}

				if gs != GameStateRunning {
					block7utils.Info("ai2L0:Click",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}
			}

			v.LevelList[0] = lst0
		}
	}

	return clicknums, false
}

func ai2L1(rng IRng, scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
	clicknums := 0
	v, isok := mapbi.MapBlockInfo[symbol]
	if !isok {
		return 0, false
	}

	lsn := scene.CountBlockSymbols(symbol)
	if lsn > 0 {
		if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums-lsn {
			ln0 := len(v.LevelList[0])
			if ln0 > 0 {
				lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
				if err != nil {
					block7utils.Error("ai2L1:RandBlockData:BlockNums-lsn:l0",
						zap.Error(err))

					return clicknums, true
				}

				// cn := 0
				for _, b := range lst1 {
					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						block7utils.Warn("ai2L1:Click:L1L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						// cn++
						clicknums++
					}

					if gs != GameStateRunning {
						block7utils.Info("ai2L1:Click:L1L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return clicknums, true
					}
				}

				v.LevelList[0] = lst0
			}

			if len(v.LevelList[1]) > 0 && BlockNums-lsn-ln0 > 0 {
				lst2, lst3, err := RandBlockData(rng, v.LevelList[1], BlockNums-lsn-ln0)
				if err != nil {
					block7utils.Error("ai2L1:RandBlockData:BlockNums-lsn:l1",
						zap.Error(err))

					return clicknums, true
				}

				for _, b := range lst3 {
					if len(b.Parent) > 0 {
						for _, bc := range b.Parent {
							gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
							if !isok {
								block7utils.Warn("ai2L1:Click:L1:Parent",
									zap.Int("x", bc.X),
									zap.Int("y", bc.Y),
									zap.Int("z", bc.Z))
							} else {
								clicknums++
							}

							if gs != GameStateRunning {
								block7utils.Info("ai2L1:Click:L1:Parent",
									zap.Int("x", bc.X),
									zap.Int("y", bc.Y),
									zap.Int("z", bc.Z),
									zap.Int("gameState", gs))

								return clicknums, true
							}
						}
					}

					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						block7utils.Warn("ai2L1:Click:L1",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						// cn++
						clicknums++
					}

					if gs != GameStateRunning {
						block7utils.Info("ai2L1:Click:L1",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return clicknums, true
					}
				}

				v.LevelList[1] = lst2
			}

			return clicknums, true
		}

		return 0, false
	}

	if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
		ln := len(v.LevelList[0]) + len(v.LevelList[1]) - (len(v.LevelList[0])+len(v.LevelList[1]))%BlockNums

		ln0 := len(v.LevelList[0])
		if ln0 > 0 {
			lst0, lst1, err := RandBlockData(rng, v.LevelList[0], ln0)
			if err != nil {
				block7utils.Error("ai2L1:RandBlockData:l0",
					zap.Error(err))

				return clicknums, true
			}

			cn := 0
			for _, b := range lst1 {
				gs, isok := scene.Click(b.X, b.Y, b.Z)
				if !isok {
					block7utils.Warn("ai2L1:Click:L1L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					cn++
					clicknums++
				}

				if gs != GameStateRunning {
					block7utils.Info("ai2L1:Click:L1L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}
			}

			v.LevelList[0] = lst0
		}

		if len(v.LevelList[1])-ln > 0 {
			lst2, lst3, err := RandBlockData(rng, v.LevelList[1], len(v.LevelList[1])-ln)
			if err != nil {
				block7utils.Error("ai2L1:RandBlockData:l0",
					zap.Error(err))

				return clicknums, true
			}

			for _, b := range lst3 {
				if len(b.Parent) > 0 {
					for _, bc := range b.Parent {
						gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
						if !isok {
							block7utils.Warn("ai2L1:Click:L1:Parent",
								zap.Int("x", bc.X),
								zap.Int("y", bc.Y),
								zap.Int("z", bc.Z))
						} else {
							clicknums++
						}

						if gs != GameStateRunning {
							block7utils.Info("ai2L1:Click:L1:Parent",
								zap.Int("x", bc.X),
								zap.Int("y", bc.Y),
								zap.Int("z", bc.Z),
								zap.Int("gameState", gs))

							return clicknums, true
						}
					}
				}

				gs, isok := scene.Click(b.X, b.Y, b.Z)
				if !isok {
					block7utils.Warn("ai2L1:Click:L1",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					// cn++
					clicknums++
				}

				if gs != GameStateRunning {
					block7utils.Info("ai2L1:Click:L1",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}
			}

			v.LevelList[1] = lst2
		}

		return clicknums, true
	}

	return 0, false
}

func ai2(rng IRng, scene *Scene) bool {
	scene.Restart()

	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		if len(scene.Block) > 0 {
			for _, b := range scene.Block {
				cn, isbreak := ai2L0(rng, scene, mapbi, b.Symbol)
				clicknums += cn
				if isbreak {
					break
				}
			}

			for _, b := range scene.Block {
				cn, isbreak := ai2L1(rng, scene, mapbi, b.Symbol)
				clicknums += cn
				if isbreak {
					break
				}
			}
		}

		if len(mapbi.BlockSymbols) > 0 {
			for _, v := range mapbi.BlockSymbols {
				cn, isbreak := ai2L0(rng, scene, mapbi, v)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for _, v := range mapbi.BlockSymbols {
					cn, isbreak := ai2L1(rng, scene, mapbi, v)
					clicknums += cn
					if isbreak {
						break
					}
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn, isbreak := ai2L0(rng, scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}

				if clicknums == 0 {
					for k := range mapbi.MapBlockInfo {
						cn, isbreak := ai2L1(rng, scene, mapbi, k)
						clicknums += cn
						if isbreak {
							break
						}
					}
				}
			}
		} else {
			for k := range mapbi.MapBlockInfo {
				cn, isbreak := ai2L0(rng, scene, mapbi, k)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn, isbreak := ai2L1(rng, scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}
			}
		}

		if scene.CountSymbols() == 0 {
			return true
		}

		if clicknums > 0 {
			block7utils.Info("ai2:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))
		} else {
			block7utils.Info("ai2:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))

			// fn := fmt.Sprintf("%v.%v.json", "fail", name)
			// scene.Save(path.Join("./ai2_output", fn))

			return false
		}
	}
}

func AI2(rng IRng, scene *Scene, name string, totalnums int) {
	os.MkdirAll(AI2OutputPath, os.ModePerm)

	finishedNums := 0
	for i := 0; i < totalnums; i++ {
		if ai2(rng, scene) {
			finishedNums++
		}
	}

	scene.FinishedPer = float32(finishedNums) / float32(totalnums)

	fn := fmt.Sprintf("%v-%v.json", scene.FinishedPer, name)
	scene.Save(path.Join(AI2OutputPath, fn))
}

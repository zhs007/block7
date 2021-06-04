package block7

import (
	"fmt"
	"path"

	"go.uber.org/zap"
)

func ai2L0(scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
	clicknums := 0
	v, isok := mapbi.MapBlockInfo[symbol]
	if isok {
		lsn := scene.CountBlockSymbols(symbol)
		if lsn > 0 {
			if len(v.LevelList[0]) >= BlockNums-lsn {
				for _, b := range v.LevelList[0] {
					if clicknums >= BlockNums-lsn {
						break
					}

					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						Warn("ai2:Click:L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						Info("ai2:Click:L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return clicknums, true
					}
				}
			}

			return clicknums, true
		}

		if len(v.LevelList[0]) >= BlockNums {
			cn := len(v.LevelList[0]) / BlockNums
			for _, b := range v.LevelList[0] {
				if clicknums >= cn*BlockNums {
					break
				}

				gs, isok := scene.Click(b.X, b.Y, b.Z)
				if !isok {
					Warn("ai2:Click:L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					clicknums++
				}

				if gs != GameStateRunning {
					Info("ai2:Click:L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}
			}
		}
	}

	return clicknums, false
}

func ai2L1(scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
	clicknums := 0
	v, isok := mapbi.MapBlockInfo[symbol]
	if !isok {
		return 0, false
	}

	lsn := scene.CountBlockSymbols(symbol)
	if lsn > 0 {
		if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums-lsn {
			cn := 0
			for _, b := range v.LevelList[0] {
				gs, isok := scene.Click(b.X, b.Y, b.Z)
				if !isok {
					Warn("ai2:Click:L1L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					cn++
					clicknums++
				}

				if gs != GameStateRunning {
					Info("ai2:Click:L1L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}
			}

			for _, b := range v.LevelList[1] {

				if len(b.Parent) > 0 {
					for _, bc := range b.Parent {
						gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
						if !isok {
							Warn("ai2:Click:L1:Parent",
								zap.Int("x", bc.X),
								zap.Int("y", bc.Y),
								zap.Int("z", bc.Z))
						} else {
							clicknums++
						}

						if gs != GameStateRunning {
							Info("ai2:Click:L1:Parent",
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
					Warn("ai2:Click:L1",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					cn++
					clicknums++
				}

				if gs != GameStateRunning {
					Info("ai2:Click:L1",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z),
						zap.Int("gameState", gs))

					return clicknums, true
				}

				if cn == BlockNums-lsn {
					break
				}
			}

			return clicknums, true
		}

		return 0, false
	}

	if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
		cn := 0
		for _, b := range v.LevelList[0] {
			gs, isok := scene.Click(b.X, b.Y, b.Z)
			if !isok {
				Warn("ai2:Click:L1L0",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			} else {
				cn++
				clicknums++
			}

			if gs != GameStateRunning {
				Info("ai2:Click:L1L0",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z),
					zap.Int("gameState", gs))

				return clicknums, true
			}
		}

		for _, b := range v.LevelList[1] {

			if len(b.Parent) > 0 {
				for _, bc := range b.Parent {
					gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
					if !isok {
						Warn("ai2:Click:L1:Parent",
							zap.Int("x", bc.X),
							zap.Int("y", bc.Y),
							zap.Int("z", bc.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						Info("ai2:Click:L1:Parent",
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
				Warn("ai2:Click:L1",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			} else {
				cn++
				clicknums++
			}

			if gs != GameStateRunning {
				Info("ai2:Click:L1",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z),
					zap.Int("gameState", gs))

				return clicknums, true
			}

			if cn == BlockNums {
				break
			}
		}

		return clicknums, true
	}

	return 0, false
}

func AI2(scene *Scene, name string) {
	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		if len(mapbi.BlockSymbols) > 0 {
			for _, v := range mapbi.BlockSymbols {
				cn, isbreak := ai2L0(scene, mapbi, v)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for _, v := range mapbi.BlockSymbols {
					cn, isbreak := ai2L1(scene, mapbi, v)
					clicknums += cn
					if isbreak {
						break
					}
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn, isbreak := ai2L0(scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}

				if clicknums == 0 {
					for k := range mapbi.MapBlockInfo {
						cn, isbreak := ai2L1(scene, mapbi, k)
						clicknums += cn
						if isbreak {
							break
						}
					}
				}
			}
		} else {
			for k := range mapbi.MapBlockInfo {
				cn, isbreak := ai2L0(scene, mapbi, k)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn, isbreak := ai2L1(scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}
			}
		}

		if scene.CountSymbols() == 0 {
			fn := fmt.Sprintf("%v.%v.json", "ok", name)
			scene.Save(path.Join("./ai2_output", fn))

			break
		}

		if clicknums > 0 {
			Info("ai2:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))
		} else {
			Info("ai2:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))

			fn := fmt.Sprintf("%v.%v.json", "fail", name)
			scene.Save(path.Join("./ai2_output", fn))

			break
		}
	}
}

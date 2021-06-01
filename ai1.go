package block7

import "go.uber.org/zap"

func ai1L0(scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
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
						Warn("AI1:Click:L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						Info("AI1:Click:L0",
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
					Warn("AI1:Click:L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					clicknums++
				}

				if gs != GameStateRunning {
					Info("AI1:Click:L0",
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

func ai1L1(scene *Scene, mapbi *BlockInfoMap, symbol int) int {
	clicknums := 0
	v, isok := mapbi.MapBlockInfo[symbol]
	if !isok {
		return 0
	}

	if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
		cn := 0
		for _, b := range v.LevelList[0] {
			cn++
			clicknums++

			gs, isok := scene.Click(b.X, b.Y, b.Z)
			if !isok {
				Warn("AI1:Click:L1L0",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			}

			if gs != GameStateRunning {
				Info("AI1:Click:L1L0",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z),
					zap.Int("gameState", gs))

				return -1
			}
		}

		for _, b := range v.LevelList[1] {
			cn++
			clicknums++

			if len(b.Parent) > 0 {
				for _, bc := range b.Parent {
					gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
					if !isok {
						Warn("AI1:Click:L1:Parent",
							zap.Int("x", bc.X),
							zap.Int("y", bc.Y),
							zap.Int("z", bc.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						Info("AI1:Click:L1:Parent",
							zap.Int("x", bc.X),
							zap.Int("y", bc.Y),
							zap.Int("z", bc.Z),
							zap.Int("gameState", gs))

						return -1
					}
				}
			}

			gs, isok := scene.Click(b.X, b.Y, b.Z)
			if !isok {
				Warn("AI1:Click:L1",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			}

			if gs != GameStateRunning {
				Info("AI1:Click:L1",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z),
					zap.Int("gameState", gs))

				return -1
			}

			if cn == BlockNums {
				break
			}
		}
	}

	return clicknums
}

func AI1(scene *Scene) {
	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		if len(scene.Block) > 0 {
			for _, v := range scene.Block {
				cn, isbreak := ai1L0(scene, mapbi, v.Symbol)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for _, v := range scene.Block {
					cn := ai1L1(scene, mapbi, v.Symbol)
					if cn > 0 {
						clicknums += cn

						break
					} else if cn < 0 {
						break
					}
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn, isbreak := ai1L0(scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}

				if clicknums == 0 {
					for k := range mapbi.MapBlockInfo {
						cn := ai1L1(scene, mapbi, k)
						if cn > 0 {
							clicknums += cn

							break
						} else if cn < 0 {
							break
						}
					}
				}
			}
		} else {
			for k := range mapbi.MapBlockInfo {
				cn, isbreak := ai1L0(scene, mapbi, k)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for k := range mapbi.MapBlockInfo {
					cn := ai1L1(scene, mapbi, k)
					if cn > 0 {
						clicknums += cn

						break
					} else if cn < 0 {
						break
					}
				}
			}
		}

		if clicknums > 0 {
			Info("AI1:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()))
		} else {
			Info("AI1:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()))

			// break
		}
	}
}

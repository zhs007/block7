package block7game

import (
	"fmt"
	"os"
	"path"

	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

const AI1OutputPath = "./ai1_output"

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
						goutils.Warn("AI1:Click:L0",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						goutils.Info("AI1:Click:L0",
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
					goutils.Warn("AI1:Click:L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					clicknums++
				}

				if gs != GameStateRunning {
					goutils.Info("AI1:Click:L0",
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

func ai1L1(scene *Scene, mapbi *BlockInfoMap, symbol int) (int, bool) {
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
					goutils.Warn("AI1:Click:L1L0",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					cn++
					clicknums++
				}

				if gs != GameStateRunning {
					goutils.Info("AI1:Click:L1L0",
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
							goutils.Warn("AI1:Click:L1:Parent",
								zap.Int("x", bc.X),
								zap.Int("y", bc.Y),
								zap.Int("z", bc.Z))
						} else {
							clicknums++
						}

						if gs != GameStateRunning {
							goutils.Info("AI1:Click:L1:Parent",
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
					goutils.Warn("AI1:Click:L1",
						zap.Int("x", b.X),
						zap.Int("y", b.Y),
						zap.Int("z", b.Z))
				} else {
					cn++
					clicknums++
				}

				if gs != GameStateRunning {
					goutils.Info("AI1:Click:L1",
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
				goutils.Warn("AI1:Click:L1L0",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			} else {
				cn++
				clicknums++
			}

			if gs != GameStateRunning {
				goutils.Info("AI1:Click:L1L0",
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
						goutils.Warn("AI1:Click:L1:Parent",
							zap.Int("x", bc.X),
							zap.Int("y", bc.Y),
							zap.Int("z", bc.Z))
					} else {
						clicknums++
					}

					if gs != GameStateRunning {
						goutils.Info("AI1:Click:L1:Parent",
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
				goutils.Warn("AI1:Click:L1",
					zap.Int("x", b.X),
					zap.Int("y", b.Y),
					zap.Int("z", b.Z))
			} else {
				cn++
				clicknums++
			}

			if gs != GameStateRunning {
				goutils.Info("AI1:Click:L1",
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

func AI1(scene *Scene, name string) {
	os.MkdirAll(AI1OutputPath, os.ModePerm)

	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		if len(scene.Block) > 0 {
			for _, b := range scene.Block {
				cn, isbreak := ai1L0(scene, mapbi, b.Symbol)
				clicknums += cn
				if isbreak {
					break
				}
			}

			for _, b := range scene.Block {
				cn, isbreak := ai1L1(scene, mapbi, b.Symbol)
				clicknums += cn
				if isbreak {
					break
				}
			}
		}

		if clicknums == 0 && len(mapbi.BlockSymbols) > 0 {
			for _, v := range mapbi.BlockSymbols {
				cn, isbreak := ai1L0(scene, mapbi, v)
				clicknums += cn
				if isbreak {
					break
				}
			}

			if clicknums == 0 {
				for _, v := range mapbi.BlockSymbols {
					cn, isbreak := ai1L1(scene, mapbi, v)
					clicknums += cn
					if isbreak {
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
						cn, isbreak := ai1L1(scene, mapbi, k)
						clicknums += cn
						if isbreak {
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
					cn, isbreak := ai1L1(scene, mapbi, k)
					clicknums += cn
					if isbreak {
						break
					}
				}
			}
		}

		if scene.CountSymbols() == 0 {
			fn := fmt.Sprintf("%v.%v.json", "ok", name)
			scene.Save(path.Join(AI1OutputPath, fn))

			break
		}

		if clicknums > 0 {
			goutils.Info("AI1:Turn",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))
		} else {
			goutils.Info("AI1:Turn:fail",
				zap.Int("iturn", iturn),
				zap.Int("clicknums", clicknums),
				zap.Int("blocknums", scene.CountSymbols()),
				zap.Int("block", len(scene.Block)))

			fn := fmt.Sprintf("%v.%v.json", "fail", name)
			scene.Save(path.Join(AI1OutputPath, fn))

			break
		}
	}
}

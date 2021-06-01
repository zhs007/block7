package block7

import "go.uber.org/zap"

func AI1(scene *Scene) {
	iturn := 0
	for {
		iturn++

		clicknums := 0
		mapbi := scene.Analysis()
		for _, v := range mapbi.MapBlockInfo {
			if len(v.LevelList[0]) >= BlockNums {
				cn := len(v.LevelList[0]) / BlockNums
				for i, b := range v.LevelList[0] {
					if i >= cn*BlockNums {
						break
					}

					clicknums++

					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						Warn("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					}

					if gs != GameStateRunning {
						Info("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return
					}
				}
			} else if len(v.LevelList[0])+len(v.LevelList[1]) >= BlockNums {
				cn := 0
				for _, b := range v.LevelList[0] {
					cn++
					clicknums++

					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						Warn("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					}

					if gs != GameStateRunning {
						Info("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return
					}
				}

				for _, b := range v.LevelList[1] {
					cn++
					clicknums++

					if len(b.Parent) > 0 {
						for _, bc := range b.Parent {
							clicknums++

							gs, isok := scene.Click(bc.X, bc.Y, bc.Z)
							if !isok {
								Warn("AI1:Click",
									zap.Int("x", bc.X),
									zap.Int("y", bc.Y),
									zap.Int("z", bc.Z))
							}

							if gs != GameStateRunning {
								Info("AI1:Click",
									zap.Int("x", bc.X),
									zap.Int("y", bc.Y),
									zap.Int("z", bc.Z),
									zap.Int("gameState", gs))

								return
							}
						}
					}

					gs, isok := scene.Click(b.X, b.Y, b.Z)
					if !isok {
						Warn("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z))
					}

					if gs != GameStateRunning {
						Info("AI1:Click",
							zap.Int("x", b.X),
							zap.Int("y", b.Y),
							zap.Int("z", b.Z),
							zap.Int("gameState", gs))

						return
					}

					if cn == BlockNums {
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

			break
		}
	}
}

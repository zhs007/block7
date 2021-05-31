package block7

// Scene - scene
type Scene struct {
	Width        int
	Height       int
	Layers       int
	XOff         int
	YOff         int
	Arr          [][][]int
	Block        []*BlockData
	MaxBlockNums int
}

// NewScene - new a scene
func NewScene(rng Rng, stage *Stage, symbols []int, blockNums int) (*Scene, error) {
	ss, err := genSymbols(rng, symbols, stage.IconNums)
	if err != nil {
		return nil, err
	}

	scene := &Scene{
		Width:        stage.Width,
		Height:       stage.Height,
		Layers:       len(stage.Layer),
		XOff:         1,
		YOff:         1,
		MaxBlockNums: blockNums,
	}

	for _, arrlayer := range stage.Layer {
		arrslayer := [][]int{}
		for _, arrrow := range arrlayer {
			arrsrow := []int{}
			for _, v := range arrrow {
				if v == 0 {
					arrsrow = append(arrsrow, 0)
				} else {
					nss, cs, err := randSymbols(rng, ss)
					if err != nil {
						return nil, err
					}

					arrsrow = append(arrsrow, cs)
					ss = nss
				}
			}

			arrslayer = append(arrslayer, arrsrow)
		}

		scene.Arr = append(scene.Arr, arrslayer)
	}

	return scene, nil
}

func (scene *Scene) CountSymbols() int {
	n := 0
	for _, arrlayer := range scene.Arr {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v > 0 {
					n++
				}
			}
		}
	}

	return n
}

func (scene *Scene) CountSymbol(symbol int) int {
	n := 0
	for _, arrlayer := range scene.Arr {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v == symbol {
					n++
				}
			}
		}
	}

	return n
}

func (scene *Scene) CanClick(x, y, z int) bool {
	if x < 0 || x >= scene.Width {
		return false
	}

	if x < 0 || y >= scene.Height {
		return false
	}

	if z < 0 || z >= len(scene.Arr) {
		return false
	}

	if z == len(scene.Arr)-1 {
		return scene.Arr[z][y][x] > 0
	}

	for zi := 1; z+zi < len(scene.Arr); zi++ {
		if zi%1 == 1 {
			if scene.Arr[z+zi][y][x] > 0 || scene.Arr[z+zi][y][x+scene.XOff] > 0 || scene.Arr[z+zi][y+scene.YOff][x] > 0 || scene.Arr[z+zi][y+scene.YOff][x+scene.XOff] > 0 {
				return false
			}
		} else {
			if scene.Arr[z+zi][y][x] > 0 {
				return false
			}
		}
	}

	return scene.Arr[z][y][x] > 0
}

func (scene *Scene) CanClickEx(x, y, z int, lst []*BlockData) bool {
	if x < 0 || x >= scene.Width {
		return false
	}

	if x < 0 || y >= scene.Height {
		return false
	}

	if z < 0 || z >= len(scene.Arr) {
		return false
	}

	if z == len(scene.Arr)-1 {
		return scene.Arr[z][y][x] > 0
	}

	for zi := 1; z+zi < len(scene.Arr); zi++ {
		if zi%1 == 1 {
			if (scene.Arr[z+zi][y][x] > 0 && !HasBlockData(lst, x, y, z+zi)) ||
				(scene.Arr[z+zi][y][x+scene.XOff] > 0 && !HasBlockData(lst, x+scene.XOff, y, z+zi)) ||
				(scene.Arr[z+zi][y+scene.YOff][x] > 0 && !HasBlockData(lst, x, y+scene.YOff, z+zi)) ||
				(scene.Arr[z+zi][y+scene.YOff][x+scene.XOff] > 0 && !HasBlockData(lst, x+scene.XOff, y+scene.YOff, z+zi)) {
				return false
			}
		} else {
			if scene.Arr[z+zi][y][x] > 0 && !HasBlockData(lst, x, y, z+zi) {
				return false
			}
		}
	}

	return scene.Arr[z][y][x] > 0
}

func (scene *Scene) GetLevel1(x, y, z int) []*BlockData {
	if z == 0 {
		return nil
	}

	arr := []*BlockData{}

	return arr
}

func (scene *Scene) GetMaxZ(x, y int) int {
	if len(scene.Arr) == 1 {
		return 0
	}

	for z := len(scene.Arr) - 1; z >= 0; z-- {
		if scene.Arr[z][y][x] > 0 {
			return z
		}
	}

	return 0
}

func (scene *Scene) Analysis() *BlockInfoMap {
	mapBI := NewBlockInfoMap(DefaultMaxBlockLevel)

	for x := 0; x < scene.Width; x++ {
		for y := 0; y < scene.Height; y++ {
			mz := scene.GetMaxZ(x, y)
			if scene.CanClick(x, y, mz) {
				arr := []*BlockData{NewBlockData(x, y, mz, scene.Arr[mz][y][x])}

				mapBI.AddBlockData(arr[0], 0)

				if mz > 0 {
					if scene.CanClickEx(x, y, mz-1, arr) {
						mapBI.AddBlockDataEx(x, y, mz-1, scene.Arr[mz-1][y][x], 1)
					}

					if scene.CanClickEx(x-scene.XOff, y, mz-1, arr) {
						mapBI.AddBlockDataEx(x-scene.XOff, y, mz-1, scene.Arr[mz-1][y][x-scene.XOff], 1)
					}

					if scene.CanClickEx(x, y-scene.YOff, mz-1, arr) {
						mapBI.AddBlockDataEx(x, y-scene.YOff, mz-1, scene.Arr[mz-1][y-scene.YOff][x], 1)
					}

					if scene.CanClickEx(x-scene.XOff, y-scene.YOff, mz-1, arr) {
						mapBI.AddBlockDataEx(x-scene.XOff, y-scene.YOff, mz-1, scene.Arr[mz-1][y-scene.YOff][x-scene.XOff], 1)
					}
				}
			}
		}
	}

	return mapBI
}

// Click - return gamestate, isok
func (scene *Scene) Click(x, y, z int) (int, bool) {
	if !scene.CanClick(x, y, z) {
		return GameStateRunning, false
	}

	if len(scene.Block) >= scene.MaxBlockNums {
		return GameStateFail, false
	}

	b := NewBlockData(x, y, z, scene.Arr[z][y][x])
	scene.Arr[z][y][x] = 0

	scene.Block = insBlockDataAndProc(scene.Block, b)

	if len(scene.Block) >= scene.MaxBlockNums {
		return GameStateFail, true
	}

	if scene.CountSymbols() == 0 {
		return GameStateSucess, true
	}

	return GameStateRunning, true
}

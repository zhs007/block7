package block7

// Scene - scene
type Scene struct {
	Width  int
	Height int
	Layers int
	XOff   int
	YOff   int
	Arr    [][][]int
}

// NewScene - new a scene
func NewScene(rng Rng, stage *Stage, symbols []int) (*Scene, error) {
	ss, err := genSymbols(rng, symbols, stage.IconNums)
	if err != nil {
		return nil, err
	}

	scene := &Scene{
		Width:  stage.Width,
		Height: stage.Height,
		Layers: len(stage.Layer),
		XOff:   1,
		YOff:   1,
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
	mapBI := NewBlockInfoMap()

	for x := 0; x < scene.Width; x++ {
		for y := 0; y < scene.Height; y++ {
			mz := scene.GetMaxZ(x, y)
			if scene.CanClick(x, y, mz) {
				mapBI.AddBlockDataEx(x, y, mz, scene.Arr[mz][y][x])
			}
		}
	}

	return mapBI
}

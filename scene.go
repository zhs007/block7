package block7

// Scene - scene
type Scene struct {
	Width     int
	Height    int
	Layers    int
	BlockNums int
	Arr       [][][]int
}

// NewScene - new a scene
func NewScene(rng Rng, stage *Stage, symbols []int) (*Scene, error) {
	ss, err := genSymbols(rng, symbols, stage.IconNums)
	if err != nil {
		return nil, err
	}

	scene := &Scene{
		Width:     stage.Width,
		Height:    stage.Height,
		Layers:    len(stage.Layer),
		BlockNums: stage.IconNums,
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

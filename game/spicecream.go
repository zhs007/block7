package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type icecreamData struct {
	pos [][]int
}

// SpecialIceCream - icecream
type SpecialIceCream struct {
	specialID   int
	specialType int
	spoonID     int
	layer       int
	w           int
	h           int
}

func NewIceCream(specialid int, icecreamid int, spoonid int, w, h int) *SpecialIceCream {
	return &SpecialIceCream{
		specialID:   specialid,
		specialType: icecreamid,
		spoonID:     spoonid,
		layer:       2,
		w:           w,
		h:           h,
	}
}

// GetSpecialID - GetSpecialID
func (icecream *SpecialIceCream) GetSpecialID() int {
	return icecream.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (icecream *SpecialIceCream) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	for i := 0; i < std.Nums; i++ {
		arr = append(arr, icecream.spoonID, icecream.spoonID, icecream.spoonID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (icecream *SpecialIceCream) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	// lst, err := GetAllBlocksEx(scene, icecream.w, icecream.h, func(x, y, z int) bool {
	// 	if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
	// 		return false
	// 	}

	// 	return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, icecream.layer)
	// }, func(x, y, z int, w, h int) bool {
	// 	nums := scene.CountChildrenNumsEx(x, y, z, w, h, func(x, y, z int) bool {
	// 		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
	// 			return false
	// 		}

	// 		return scene.InitArr[z][y][x] > 0
	// 	})

	// 	return scene.BlockNums-nums >= std.Nums*3
	// })
	// if err != nil {
	// 	goutils.Error("SpecialIceCream.OnFixScene:GetAllBlocksEx",
	// 		zap.Error(err))

	// 	return err
	// }

	// stddata := &icecreamData{}

	// for i := 0; i < std.Nums; i++ {
	// 	cr, err := rng.Rand(len(lst))
	// 	if err != nil {
	// 		goutils.Error("SpecialIceCream.OnFixScene:GetAllBlocksEx",
	// 			zap.Int("i", i),
	// 			zap.Error(err))

	// 		return err
	// 	}

	// 	stddata.pos = append(stddata.pos, lst[cr])

	// 	lst = append(lst[:cr], lst[:cr+1]...)
	// }

	// std.Data = stddata

	// lstSpooon, err := GetAllBlocks(scene, func(x, y, z int) bool {
	// 	if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
	// 		return false
	// 	}

	// 	return scene.InitArr[z][y][x] == icecream.spoonID
	// })
	// if err != nil {
	// 	goutils.Error("SpecialIceCream.OnFixScene:GetAllBlocks",
	// 		zap.Error(err))

	// 	return err
	// }

	// nlst, err := AddBlocksVal(lst, func(x, y, z int) int {
	// 	blockval := 0
	// 	for ox := 0; ox < icecream.w; ox++ {
	// 		for oy := 0; oy < icecream.h; oy++ {
	// 			cbd := &BlockData{X: x + ox, Y: y + oy, Z: z}

	// 			for _, spoonpos := range lstSpooon {
	// 				if scene.IsParent2(cbd, &BlockData{X: spoonpos[0], Y: spoonpos[1], Z: spoonpos[2]}, func(x, y, z int) bool {
	// 					return scene.InitArr[z][y][x] > 0
	// 				}) {
	// 					blockval++
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return 0
	// })
	// if err != nil {
	// 	goutils.Error("SpecialIceCream.OnFixScene:AddBlocksVal",
	// 		zap.Error(err))

	// 	return err
	// }

	// sort.Slice(nlst, func(i, j int) bool {
	// 	return nlst[i][3] > nlst[j][3]
	// })

	// lst := FindAllSymbolsEx(scene.InitArr, []int{icecream.icecreamID, icecream.spoonID})
	// if len(lst[0]) > 0 && len(lst[1]) > 0 {
	// 	icecream.fixScene(scene, lst)
	// }

	return nil
}

// fixScene - fixScene
func (icecream *SpecialIceCream) fixSpoon(rng IRng, scene *Scene, lstchild []int) error {
	lstValid := CheckScene(scene, func(x, y, z int) bool {
		ci := goutils.FindInt3(lstchild, x, y, z)
		if ci < 0 {
			if scene.InitArr[z][y][x] != icecream.spoonID {
				return true
			}
		}

		return false
	})

	for i := 0; i < len(lstchild)/3; i++ {
		if scene.InitArr[lstchild[i*3+2]][lstchild[i*3+1]][lstchild[i*3]] == icecream.spoonID {
			if len(lstValid) <= 0 {
				goutils.Error("SpecialIceCream.fixSpoon:",
					zap.Error(ErrInvalidSafeList))

				return ErrInvalidSafeList
			}

			cr, err := rng.Rand(len(lstValid))
			if err != nil {
				goutils.Error("SpecialIceCream.fixSpoon:Rand",
					zap.Error(err))

				return err
			}

			scene.InitArr[lstchild[i*3+2]][lstchild[i*3+1]][lstchild[i*3]] = scene.InitArr[lstValid[cr][2]][lstValid[cr][1]][lstValid[cr][0]]
			scene.InitArr[lstValid[cr][2]][lstValid[cr][1]][lstValid[cr][0]] = icecream.spoonID

			scene.Arr[lstchild[i*3+2]][lstchild[i*3+1]][lstchild[i*3]] = scene.Arr[lstValid[cr][2]][lstValid[cr][1]][lstValid[cr][0]]
			scene.Arr[lstValid[cr][2]][lstValid[cr][1]][lstValid[cr][0]] = icecream.spoonID

			lstValid = append(lstValid[:cr], lstValid[cr+1:]...)
		}
	}

	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (icecream *SpecialIceCream) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     icecream.layer,
		LayerType: icecream.specialType,
		Special:   icecream,
	}

	lst, err := GetAllBlocksEx(scene, icecream.w, icecream.h, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 &&
			scene.InitArr[z][y][x] != 403 &&
			scene.InitArr[z][y][x] != icecream.spoonID &&
			!scene.HasSpecialLayer(x, y, z, icecream.layer)
	}, func(x, y, z int, w, h int) bool {
		nums := scene.CountChildrenNumsEx(x, y, z, w, h, func(x, y, z int) bool {
			if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
				return false
			}

			return scene.InitArr[z][y][x] > 0
		})

		return scene.BlockNums-nums >= std.Nums*3
	})
	if err != nil {
		goutils.Error("SpecialIceCream.OnGenSymbolLayers:GetAllBlocksEx",
			zap.Error(err))

		return nil, err
	}

	stddata := &icecreamData{}

	for i := 0; i < std.Nums; i++ {
		cr, err := rng.Rand(len(lst))
		if err != nil {
			goutils.Error("SpecialIceCream.OnGenSymbolLayers:GetAllBlocksEx",
				zap.Int("i", i),
				zap.Error(err))

			return nil, err
		}

		stddata.pos = append(stddata.pos, lst[cr])

		lst = append(lst[:cr], lst[cr+1:]...)
	}

	// stddata, isok := std.Data.(*icecreamData)
	// if isok {
	sl.Pos = stddata.pos
	// }

	lstchild := GetChildrenEx(scene, sl.Pos, icecream.w, icecream.h, func(x, y, z int) bool {
		return scene.InitArr[z][y][x] > 0
	})

	err = icecream.fixSpoon(rng, scene, lstchild)
	if err != nil {
		goutils.Error("SpecialIceCream.OnGenSymbolLayers:fixSpoon",
			zap.Error(err))

		return nil, err
	}

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (icecream *SpecialIceCream) GetSpecialLayerType() int {
	return icecream.specialType
}

// OnGen2 - OnGen2
func (icecream *SpecialIceCream) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     icecream.layer,
		LayerType: icecream.specialType,
		Special:   icecream,
		Pos: [][]int{
			{x, y, z},
		},
	}

	return sl, nil
}

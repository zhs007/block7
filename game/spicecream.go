package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

type icecreamData struct {
	pos [][]int
}

// SpecialIceCream - icecream
type SpecialIceCream struct {
	specialID  int
	icecreamID int
	spoonID    int
	layer      int
	w          int
	h          int
}

func NewIceCream(specialid int, icecreamid int, spoonid int, w, h int) *SpecialIceCream {
	return &SpecialIceCream{
		specialID:  specialid,
		icecreamID: icecreamid,
		spoonID:    spoonid,
		layer:      2,
		w:          w,
		h:          h,
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
func (icecream *SpecialIceCream) OnFixScene(std *SpecialTypeData, scene *Scene) error {
	_, err := GetAllBlocksEx(scene, icecream.w, icecream.h, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, icecream.layer)
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
		block7utils.Error("SpecialIceCream.OnFixScene:GetAllBlocksEx",
			zap.Error(err))

		return err
	}

	// lstSpooon, err := GetAllBlocks(scene, func(x, y, z int) bool {
	// 	if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
	// 		return false
	// 	}

	// 	return scene.InitArr[z][y][x] == icecream.spoonID
	// })
	// if err != nil {
	// 	block7utils.Error("SpecialIceCream.OnFixScene:GetAllBlocks",
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
	// 	block7utils.Error("SpecialIceCream.OnFixScene:AddBlocksVal",
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

// // fixScene - fixScene
// func (icecream *SpecialIceCream) fixScene(scene *Scene, lst [][]*BlockData) {
// 	for _, fv := range lst[1] {
// 		for _, cv := range lst[0] {
// 			if scene.IsParent2(fv, cv, func(x, y, z int) bool {
// 				return scene.InitArr[z][y][x] > 0
// 			}) {
// 				block7utils.Debug("SpecialCake.fixScene",
// 					block7utils.JSON("icecream", cv),
// 					block7utils.JSON("spoon", fv))

// 				scene.InitArr[fv.Z][fv.Y][fv.X] = icecream.icecreamID
// 				scene.InitArr[cv.Z][cv.Y][cv.X] = icecream.spoonID

// 				tx := fv.X
// 				fv.X = cv.X
// 				cv.X = tx

// 				ty := fv.Y
// 				fv.Y = cv.Y
// 				cv.Y = ty

// 				tz := fv.Z
// 				fv.Z = cv.Z
// 				cv.Z = tz

// 				icecream.fixScene(scene, lst)
// 			}
// 		}
// 	}
// }

// OnGenSymbolLayer - OnGenSymbolLayer
func (icecream *SpecialIceCream) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (icecream *SpecialIceCream) GetSpecialLayerType() int {
	return 0
}

package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// SpecialCurtain - curtain
type SpecialCurtain struct {
	specialID   int
	specialType int
	layer       int
}

func NewSpecialCurtain(specialid int, curtain int) *SpecialCurtain {
	return &SpecialCurtain{
		specialID:   specialid,
		specialType: curtain,
		layer:       2,
	}
}

// GetSpecialID - GetSpecialID
func (curtain *SpecialCurtain) GetSpecialID() int {
	return curtain.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (curtain *SpecialCurtain) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (curtain *SpecialCurtain) OnFixScene(scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (curtain *SpecialCurtain) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     curtain.layer,
		LayerType: curtain.specialType,
		Special:   curtain,
	}

	lst, err := GenBlocks(rng, scene, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	}, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, curtain.layer)
	})
	if err != nil {
		block7utils.Error("SpecialCurtain.OnGenSymbolLayers:GenBlocks",
			block7utils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (curtain *SpecialCurtain) GetSpecialLayerType() int {
	return curtain.specialType
}

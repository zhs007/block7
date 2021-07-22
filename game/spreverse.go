package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// SpecialReverse - reverse
type SpecialReverse struct {
	specialID   int
	specialType int
	layer       int
}

func NewReverse(specialid int, iceid int) *SpecialReverse {
	return &SpecialReverse{
		specialID:   specialid,
		specialType: iceid,
		layer:       1,
	}
}

// GetSpecialID - GetSpecialID
func (reverse *SpecialReverse) GetSpecialID() int {
	return reverse.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (reverse *SpecialReverse) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (reverse *SpecialReverse) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (reverse *SpecialReverse) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     reverse.layer,
		LayerType: reverse.specialType,
		Special:   reverse,
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

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, reverse.layer)
	})
	if err != nil {
		block7utils.Error("SpecialReverse.OnGenSymbolLayers:GenBlocks",
			block7utils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (weeds *SpecialReverse) GetSpecialLayerType() int {
	return weeds.specialType
}

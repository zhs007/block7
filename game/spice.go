package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// SpecialIce - ice
type SpecialIce struct {
	specialID   int
	specialType int
	layer       int
	brother     int
}

func NewIce(specialid int, iceid int, brother int) *SpecialIce {
	return &SpecialIce{
		specialID:   specialid,
		specialType: iceid,
		layer:       1,
		brother:     brother,
	}
}

// GetSpecialID - GetSpecialID
func (ice *SpecialIce) GetSpecialID() int {
	return ice.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (ice *SpecialIce) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (ice *SpecialIce) OnFixScene(scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (ice *SpecialIce) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     ice.layer,
		LayerType: ice.specialType,
		Special:   ice,
	}

	lst, err := GenBrotherBlocks(rng, scene, ice.brother, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	})
	if err != nil {
		block7utils.Error("SpecialIce.OnGenSymbolLayers:GenBrotherBlocks",
			block7utils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (ice *SpecialIce) GetSpecialLayerType() int {
	return ice.specialType
}

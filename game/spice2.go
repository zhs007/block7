package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialIce2 - ice2
type SpecialIce2 struct {
	specialID   int
	specialType int
	layer       int
	brother     int
}

func NewIce2(specialid int, iceid int, brother int) *SpecialIce2 {
	return &SpecialIce2{
		specialID:   specialid,
		specialType: iceid,
		layer:       1,
		brother:     brother,
	}
}

// GetSpecialID - GetSpecialID
func (ice *SpecialIce2) GetSpecialID() int {
	return ice.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (ice *SpecialIce2) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (ice *SpecialIce2) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (ice *SpecialIce2) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     ice.layer,
		LayerType: ice.specialType,
		Special:   ice,
	}

	lst, err := GenBrotherBlocks(rng, scene, ice.brother, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	}, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, ice.layer)
	})
	if err != nil {
		goutils.Error("SpecialIce.OnGenSymbolLayers:GenBrotherBlocks",
			goutils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (ice *SpecialIce2) GetSpecialLayerType() int {
	return ice.specialType
}

// OnGen2 - OnGen2
func (ice *SpecialIce2) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     ice.layer,
		LayerType: ice.specialType,
		Special:   ice,
		Pos: [][]int{
			{x, y, z},
		},
	}

	return sl, nil
}

package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialWeeds - weeds
type SpecialWeeds struct {
	specialID   int
	specialType int
	layer       int
}

func NewWeeds(specialid int, iceid int) *SpecialWeeds {
	return &SpecialWeeds{
		specialID:   specialid,
		specialType: iceid,
		layer:       1,
	}
}

// GetSpecialID - GetSpecialID
func (weeds *SpecialWeeds) GetSpecialID() int {
	return weeds.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (weeds *SpecialWeeds) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (weeds *SpecialWeeds) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (weeds *SpecialWeeds) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     weeds.layer,
		LayerType: weeds.specialType,
		Special:   weeds,
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

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, weeds.layer)
	})
	if err != nil {
		goutils.Error("SpecialWeeds.OnGenSymbolLayers:GenBlocks",
			goutils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (weeds *SpecialWeeds) GetSpecialLayerType() int {
	return weeds.specialType
}

// OnGen2 - OnGen2
func (weeds *SpecialWeeds) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	return nil, nil
}

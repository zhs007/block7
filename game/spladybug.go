package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialLadybug - ladybug
type SpecialLadybug struct {
	specialID   int
	specialType int
	layer       int
	brother     int
}

func NewLadybug(specialid int, ladybug int, brother int) *SpecialLadybug {
	return &SpecialLadybug{
		specialID:   specialid,
		specialType: ladybug,
		layer:       3,
		brother:     brother,
	}
}

// GetSpecialID - GetSpecialID
func (ladybug *SpecialLadybug) GetSpecialID() int {
	return ladybug.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (ladybug *SpecialLadybug) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (ladybug *SpecialLadybug) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (ladybug *SpecialLadybug) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     ladybug.layer,
		LayerType: ladybug.specialType,
		Special:   ladybug,
	}

	lst, err := GenBrotherBlocks(rng, scene, ladybug.brother, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	}, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, ladybug.layer)
	})
	if err != nil {
		goutils.Error("SpecialLadybug.OnGenSymbolLayers:GenBrotherBlocks",
			goutils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (ladybug *SpecialLadybug) GetSpecialLayerType() int {
	return ladybug.specialType
}

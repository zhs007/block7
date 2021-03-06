package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialRainbow - rainbow
type SpecialRainbow struct {
	specialID int
	rainbowID int
}

func NewRainbow(specialid int, rainbow int) *SpecialRainbow {
	return &SpecialRainbow{
		specialID: specialid,
		rainbowID: rainbow,
	}
}

// GetSpecialID - GetSpecialID
func (rainbow *SpecialRainbow) GetSpecialID() int {
	return rainbow.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (rainbow *SpecialRainbow) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidRainbowNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, rainbow.rainbowID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (rainbow *SpecialRainbow) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (rainbow *SpecialRainbow) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (rainbow *SpecialRainbow) GetSpecialLayerType() int {
	return 0
}

// OnGen2 - OnGen2
func (rainbow *SpecialRainbow) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	if scene.InitArr[z][y][x] > 0 {
		goutils.Error("SpecialRainbow:OnGen2",
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Int("z", z),
			zap.Error(ErrRecoveBlock))

		return nil, ErrRecoveBlock
	}

	scene.InitArr[z][y][x] = rainbow.rainbowID

	return nil, nil
}

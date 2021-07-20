package block7game

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
		return nil, ErrInvalidBombNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, rainbow.rainbowID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (rainbow *SpecialRainbow) OnFixScene(scene *Scene) error {
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

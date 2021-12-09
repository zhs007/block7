package block7game

// SpecialBomb - bomb
type SpecialBomb struct {
	specialID int
	bombID    int
}

func NewBomb(specialid int, bombid int) *SpecialBomb {
	return &SpecialBomb{
		specialID: specialid,
		bombID:    bombid,
	}
}

// GetSpecialID - GetSpecialID
func (bomb *SpecialBomb) GetSpecialID() int {
	return bomb.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (bomb *SpecialBomb) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidBombNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, bomb.bombID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (bomb *SpecialBomb) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (bomb *SpecialBomb) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (bomb *SpecialBomb) GetSpecialLayerType() int {
	return 0
}

// OnGen2 - OnGen2
func (bomb *SpecialBomb) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	scene.InitArr[z][y][x] = bomb.bombID

	return nil, nil
}

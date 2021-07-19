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
func (bomb *SpecialBomb) OnGenSymbolBlocks(std SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidBombNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, bomb.bombID)
	}

	return arr, nil
}

// // OnGenSymbolLayers - OnGenSymbolLayers
// func (bomb *SpecialBomb) OnGenSymbolLayers(std SpecialTypeData, arr []int) ([]int, error) {

// }

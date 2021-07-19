package block7game

// SpecialCake - cake
type SpecialCake struct {
	specialID int
	cakeID    int
	forkID    int
}

func NewCake(specialid int, cakeid int, forkid int) *SpecialCake {
	return &SpecialCake{
		specialID: specialid,
		cakeID:    cakeid,
		forkID:    forkid,
	}
}

// GetSpecialID - GetSpecialID
func (cake *SpecialCake) GetSpecialID() int {
	return cake.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (cake *SpecialCake) OnGenSymbolBlocks(std SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidBombNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, cake.cakeID)
		arr = append(arr, cake.forkID, cake.forkID, cake.forkID)
	}

	return arr, nil
}

// // OnGenSymbolLayers - OnGenSymbolLayers
// func (bomb *SpecialBomb) OnGenSymbolLayers(std SpecialTypeData, arr []int) ([]int, error) {

// }

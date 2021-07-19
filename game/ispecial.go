package block7game

// ISpecial - spicial
type ISpecial interface {
	// GetSpecialID - GetSpecialID
	GetSpecialID() int
	// OnGenSymbolBlocks - OnGenSymbolBlocks
	OnGenSymbolBlocks(std SpecialTypeData, arr []int) ([]int, error)
	// // OnGenSymbolLayers - OnGenSymbolLayers
	// OnGenSymbolLayers(std SpecialTypeData, arr []int) ([]int, error)
}

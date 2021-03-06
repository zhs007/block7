package block7game

// ISpecial - spicial
type ISpecial interface {
	// GetSpecialID - GetSpecialID
	GetSpecialID() int
	// GetSpecialLayerType - GetSpecialLayerType
	GetSpecialLayerType() int
	// OnGenSymbolBlocks - OnGenSymbolBlocks
	OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error)
	// OnFixScene - OnFixScene
	OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error
	// OnGenSymbolLayer - OnGenSymbolLayer
	OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error)
	// OnGen2 - OnGen2
	OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error)
}

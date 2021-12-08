package block7game

// SpecialFork - fork
type SpecialFork struct {
	specialID int
	forkID    int
}

func NewFork(specialid int, forkid int) *SpecialFork {
	return &SpecialFork{
		specialID: specialid,
		forkID:    forkid,
	}
}

// GetSpecialID - GetSpecialID
func (fork *SpecialFork) GetSpecialID() int {
	return fork.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (fork *SpecialFork) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (fork *SpecialFork) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// fixScene - fixScene
func (fork *SpecialFork) fixScene(scene *Scene, lst [][]*BlockData) {
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (fork *SpecialFork) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (fork *SpecialFork) GetSpecialLayerType() int {
	return 0
}

// OnGen2 - OnGen2
func (fork *SpecialFork) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	scene.InitArr[z][y][x] = fork.forkID

	return nil, nil
}

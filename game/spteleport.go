package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialTeleport - teleport
type SpecialTeleport struct {
	specialID  int
	teleportID int
}

func NewTeleport(specialid int, teleport int) *SpecialTeleport {
	return &SpecialTeleport{
		specialID:  specialid,
		teleportID: teleport,
	}
}

// GetSpecialID - GetSpecialID
func (teleport *SpecialTeleport) GetSpecialID() int {
	return teleport.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (teleport *SpecialTeleport) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidTeleportNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, teleport.teleportID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (teleport *SpecialTeleport) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (teleport *SpecialTeleport) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (teleport *SpecialTeleport) GetSpecialLayerType() int {
	return 0
}

// OnGen2 - OnGen2
func (teleport *SpecialTeleport) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	if scene.InitArr[z][y][x] > 0 {
		goutils.Error("SpecialTeleport:OnGen2",
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Int("z", z),
			zap.Error(ErrRecoveBlock))

		return nil, ErrRecoveBlock
	}

	scene.InitArr[z][y][x] = teleport.teleportID

	return nil, nil
}

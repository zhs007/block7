package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

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
func (cake *SpecialCake) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	if std.Nums%3 > 0 {
		return nil, ErrInvalidCakeNums
	}

	for i := 0; i < std.Nums; i++ {
		arr = append(arr, cake.cakeID)
		arr = append(arr, cake.forkID, cake.forkID, cake.forkID)
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (cake *SpecialCake) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	lst := FindAllSymbolsEx(scene.InitArr, []int{cake.cakeID, cake.forkID})
	if len(lst[0]) > 0 && len(lst[1]) > 0 {
		cake.fixScene(scene, lst)
	}

	return nil
}

// fixScene - fixScene
func (cake *SpecialCake) fixScene(scene *Scene, lst [][]*BlockData) {
	for _, fv := range lst[1] {
		for _, cv := range lst[0] {
			if scene.IsParent2(fv, cv, func(x, y, z int) bool {
				return scene.InitArr[z][y][x] > 0
			}) {
				goutils.Debug("SpecialCake.fixScene",
					goutils.JSON("cake", cv),
					goutils.JSON("fork", fv))

				scene.InitArr[fv.Z][fv.Y][fv.X] = cake.cakeID
				scene.InitArr[cv.Z][cv.Y][cv.X] = cake.forkID

				tx := fv.X
				fv.X = cv.X
				cv.X = tx

				ty := fv.Y
				fv.Y = cv.Y
				cv.Y = ty

				tz := fv.Z
				fv.Z = cv.Z
				cv.Z = tz

				cake.fixScene(scene, lst)
			}
		}
	}
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (cake *SpecialCake) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	return nil, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (cake *SpecialCake) GetSpecialLayerType() int {
	return 0
}

// OnGen2 - OnGen2
func (cake *SpecialCake) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	if scene.InitArr[z][y][x] > 0 {
		goutils.Error("SpecialCake:OnGen2",
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Int("z", z),
			zap.Error(ErrRecoveBlock))

		return nil, ErrRecoveBlock
	}

	scene.InitArr[z][y][x] = cake.cakeID

	return nil, nil
}

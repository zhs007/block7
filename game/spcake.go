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

// OnFixScene - OnFixScene
func (cake *SpecialCake) OnFixScene(scene *Scene) error {
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
			if scene.IsParent(fv, cv) {
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

// // OnGenSymbolLayers - OnGenSymbolLayers
// func (bomb *SpecialBomb) OnGenSymbolLayers(std SpecialTypeData, arr []int) ([]int, error) {

// }

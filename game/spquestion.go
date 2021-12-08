package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// SpecialQuestion - question
type SpecialQuestion struct {
	specialID   int
	specialType int
	layer       int
}

func NewQuestion(specialid int, question int) *SpecialQuestion {
	return &SpecialQuestion{
		specialID:   specialid,
		specialType: question,
		layer:       0,
	}
}

// GetSpecialID - GetSpecialID
func (question *SpecialQuestion) GetSpecialID() int {
	return question.specialID
}

// OnGenSymbolBlocks - OnGenSymbolBlocks
func (question *SpecialQuestion) OnGenSymbolBlocks(std *SpecialTypeData, arr []int) ([]int, error) {
	return arr, nil
}

// OnFixScene - OnFixScene
func (question *SpecialQuestion) OnFixScene(rng IRng, std *SpecialTypeData, scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (question *SpecialQuestion) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     question.layer,
		LayerType: question.specialType,
		Special:   question,
	}

	lst, err := GenBlocks(rng, scene, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	}, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403 && !scene.HasSpecialLayer(x, y, z, question.layer)
	})
	if err != nil {
		goutils.Error("SpecialQuestion.OnGenSymbolLayers:GenBlocks",
			goutils.JSON("SpecialTypeData", std),
			zap.Error(err))

		return nil, err
	}

	sl.Pos = lst

	return sl, nil
}

// GetSpecialLayerType - GetSpecialLayerType
func (question *SpecialQuestion) GetSpecialLayerType() int {
	return question.specialType
}

// OnGen2 - OnGen2
func (question *SpecialQuestion) OnGen2(scene *Scene, x, y, z int) (*SpecialLayer, error) {
	return nil, nil
}

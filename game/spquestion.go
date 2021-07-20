package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// SpecialQuestion - question
type SpecialQuestion struct {
	specialID   int
	specialType int
	layer       int
	brother     int
}

func NewQuestion(specialid int, iceid int, brother int) *SpecialIce {
	return &SpecialIce{
		specialID:   specialid,
		specialType: iceid,
		layer:       1,
		brother:     brother,
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
func (question *SpecialQuestion) OnFixScene(scene *Scene) error {
	return nil
}

// OnGenSymbolLayer - OnGenSymbolLayer
func (question *SpecialQuestion) OnGenSymbolLayers(rng IRng, std *SpecialTypeData, scene *Scene) (*SpecialLayer, error) {
	sl := &SpecialLayer{
		Layer:     question.layer,
		LayerType: question.specialType,
		Special:   question,
	}

	lst, err := GenBrotherBlocks(rng, scene, question.brother, std.Nums, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0
	}, func(x, y, z int) bool {
		if x < 0 || y < 0 || z < 0 || x >= scene.Width || y >= scene.Height || z >= scene.Layers {
			return false
		}

		return scene.InitArr[z][y][x] > 0 && scene.InitArr[z][y][x] != 403
	})
	if err != nil {
		block7utils.Error("SpecialIce.OnGenSymbolLayers:GenBrotherBlocks",
			block7utils.JSON("SpecialTypeData", std),
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

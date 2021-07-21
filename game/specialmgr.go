package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

type SpecialMgr struct {
	MapSpecial map[int]ISpecial
}

func NewSpecialMgr() *SpecialMgr {
	return &SpecialMgr{
		MapSpecial: make(map[int]ISpecial),
	}
}

func (mgr *SpecialMgr) RegSpecial(specialid int, special ISpecial) {
	mgr.MapSpecial[specialid] = special
}

func (mgr *SpecialMgr) GenSymbols(ld2 *LevelData2) ([]int, error) {
	if ld2 == nil {
		return nil, nil
	}

	arr := []int{}

	for _, v := range ld2.SpecialType {
		sp, isok := mgr.MapSpecial[v.SpecialID]
		if isok {
			arr1, err := sp.OnGenSymbolBlocks(v, arr)
			if err != nil {
				block7utils.Error("SpecialMgr.GenSymbols:OnGenSymbolBlocks",
					zap.Int("SpecialID", v.SpecialID),
					block7utils.JSON("SpecialTypeData", v),
					zap.Error(err))

				return nil, err
			}

			arr = arr1
		}
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (mgr *SpecialMgr) OnFixScene(ld2 *LevelData2, scene *Scene) error {
	if ld2 == nil {
		return nil
	}

	for _, v := range ld2.SpecialType {
		sp, isok := mgr.MapSpecial[v.SpecialID]
		if isok {
			err := sp.OnFixScene(v, scene)
			if err != nil {
				block7utils.Error("SpecialMgr.OnFixScene:OnFixScene",
					zap.Int("SpecialID", v.SpecialID),
					block7utils.JSON("SpecialTypeData", v),
					zap.Error(err))

				return err
			}
		}
	}

	return nil
}

// GenSymbolLayers - GenSymbolLayers
func (mgr *SpecialMgr) GenSymbolLayers(rng IRng, ld2 *LevelData2, scene *Scene) error {
	if ld2 == nil {
		return nil
	}

	for _, v := range ld2.SpecialType {
		sp, isok := mgr.MapSpecial[v.SpecialID]
		if isok {
			layer, err := sp.OnGenSymbolLayers(rng, v, scene)
			if err != nil {
				block7utils.Error("SpecialMgr.GenSymbolLayers:OnGenSymbolLayers",
					zap.Int("SpecialID", v.SpecialID),
					block7utils.JSON("SpecialTypeData", v),
					zap.Error(err))

				return err
			}

			if layer != nil {
				scene.SpecialLayers = append(scene.SpecialLayers, layer)
			}
		}
	}

	return nil
}

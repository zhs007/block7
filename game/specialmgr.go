package block7game

import (
	goutils "github.com/zhs007/goutils"
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
				goutils.Error("SpecialMgr.GenSymbols:OnGenSymbolBlocks",
					zap.Int("SpecialID", v.SpecialID),
					goutils.JSON("SpecialTypeData", v),
					zap.Error(err))

				return nil, err
			}

			arr = arr1
		}
	}

	return arr, nil
}

// OnFixScene - OnFixScene
func (mgr *SpecialMgr) OnFixScene(rng IRng, ld2 *LevelData2, scene *Scene) error {
	if ld2 == nil {
		return nil
	}

	for _, v := range ld2.SpecialType {
		sp, isok := mgr.MapSpecial[v.SpecialID]
		if isok {
			err := sp.OnFixScene(rng, v, scene)
			if err != nil {
				goutils.Error("SpecialMgr.OnFixScene:OnFixScene",
					zap.Int("SpecialID", v.SpecialID),
					goutils.JSON("SpecialTypeData", v),
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
				goutils.Error("SpecialMgr.GenSymbolLayers:OnGenSymbolLayers",
					zap.Int("SpecialID", v.SpecialID),
					goutils.JSON("SpecialTypeData", v),
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

// Gen2 - GenSymbolLayers
func (mgr *SpecialMgr) Gen2(scene *Scene, x, y, z int, specialLayer int, specialBlock int) error {
	if specialLayer > 0 {
		sp, isok := mgr.MapSpecial[specialLayer]
		if isok {
			spl, err := sp.OnGen2(scene, x, y, z)
			if err != nil {
				goutils.Error("SpecialMgr.Gen2:OnGen2",
					zap.Int("specialLayer", specialLayer),
					zap.Error(err))

				return err
			}

			if spl != nil {
				scene.AddSpecialLayers(spl)
				// scene.SpecialLayers = append(scene.SpecialLayers, spl)
			}
		}
	}

	if specialBlock > 0 {
		sp, isok := mgr.MapSpecial[specialBlock]
		if isok {
			spl, err := sp.OnGen2(scene, x, y, z)
			if err != nil {
				goutils.Error("SpecialMgr.Gen2:OnGen2",
					zap.Int("specialBlock", specialBlock),
					zap.Error(err))

				return err
			}

			if spl != nil {
				scene.AddSpecialLayers(spl)
				// scene.SpecialLayers = append(scene.SpecialLayers, spl)
			}
		}
	}

	return nil
}

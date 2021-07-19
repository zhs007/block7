package block7game

import (
	"fmt"

	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

type AIResult2 struct {
	MaxBlockNums int
	Blocks       []*BlockData
	MapSymbols   map[string]*AIBlockResult
}

func NewAIResult2(scene *Scene, mapBI *BlockInfoMap) *AIResult2 {
	air := &AIResult2{
		MaxBlockNums: scene.MaxBlockNums,
		MapSymbols:   make(map[string]*AIBlockResult),
	}

	air.Blocks = append(air.Blocks, scene.Block...)

	return air
}

func (aiResult *AIResult2) genKey(level int, symbol int) string {
	return fmt.Sprintf("%v-%v", level, symbol)
}

func (aiResult *AIResult2) HasSymbol(level int, symbol int) bool {
	k := aiResult.genKey(level, symbol)
	aibr, isok := aiResult.MapSymbols[k]
	if !isok {
		return false
	}

	return aibr.State != 0 || len(aibr.Arr) > 0
}

func (aiResult *AIResult2) StartSymbol(level int, symbol int) {
	k := aiResult.genKey(level, symbol)
	_, isok := aiResult.MapSymbols[k]
	if !isok {
		aibr := &AIBlockResult{
			Symbol:     symbol,
			LastBlocks: append([]*BlockData{}, aiResult.Blocks...),
		}

		aiResult.MapSymbols[k] = aibr
	}
}

func (aiResult *AIResult2) StopSymbol(level int, symbol int, state int) {
	k := aiResult.genKey(level, symbol)
	aibr, isok := aiResult.MapSymbols[k]
	if isok {
		aibr.State = state
	}
}

func (aiResult *AIResult2) Click(level int, symbol int, scene *Scene, bd *BlockData) bool {
	k := aiResult.genKey(level, symbol)
	aibr, isok := aiResult.MapSymbols[k]
	if isok {
		if !HasBlockData(aibr.Arr, bd.X, bd.Y, bd.Z) {
			if scene.HasBlock(bd.X, bd.Y, bd.Z) {
				if scene.CanClickEx(bd.X, bd.Y, bd.Z, aibr.Arr) {
					aibr.Arr = append(aibr.Arr, bd)
					aibr.LastBlocks = append(aibr.LastBlocks, bd)
					cn := CountBlockData(aibr.LastBlocks, bd.Symbol)
					if cn >= BlockNums {
						aibr.LastBlocks = RemoveBlockData(aibr.LastBlocks, bd.Symbol, BlockNums*cn/BlockNums)
					}

					if len(aibr.LastBlocks) >= aiResult.MaxBlockNums {
						return false
					}
				} else {
					block7utils.Warn("AIResult2.Click:CanClickEx",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z))
				}
			} else {
				block7utils.Warn("AIResult2.Click:HasBlock",
					zap.Int("x", bd.X),
					zap.Int("y", bd.Y),
					zap.Int("z", bd.Z))
			}
		}
	}

	return true
}

func (aiResult *AIResult2) ClickEx(level int, symbol int, scene *Scene, bd *BlockData) bool {
	if len(bd.Parent) > 0 {
		for _, v := range bd.Parent {
			ret := aiResult.ClickEx(level, symbol, scene, v)
			if !ret {
				return false
			}
		}
	}

	return aiResult.Click(level, symbol, scene, bd)
}

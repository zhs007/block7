package block7game

import (
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type AIBlockResult struct {
	Symbol     int
	Arr        []*BlockData
	LastBlocks []*BlockData
	State      int
}

type AIResult struct {
	MaxBlockNums int
	Blocks       []*BlockData
	MapSymbols   map[int]*AIBlockResult
}

func NewAIResult(scene *Scene, mapBI *BlockInfoMap) *AIResult {
	air := &AIResult{
		MaxBlockNums: scene.MaxBlockNums,
		MapSymbols:   make(map[int]*AIBlockResult),
	}

	air.Blocks = append(air.Blocks, scene.Block...)

	return air
}

func (aiResult *AIResult) HasSymbol(symbol int) bool {
	aibr, isok := aiResult.MapSymbols[symbol]
	if !isok {
		return false
	}

	return aibr.State != 0 || len(aibr.Arr) > 0
}

func (aiResult *AIResult) StartSymbol(symbol int) {
	_, isok := aiResult.MapSymbols[symbol]
	if !isok {
		aibr := &AIBlockResult{
			Symbol:     symbol,
			LastBlocks: append([]*BlockData{}, aiResult.Blocks...),
		}

		aiResult.MapSymbols[symbol] = aibr
	}
}

func (aiResult *AIResult) StopSymbol(symbol int, state int) {
	aibr, isok := aiResult.MapSymbols[symbol]
	if isok {
		aibr.State = state
	}
}

func (aiResult *AIResult) Click(symbol int, scene *Scene, bd *BlockData) bool {
	aibr, isok := aiResult.MapSymbols[symbol]
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
					goutils.Warn("AIResult.Click:CanClickEx",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z))
				}
			} else {
				goutils.Warn("AIResult.Click:HasBlock",
					zap.Int("x", bd.X),
					zap.Int("y", bd.Y),
					zap.Int("z", bd.Z))
			}
		}
	}

	return true
}

func (aiResult *AIResult) ClickEx(symbol int, scene *Scene, bd *BlockData) bool {
	if len(bd.Parent) > 0 {
		for _, v := range bd.Parent {
			ret := aiResult.ClickEx(symbol, scene, v)
			if !ret {
				return false
			}
		}
	}

	return aiResult.Click(symbol, scene, bd)
}

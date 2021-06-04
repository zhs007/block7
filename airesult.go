package block7

type AIBlockResult struct {
	Symbol     int
	Arr        []*BlockData
	LastBlocks []*BlockData
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

	return len(aibr.Arr) > 0
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

func (aiResult *AIResult) Click(symbol int, scene *Scene, bd *BlockData) bool {
	aibr, isok := aiResult.MapSymbols[symbol]
	if isok {
		if scene.HasBlock(bd.X, bd.Y, bd.Z) {
			aibr.Arr = append(aibr.Arr, bd)
			aibr.LastBlocks = append(aibr.LastBlocks, bd)
			cn := CountBlockData(aibr.LastBlocks, bd.Symbol)
			if cn >= BlockNums {
				aibr.LastBlocks = RemoveBlockData(aibr.LastBlocks, bd.Symbol, BlockNums*cn/BlockNums)
			}

			if len(aibr.LastBlocks) >= aiResult.MaxBlockNums {
				return false
			}
		}
	}

	return true
}

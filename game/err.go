package block7game

import "errors"

var (
	// ErrInvalidSpecialType - invalid SpecialType
	ErrInvalidSpecialType = errors.New("invalid SpecialType")
	// ErrInvalidMinMaxType - invalid minType or maxType
	ErrInvalidMinMaxType = errors.New("invalid minType or maxType")
	// ErrInvalidBombNums - invalid bomb nums
	ErrInvalidBombNums = errors.New("invalid bomb nums")
	// ErrInvalidParams - invalid params
	ErrInvalidParams = errors.New("invalid params")
	// ErrInvalidLevel - invalid level
	ErrInvalidLevel = errors.New("invalid level")
	// ErrInvalidSpecialNums - invalid special nums
	ErrInvalidSpecialNums = errors.New("invalid special nums")
	// ErrInvalidSymbolsLength - invalid symbols length
	ErrInvalidSymbolsLength = errors.New("invalid symbols length")
	// ErrInvalidGenBrotherBlocksNums - invalid GenBrotherBlocks nums
	ErrInvalidGenBrotherBlocksNums = errors.New("invalid GenBrotherBlocks nums")
	// ErrInvalidSceneWHL - invalid scene width or height or layers
	ErrInvalidSceneWHL = errors.New("invalid scene width or height or layers")
	// ErrInvalidHistoryWidth - invalid history width
	ErrInvalidHistoryWidth = errors.New("invalid history width")
	// ErrInvalidPositionList - invalid position list
	ErrInvalidPositionList = errors.New("invalid position list")
	// ErrInvalidSafeList - invalid safe list
	ErrInvalidSafeList = errors.New("invalid safe list")
	// ErrInvalidCakeNums - invalid cake nums
	ErrInvalidCakeNums = errors.New("invalid cake nums")
	// ErrInvalidRainbowNums - invalid rainbow nums
	ErrInvalidRainbowNums = errors.New("invalid rainbow nums")
	// ErrInvalidTeleportNums - invalid teleport nums
	ErrInvalidTeleportNums = errors.New("invalid teleport nums")
	// ErrInvalidMap2BlockNums - invalid map2 block nums
	ErrInvalidMap2BlockNums = errors.New("invalid map2 block nums")
	// ErrInvalidIconTypes2 - invalid IconTypes2
	ErrInvalidIconTypes2 = errors.New("invalid IconTypes2")
	// ErrInvalidMapExcelFile - invalid MapExcelFile
	ErrInvalidMapExcelFile = errors.New("invalid MapExcelFile")
	// ErrInvalidMapExcelWidthHeight - invalid MapExcel Width Height
	ErrInvalidMapExcelWidthHeight = errors.New("invalid MapExcel Width Height")
	// ErrInvalidBlockNumber - invalid block number
	ErrInvalidBlockNumber = errors.New("invalid block number")
	// ErrRecoveBlock - invalid recove block
	ErrRecoveBlock = errors.New("invalid recove block")
)

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
)

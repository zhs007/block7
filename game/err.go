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
)

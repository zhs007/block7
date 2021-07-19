package block7

import "errors"

var (
	// ErrInvalidSymbolsLength - invalid symbols length
	ErrInvalidSymbolsLength = errors.New("invalid symbols length")

	// ErrInvalidLevel - invalid level
	ErrInvalidLevel = errors.New("invalid level")

	// ErrInvalidParams - invalid params
	ErrInvalidParams = errors.New("invalid params")

	// ErrInvalidSpecialNums - invalid special nums
	ErrInvalidSpecialNums = errors.New("invalid special nums")
)

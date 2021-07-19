package block7game

import "errors"

var (
	// ErrInvalidSpecialType - invalid SpecialType
	ErrInvalidSpecialType = errors.New("invalid SpecialType")
	// ErrInvalidMinMaxType - invalid minType or maxType
	ErrInvalidMinMaxType = errors.New("invalid minType or maxType")
	// ErrInvalidBombNums - invalid bomb nums
	ErrInvalidBombNums = errors.New("invalid bomb nums")
)

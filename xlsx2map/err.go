package main

import "errors"

var (
	// ErrInvalidMapExcelFile - invalid MapExcelFile
	ErrInvalidMapExcelFile = errors.New("invalid MapExcelFile")
	// ErrInvalidMapExcelWidthHeight - invalid MapExcel Width Height
	ErrInvalidMapExcelWidthHeight = errors.New("invalid MapExcel Width Height")
)

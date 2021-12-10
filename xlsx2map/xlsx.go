package main

import (
	"github.com/xuri/excelize/v2"
	block7game "github.com/zhs007/block7/game"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

func loadExcel(fn string) (*block7game.Stage, error) {
	f, err := excelize.OpenFile(fn)
	if err != nil {
		goutils.Error("loadExcel:OpenFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	lstname := f.GetSheetList()
	if len(lstname) <= 0 {
		goutils.Error("loadExcel:GetSheetList",
			goutils.JSON("SheetList", lstname),
			zap.String("fn", fn),
			zap.Error(ErrInvalidMapExcelFile))

		return nil, ErrInvalidMapExcelFile
	}

	xoff, yoff := block7game.OffsetStringToXYOff("1,1,1")

	stage := &block7game.Stage{
		MapType: 1,
		Offset:  "1,1,1",
		XOff:    xoff,
		YOff:    yoff,
	}

	for _, sheet := range lstname {
		rows, err := f.GetRows(sheet)
		if err != nil {
			goutils.Error("loadExcel:GetRows",
				zap.String("fn", fn),
				zap.Error(err))

			return nil, err
		}

		h := len(rows)
		w := len(rows[0])

		if stage.Width == 0 {
			stage.Width = w
		}

		if stage.Height == 0 {
			stage.Height = h
		}

		if stage.Width != w || stage.Height != h {
			goutils.Error("loadExcel:checkWeightHeight",
				zap.String("fn", fn),
				zap.String("sheet", sheet),
				zap.Error(ErrInvalidMapExcelWidthHeight))

			return nil, ErrInvalidMapExcelWidthHeight
		}

		arr2 := [][]int{}
		for y, yarr := range rows {

			arr1 := []int{}
			for x, v := range yarr {
				iv, err := goutils.String2Int64(v)
				if err != nil {
					goutils.Error("loadExcel:checkWeightHeight",
						zap.String("fn", fn),
						zap.String("sheet", sheet),
						zap.Int("x", x),
						zap.Int("y", y),
						zap.String("val", v),
						zap.Error(err))

					return nil, ErrInvalidMapExcelWidthHeight
				}

				arr1 = append(arr1, int(iv))
			}

			arr2 = append(arr2, arr1)
		}

		stage.Layer = append(stage.Layer, arr2)
	}

	stage.IconNums = stage.CountSymbols()

	return stage, nil
}

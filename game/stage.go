package block7game

import (
	"fmt"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	"github.com/xuri/excelize/v2"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// Stage - stage
type Stage struct {
	Width      int       `json:"width"`
	Height     int       `json:"height"`
	Offset     string    `json:"offset"`
	Layer      [][][]int `json:"layer"`
	IconNums   int       `json:"iconnums"`
	XOff       int       `json:"xoff"`
	YOff       int       `json:"yoff"`
	MapType    int       `json:"mapTypes"`   // 地图类型，0是老版本方式，1是新版本
	LayerLevel []int     `json:"layerlevel"` // 分章节，[1,1,1,0,0]，就是5层分2个章节，最上面2层全部消除完才允许操作下面3层
	// ComboEnable bool      `json:"comboEnable"`
}

func LoadStage(fn string) (*Stage, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	stage := &Stage{}
	err = json.Unmarshal(data, stage)
	if err != nil {
		return nil, err
	}

	xoff, yoff := OffsetStringToXYOff(stage.Offset)
	stage.XOff = xoff
	stage.YOff = yoff

	// if len(stage.Offset) > 0 {
	// 	arr := strings.Split(stage.Offset, ",")
	// 	if len(arr) == 3 {
	// 		if arr[0] == "0" {
	// 			stage.XOff = 1
	// 			stage.YOff = -1
	// 		} else {
	// 			stage.XOff = -1
	// 			stage.YOff = 1
	// 		}
	// 	}
	// }

	stage.IconNums = stage.CountSymbols()

	return stage, nil
}

func (stage *Stage) CountSymbols() int {
	n := 0
	for _, arrlayer := range stage.Layer {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v > 0 {
					n++
				}
			}
		}
	}

	return n
}

func (stage *Stage) Analyze(str string) error {
	layerNums := make(map[int]int)
	layerLayer := make(map[int]int)
	layerBlock := make(map[int]int)

	for _, arrlayer := range stage.Layer {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v > 0 {
					area := getBlockArea(v)
					spl := getBlockSpecialLayer(v)
					spb := getBlockSpecialBlock(v)

					layerNums[area]++
					if spl > 0 {
						layerLayer[area]++
					}

					if spb == 0 {
						layerBlock[area]++
					}
				}
			}
		}
	}

	for k := range layerNums {
		goutils.Info("Stage.Analyze",
			zap.String("msg", str),
			zap.Int("nums", layerNums[k]),
			zap.Int("layernums", layerLayer[k]),
			zap.Int("blocknums", layerBlock[k]),
		)
	}

	return nil
}

func (stage *Stage) Analyze2(name string) (*MapState, error) {
	ms := &MapState{
		Name:             name,
		AreaBlockNums:    make(map[int]int),
		AreaSPLayerNums:  make(map[int]int),
		SPMap:            make(map[int]int),
		Layers:           len(stage.Layer),
		Height:           len(stage.Layer[0]),
		Width:            len(stage.Layer[0][0]),
		LayerBlockNums:   make([]int, len(stage.Layer)),
		LayerSPLayerNums: make([]int, len(stage.Layer)),
	}

	for layeri, arrlayer := range stage.Layer {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v > 0 {
					area := getBlockArea(v)
					spl := getBlockSpecialLayer(v)
					spb := getBlockSpecialBlock(v)

					if spl > 0 {
						_, isok := ms.AreaSPLayerNums[area]
						if isok {
							ms.AreaSPLayerNums[area]++
						} else {
							ms.AreaSPLayerNums[area] = 1
						}

						ms.LayerSPLayerNums[layeri]++
					}

					if spb == 0 {
						_, isok := ms.AreaBlockNums[area]
						if isok {
							ms.AreaBlockNums[area]++
						} else {
							ms.AreaBlockNums[area] = 1
						}

						ms.LayerBlockNums[layeri]++
					}

					if spl > 0 {
						_, isok := ms.SPMap[spl]
						if isok {
							ms.SPMap[spl]++
						} else {
							ms.SPMap[spl] = 1
						}
					}

					if spb > 0 {
						_, isok := ms.SPMap[spb]
						if isok {
							ms.SPMap[spb]++
						} else {
							ms.SPMap[spb] = 1
						}
					}
				}
			}
		}
	}

	ms.SpecialType = ""
	for k, v := range ms.SPMap {
		if ms.SpecialType == "" {
			ms.SpecialType += fmt.Sprintf("%03d,%d", k, v)
		} else {
			ms.SpecialType += fmt.Sprintf(",%03d,%d", k, v)
		}
	}

	areaBlockNums := make([]int, len(ms.AreaBlockNums))
	for k, v := range ms.AreaBlockNums {
		areaBlockNums[k-1] = v
	}

	si := 31001
	for _, v := range areaBlockNums {
		if v%3 != 0 {
			return nil, ErrInvalidBlockNumber
		}

		strblock := ""
		for i := 0; i < v/3; i++ {
			if strblock == "" {
				strblock += fmt.Sprintf("%v", si+i)
			} else {
				strblock += fmt.Sprintf(",%v", si+i)
			}
		}
		si += v / 3

		ms.IconType2 = append(ms.IconType2, strblock)
	}

	return ms, nil
}

func LoadExcel(fn string) (*Stage, error) {
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

	xoff, yoff := OffsetStringToXYOff("1,1,1")

	stage := &Stage{
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

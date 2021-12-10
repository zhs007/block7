package block7game

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

// Stage - stage
type Stage struct {
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	Offset      string    `json:"offset"`
	Layer       [][][]int `json:"layer"`
	IconNums    int       `json:"iconnums"`
	XOff        int       `json:"xoff"`
	YOff        int       `json:"yoff"`
	MapType     int       `json:"mapTypes"` // 地图类型，0是老版本方式，1是新版本
	ComboEnable bool      `json:"comboEnable"`
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

package block7game

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type MapState struct {
	Name             string      `json:"name"`
	Width            int         `json:"width"`
	Height           int         `json:"height"`
	Layers           int         `json:"layers"`
	SpecialType      string      `json:"specialType"`
	IconType2        []string    `json:"iconType2"`
	AreaBlockNums    map[int]int `json:"areaBlockNums"`
	AreaSPLayerNums  map[int]int `json:"areaSPLayerNums"`
	LayerBlockNums   []int       `json:"layerBlockNums"`
	LayerSPLayerNums []int       `json:"layerSPLayerNums"`
	SPMap            map[int]int `json:"SPMap"`
}

func SaveMapStateList(fn string, lst []*MapState) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	fd, err := json.MarshalIndent(lst, "", "  ")
	if err != nil {
		goutils.Error("Marshal",
			goutils.JSON("lst", lst),
			zap.Error(err))

		return err
	}

	ioutil.WriteFile(fn, fd, 0644)

	return nil
}

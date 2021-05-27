package block7

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

// Stage - stage
type Stage struct {
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Offset   string    `json:"offset"`
	Layer    [][][]int `json:"layer"`
	IconNums int       `json:"iconnums"`
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

	return stage, nil
}

package block7game

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

type LevelData struct {
	ID          string `json:"id"`
	MapID       string `json:"map"`
	MinType     string `json:"minType"`
	MaxType     string `json:"maxType"`
	SpecialType string `json:"specialType"`
}

type SpecialTypeData struct {
	SpecialID int `json:"special"`
	Nums      int `json:"nums"`
}

type LevelData2 struct {
	ID          int               `json:"id"`
	MapID       int               `json:"map"`
	MinType     int               `json:"minType"`
	MaxType     int               `json:"maxType"`
	SpecialType []SpecialTypeData `json:"specialType"`
}

type LevelMgr struct {
	MapLevel map[int]*LevelData2
}

// NewLevelMgr - new a LevelMgr
func NewLevelMgr() *LevelMgr {
	return &LevelMgr{
		MapLevel: make(map[int]*LevelData2),
	}
}

// LoadLevel - load level file
func (mgr *LevelMgr) LoadLevel(fn string) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	arr := []LevelData{}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	return nil
}

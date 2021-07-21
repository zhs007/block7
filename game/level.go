package block7game

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

type LevelData struct {
	ID          string `json:"id"`
	MapID       string `json:"map"`
	MinType     string `json:"minType"`
	MaxType     string `json:"maxType"`
	SpecialType string `json:"specialType"`
}

type SpecialTypeData struct {
	SpecialID int         `json:"special"`
	Nums      int         `json:"nums"`
	Data      interface{} `json:"-"`
}

type LevelData2 struct {
	ID          int                `json:"id"`
	MapID       int                `json:"map"`
	MinType     int                `json:"minType"`
	MaxType     int                `json:"maxType"`
	SpecialType []*SpecialTypeData `json:"specialType"`
}

func (ld2 *LevelData2) GenSymbols() []int {
	if ld2.MinType == ld2.MaxType {
		return GenSymbols(ld2.MinType)
	}

	ci := rand.Int()
	n := ld2.MinType + ci%(ld2.MaxType-ld2.MinType)

	return GenSymbols(n)
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
		block7utils.Error("LevelMgr.LoadLevel:ReadFile",
			zap.String("fn", fn),
			zap.Error(err))

		return err
	}

	arr := []LevelData{}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	for i, v := range arr {
		ld2 := &LevelData2{}

		id, err := strconv.Atoi(v.ID)
		if err != nil {
			block7utils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("id", v.ID),
				zap.Error(err))

			return err
		}
		ld2.ID = id

		mapid, err := strconv.Atoi(v.MapID)
		if err != nil {
			block7utils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("map", v.MapID),
				zap.Error(err))

			return err
		}
		ld2.MapID = mapid

		mintype, err := strconv.Atoi(v.MinType)
		if err != nil {
			block7utils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("mintype", v.MinType),
				zap.Error(err))

			return err
		}
		ld2.MinType = mintype

		maxtype, err := strconv.Atoi(v.MaxType)
		if err != nil {
			block7utils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("maxtype", v.MaxType),
				zap.Error(err))

			return err
		}
		ld2.MaxType = maxtype

		if ld2.MinType > ld2.MaxType {
			block7utils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("mintype", v.MinType),
				zap.String("maxtype", v.MaxType),
				zap.Error(ErrInvalidMinMaxType))

			return ErrInvalidMinMaxType
		}

		arrstr := strings.Split(v.SpecialType, ",")
		if len(arrstr) > 1 {
			if len(arrstr)%2 == 1 {
				block7utils.Error("LevelMgr.LoadLevel:Split",
					zap.String("fn", fn),
					zap.Int("i", i),
					zap.String("specialType", v.SpecialType),
					zap.Error(ErrInvalidSpecialType))

				return ErrInvalidSpecialType
			}

			for j := 0; j < len(arrstr)/2; j++ {
				std := &SpecialTypeData{}

				sid, err := strconv.Atoi(arrstr[j*2])
				if err != nil {
					block7utils.Error("LevelMgr.LoadLevel:Atoi",
						zap.String("fn", fn),
						zap.Int("i", i),
						zap.Int("j", j*2),
						zap.String("val", arrstr[j*2]),
						zap.Error(err))

					return err
				}
				std.SpecialID = sid

				nums, err := strconv.Atoi(arrstr[j*2+1])
				if err != nil {
					block7utils.Error("LevelMgr.LoadLevel:Atoi",
						zap.String("fn", fn),
						zap.Int("i", i),
						zap.Int("j", j*2+1),
						zap.String("val", arrstr[j*2+1]),
						zap.Error(err))

					return err
				}
				std.Nums = nums

				ld2.SpecialType = append(ld2.SpecialType, std)
			}
		}

		mgr.MapLevel[id] = ld2
	}

	return nil
}

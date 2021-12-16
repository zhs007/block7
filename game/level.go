package block7game

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type LevelData struct {
	ID          string   `json:"id"`
	MapID       string   `json:"map"`
	MinType     string   `json:"minType"`
	MaxType     string   `json:"maxType"`
	SpecialType string   `json:"specialType"`
	IconType2   []string `json:"iconType2"`
}

type SpecialTypeData struct {
	SpecialID int         `json:"special"`
	Nums      int         `json:"nums"`
	Data      interface{} `json:"-"`
}

type LevelData2 struct {
	ID             int                `json:"id"`
	MapID          int                `json:"map"`
	MinType        int                `json:"minType"`
	MaxType        int                `json:"maxType"`
	SpecialType    []*SpecialTypeData `json:"specialType"`
	IconType2      [][]int            `json:"iconType2"`   // 这个是直接解码的level表里的数据
	IconType2Ex    [][]int            `json:"iconType2ex"` // 这个是处理后的数据，类似 [[1,2,3],[4,5,6]] 这样
	SpecialTypeStr string             `json:"specialTypeString"`
}

func (ld2 *LevelData2) GenSymbols() []int {
	if ld2.MinType == ld2.MaxType {
		return GenSymbols(ld2.MinType)
	}

	ci := rand.Int()
	n := ld2.MinType + ci%(ld2.MaxType-ld2.MinType)

	return GenSymbols(n)
}

func (ld2 *LevelData2) genIconType2Ex() {
	ld2.IconType2Ex = nil
	icons := []int{}

	for _, arr := range ld2.IconType2 {
		arr2 := []int{}

		for _, icon := range arr {
			ci := goutils.FindInt(icons, icon)
			if ci < 0 {
				icons = append(icons, icon)
				arr2 = append(arr2, len(icons))
			} else {
				arr2 = append(arr2, ci+1)
			}
		}

		ld2.IconType2Ex = append(ld2.IconType2Ex, arr2)
	}
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
		goutils.Error("LevelMgr.LoadLevel:ReadFile",
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
		ld2 := &LevelData2{
			SpecialTypeStr: v.SpecialType,
		}

		id, err := strconv.Atoi(v.ID)
		if err != nil {
			goutils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("id", v.ID),
				zap.Error(err))

			return err
		}
		ld2.ID = id

		mapid, err := strconv.Atoi(v.MapID)
		if err != nil {
			goutils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("map", v.MapID),
				zap.Error(err))

			return err
		}
		ld2.MapID = mapid

		mintype, err := strconv.Atoi(v.MinType)
		if err != nil {
			goutils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("mintype", v.MinType),
				zap.Error(err))

			return err
		}
		ld2.MinType = mintype

		maxtype, err := strconv.Atoi(v.MaxType)
		if err != nil {
			goutils.Error("LevelMgr.LoadLevel:Atoi",
				zap.String("fn", fn),
				zap.Int("i", i),
				zap.String("maxtype", v.MaxType),
				zap.Error(err))

			return err
		}
		ld2.MaxType = maxtype

		if ld2.MinType > ld2.MaxType {
			goutils.Error("LevelMgr.LoadLevel:Atoi",
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
				goutils.Error("LevelMgr.LoadLevel:Split",
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
					goutils.Error("LevelMgr.LoadLevel:Atoi",
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
					goutils.Error("LevelMgr.LoadLevel:Atoi",
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

		if len(v.IconType2) > 0 {
			for _, iv := range v.IconType2 {
				arrstr := strings.Split(iv, ",")
				if len(arrstr) > 0 {
					arri := []int{}

					for _, strv := range arrstr {
						i64, err := goutils.String2Int64(strv)
						if err != nil {
							goutils.Error("LevelMgr.LoadLevel:IconType2",
								zap.String("fn", fn),
								zap.String("val", strv),
								zap.Error(err))

							return err
						}

						arri = append(arri, int(i64))
					}

					ld2.IconType2 = append(ld2.IconType2, arri)
				}
			}

			ld2.genIconType2Ex()
		}

		mgr.MapLevel[id] = ld2
	}

	return nil
}

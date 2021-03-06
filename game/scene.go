package block7game

import (
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhs007/block7/block7pb"
	goutils "github.com/zhs007/goutils"
	"go.uber.org/zap"
)

// Scene - scene
type Scene struct {
	StageID           int             `json:"stageid"` // 对应missionid，就是关卡id，版本不同，可能没有对比价值
	MapID             int             `json:"mapid"`   // 实际的mapid，有对比价值
	Version           int             `json:"version"`
	SceneID           int64           `json:"sceneid"` // 关卡的动态id，同一个地图，可能随机出不同的scene，这就是随机后的id
	UserID            int64           `json:"userid"`
	Width             int             `json:"width"`
	Height            int             `json:"height"`
	Layers            int             `json:"layers"`
	XOff              int             `json:"xoff"`
	YOff              int             `json:"yoff"`
	Arr               [][][]int       `json:"-"`
	Block             []*BlockData    `json:"-"`
	BlockEx           []*BlockData    `json:"-"`
	MaxBlockNums      int             `json:"-"`
	InitArr           [][][]int       `json:"layer"`
	History           [][]int         `json:"history"`
	ClickValues       int             `json:"clickValues"`
	FinishedPer       float32         `json:"finishedPer"`
	Offset            string          `json:"offset"`
	IsOutputScene     bool            `json:"isOutputScene"`
	SpecialLayers     []*SpecialLayer `json:"specialLayers"`     // 这个是自己用的
	BlockNums         int             `json:"-"`                 // 初始化block数量
	RngData           []int64         `json:"rngdata"`           // 前端rng数据
	GameState         int32           `json:"gamestate"`         // 前端gamestate
	ClientMissionID   int             `json:"clientMissionID"`   // 前端missionID
	ClientStageType   int             `json:"clientStageType"`   // 前端stage type
	FirstItem         int             `json:"firstItem"`         // 前置道具
	IsFullHistoryData bool            `json:"isFullHistoryData"` // 是否是完整的数据
	ClientVersion     string          `json:"clientVersion"`
	LastHP            int             `json:"lastHP"`
	LastCoin          int             `json:"lastCoin"`
	RefreshTimes      int             `json:"refresh"`
	BackTimes         int             `json:"back"`
	BombTimes         int             `json:"bomb"`
	RebirthTimes      int             `json:"rebirth"`
	MapType           int             `json:"mapTypes"`     // 地图类型，0是老版本方式，1是新版本
	SpecialType       string          `json:"specialType"`  // level.json 文件内容
	LayerLevel        []int           `json:"layerlevel"`   // 分章节，[1,1,1,0,0]，就是5层分2个章节，最上面2层全部消除完才允许操作下面3层
	InitLayerArr      [][][]int       `json:"initlayerarr"` // 前端用数据，区域数组，0,1,2,3,4,5这样
}

// LoadScene - load a scene
func LoadScene(rng IRng, fn string, blockNums int) (*Scene, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	scene := &Scene{}
	err = json.Unmarshal(data, scene)
	if err != nil {
		return nil, err
	}

	scene.MaxBlockNums = blockNums

	scene.Restart()

	return scene, nil
}

// NewScene - new a scene
func NewScene(rng IRng, stage *Stage, symbols []int, blockNums int, ld2 *LevelData2) (*Scene, error) {
	ss, err := MgrSpecial.GenSymbols(ld2)
	if err != nil {
		goutils.Warn("NewScene:MgrSpecial.GenSymbols",
			zap.Error(err))

		return nil, err
	}

	if len(ss) > stage.IconNums {
		goutils.Warn("NewScene:IconNums",
			zap.Error(ErrInvalidSpecialNums))

		return nil, ErrInvalidSpecialNums
	}

	if len(ss) < stage.IconNums {
		ss1, err := genSymbols(rng, symbols, stage.IconNums-len(ss))
		if err != nil {
			goutils.Warn("NewScene:genSymbols",
				zap.Error(err))

			return nil, err
		}

		ss = append(ss, ss1...)
	}

	// goutils.Debug("NewScene",
	// 	goutils.JSON("symbols", ss))

	scene := &Scene{
		Width:        stage.Width,
		Height:       stage.Height,
		Layers:       len(stage.Layer),
		XOff:         stage.XOff,
		YOff:         stage.YOff,
		MaxBlockNums: blockNums,
		Offset:       stage.Offset,
		LayerLevel:   stage.LayerLevel,
		// SpecialType:  ld2.SpecialTypeStr,
	}

	if ld2 != nil {
		scene.SpecialType = ld2.SpecialTypeStr
	}

	nums := 0
	for _, arrlayer := range stage.Layer {
		arrslayer := [][]int{}
		for _, arrrow := range arrlayer {
			arrsrow := []int{}
			for _, v := range arrrow {
				if v == 0 {
					arrsrow = append(arrsrow, 0)
				} else {
					nss, cs, err := randSymbols(rng, ss)
					if err != nil {
						return nil, err
					}

					arrsrow = append(arrsrow, cs)
					ss = nss

					nums++
				}
			}

			arrslayer = append(arrslayer, arrsrow)
		}

		scene.Arr = append(scene.Arr, arrslayer)
	}

	scene.InitArr = goutils.CloneArr3(scene.Arr)
	scene.BlockNums = nums

	err = MgrSpecial.OnFixScene(rng, ld2, scene)
	if err != nil {
		goutils.Warn("NewScene:OnFixScene",
			zap.Error(err))

		return nil, err
	}

	err = MgrSpecial.GenSymbolLayers(rng, ld2, scene)
	if err != nil {
		goutils.Warn("NewScene:GenSymbolLayers",
			zap.Error(err))

		return nil, err
	}

	return scene, nil
}

// NewSceneFromData - new a scene
func NewSceneFromData(arr [][][]int, layers []*SpecialLayer) (*Scene, error) {
	scene := &Scene{
		Width:  len(arr[0][0]),
		Height: len(arr[0]),
		Layers: len(arr),
		XOff:   -1,
		YOff:   -1,
		Offset: "1,0,1",
	}

	for _, arrlayer := range arr {
		arrslayer := [][]int{}
		for _, arrrow := range arrlayer {
			arrsrow := []int{}
			for _, v := range arrrow {
				arrsrow = append(arrsrow, int(v))
			}

			arrslayer = append(arrslayer, arrsrow)
		}

		scene.Arr = append(scene.Arr, arrslayer)
	}

	scene.InitArr = goutils.CloneArr3(scene.Arr)
	scene.BlockNums = scene.CountNums(func(x, y, z int) bool {
		return scene.InitArr[z][y][x] > 0
	})

	for _, ssl := range layers {
		spid := SpecialType2SpecialID(ssl.LayerType)
		if spid > 0 {
			cs, isok := MgrSpecial.MapSpecial[spid]
			if isok {
				sl := &SpecialLayer{
					Layer:     int(ssl.Layer),
					LayerType: int(ssl.LayerType),
					Special:   cs,
				}

				sl.Pos = ssl.Pos

				scene.SpecialLayers = append(scene.SpecialLayers, sl)
			}
		}
	}

	return scene, nil
}

// NewSceneFromPB - new a scene
func NewSceneFromPB(pbscene *block7pb.Scene) (*Scene, error) {
	scene := &Scene{
		MapID:           int(pbscene.MapID2),
		Version:         int(pbscene.Version),
		SceneID:         pbscene.SceneID,
		Width:           int(pbscene.Width),
		Height:          int(pbscene.Height),
		Layers:          int(pbscene.Layers),
		XOff:            int(pbscene.XOff),
		YOff:            int(pbscene.YOff),
		Offset:          pbscene.Offset,
		RngData:         pbscene.RngData,
		GameState:       pbscene.GameState,
		BlockNums:       int(pbscene.BlockNums),
		ClientMissionID: int(pbscene.ClientMissionID),
		ClientStageType: int(pbscene.ClientStageType),
		FirstItem:       int(pbscene.FirstItem),
		UserID:          pbscene.UserID,
	}

	// if pbscene.ClientMissionID > 0 {

	// }

	if pbscene.InitArr2 != nil {
		arr, err := goutils.Int32ArrToIntArr3(pbscene.InitArr2, scene.Width, scene.Height, scene.Layers)
		if err != nil {
			goutils.Warn("NewSceneFromPB:Int32ArrToIntArr3:InitArr2",
				zap.Int("w", scene.Width),
				zap.Int("h", scene.Height),
				zap.Int("len", len(pbscene.InitArr2)),
				zap.Error(err))

			return nil, err
		}

		scene.Arr = arr
	} else {
		for _, arrlayer := range pbscene.InitArr {
			arrslayer := [][]int{}
			for _, arrrow := range arrlayer.Values {
				arrsrow := []int{}
				for _, v := range arrrow.Values {
					arrsrow = append(arrsrow, int(v))
				}

				arrslayer = append(arrslayer, arrsrow)
			}

			scene.Arr = append(scene.Arr, arrslayer)
		}
	}

	scene.InitArr = goutils.CloneArr3(scene.Arr)
	scene.BlockNums = scene.CountNums(func(x, y, z int) bool {
		return scene.InitArr[z][y][x] > 0
	})

	if len(pbscene.History2) > 0 {
		arr, err := goutils.Int32ArrToIntArr2(pbscene.History2, 4, len(pbscene.History2)/4)
		if err != nil {
			goutils.Warn("NewSceneFromPB:Int32ArrToIntArr2:History2",
				zap.Error(err))

			return nil, err
		}

		scene.History = arr
	}

	for _, pbsl := range pbscene.SpecialLayer {
		cs, isok := MgrSpecial.MapSpecial[int(pbsl.Special)]
		if isok {
			sl := &SpecialLayer{
				Layer:     int(pbsl.Layer),
				LayerType: int(pbsl.LayerType),
				Special:   cs,
			}

			arr, err := goutils.Int32ArrToIntArr2(pbsl.Values, 3, len(pbsl.Values)/3)
			if err != nil {
				goutils.Warn("NewSceneFromPB:Int32ArrToIntArr2:SpecialLayer",
					zap.Error(err))

				return nil, err
			}
			sl.Pos = arr

			scene.SpecialLayers = append(scene.SpecialLayers, sl)
		}
	}

	return scene, nil
}

func (scene *Scene) Restart() {
	scene.Arr = goutils.CloneArr3(scene.InitArr)
	scene.History = nil
	scene.ClickValues = 0
	scene.Block = nil
}

func (scene *Scene) CountSymbols() int {
	n := 0
	for _, arrlayer := range scene.Arr {
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

func (scene *Scene) CountSymbol(symbol int) int {
	n := 0
	for _, arrlayer := range scene.Arr {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v == symbol {
					n++
				}
			}
		}
	}

	return n
}

func (scene *Scene) IsValid() bool {
	n := 0
	for _, arrlayer := range scene.Arr {
		for _, arrrow := range arrlayer {
			for _, v := range arrrow {
				if v > 0 {
					n++
				}
			}
		}
	}

	return (n+len(scene.Block))%3 == 0
}

func (scene *Scene) HasBlock(x, y, z int) bool {
	if x < 0 || x >= scene.Width {
		return false
	}

	if y < 0 || y >= scene.Height {
		return false
	}

	if z < 0 || z >= len(scene.Arr) {
		return false
	}

	return scene.Arr[z][y][x] > 0
}

func (scene *Scene) CanClick(x, y, z int) bool {
	if x < 0 || x >= scene.Width {
		return false
	}

	if y < 0 || y >= scene.Height {
		return false
	}

	if z < 0 || z >= len(scene.Arr) {
		return false
	}

	if z == len(scene.Arr)-1 {
		return scene.Arr[z][y][x] > 0
	}

	if z%2 == 0 {
		for zi := 1; z+zi < len(scene.Arr); zi++ {
			if (z+zi)%2 == 1 {
				if scene.HasBlock(x, y, z+zi) ||
					scene.HasBlock(x+scene.XOff, y, z+zi) ||
					scene.HasBlock(x, y+scene.YOff, z+zi) ||
					scene.HasBlock(x+scene.XOff, y+scene.YOff, z+zi) {
					return false
				}
			} else {
				if scene.HasBlock(x, y, z+zi) {
					return false
				}
			}
		}
	} else {
		for zi := 1; z+zi < len(scene.Arr); zi++ {
			if (z+zi)%2 == 0 {
				if scene.HasBlock(x, y, z+zi) ||
					scene.HasBlock(x-scene.XOff, y, z+zi) ||
					scene.HasBlock(x, y-scene.YOff, z+zi) ||
					scene.HasBlock(x-scene.XOff, y-scene.YOff, z+zi) {
					return false
				}
			} else {
				if scene.HasBlock(x, y, z+zi) {
					return false
				}
			}
		}
	}

	return scene.Arr[z][y][x] > 0
}

func (scene *Scene) CanClickEx(x, y, z int, lst []*BlockData) bool {
	if x < 0 || x >= scene.Width {
		return false
	}

	if y < 0 || y >= scene.Height {
		return false
	}

	if z < 0 || z >= len(scene.Arr) {
		return false
	}

	if z == len(scene.Arr)-1 {
		return scene.Arr[z][y][x] > 0
	}

	if z%2 == 0 {
		for zi := 1; z+zi < len(scene.Arr); zi++ {
			if (z+zi)%2 == 1 {
				if (scene.HasBlock(x, y, z+zi) && !HasBlockData(lst, x, y, z+zi)) ||
					(scene.HasBlock(x+scene.XOff, y, z+zi) && !HasBlockData(lst, x+scene.XOff, y, z+zi)) ||
					(scene.HasBlock(x, y+scene.YOff, z+zi) && !HasBlockData(lst, x, y+scene.YOff, z+zi)) ||
					(scene.HasBlock(x+scene.XOff, y+scene.YOff, z+zi) && !HasBlockData(lst, x+scene.XOff, y+scene.YOff, z+zi)) {
					return false
				}
			} else {
				if scene.HasBlock(x, y, z+zi) && !HasBlockData(lst, x, y, z+zi) {
					return false
				}
			}
		}
	} else {
		for zi := 1; z+zi < len(scene.Arr); zi++ {
			if (z+zi)%2 == 0 {
				if (scene.HasBlock(x, y, z+zi) && !HasBlockData(lst, x, y, z+zi)) ||
					(scene.HasBlock(x-scene.XOff, y, z+zi) && !HasBlockData(lst, x-scene.XOff, y, z+zi)) ||
					(scene.HasBlock(x, y-scene.YOff, z+zi) && !HasBlockData(lst, x, y-scene.YOff, z+zi)) ||
					(scene.HasBlock(x-scene.XOff, y-scene.YOff, z+zi) && !HasBlockData(lst, x-scene.XOff, y-scene.YOff, z+zi)) {
					return false
				}
			} else {
				if scene.HasBlock(x, y, z+zi) && !HasBlockData(lst, x, y, z+zi) {
					return false
				}
			}
		}
	}

	return scene.Arr[z][y][x] > 0
}

func (scene *Scene) GetLevel1(x, y, z int) []*BlockData {
	if z == 0 {
		return nil
	}

	arr := []*BlockData{}

	return arr
}

func (scene *Scene) GetMaxZ(x, y int) int {
	if len(scene.Arr) == 1 {
		return 0
	}

	for z := len(scene.Arr) - 1; z >= 0; z-- {
		if scene.Arr[z][y][x] > 0 {
			return z
		}
	}

	return 0
}

func (scene *Scene) analysisNeighboring(mapBI *BlockInfoMap, arr []*BlockData, bd *BlockData, level int, depth int) error {
	for xoff := -1; xoff <= 1; xoff++ {
		for yoff := -1; yoff <= 1; yoff++ {
			if xoff == 0 && yoff == 0 {
				continue
			}

			if scene.CanClickEx(bd.X+xoff, bd.Y+yoff, bd.Z, arr) {
				cb, err := mapBI.AddBlockDataEx2(scene, bd.X+xoff, bd.Y+yoff, bd.Z, scene.Arr[bd.Z][bd.Y+yoff][bd.X+xoff], arr)
				if err != nil {
					goutils.Warn("Scene.analysisNeighboring:AddBlockDataEx2",
						zap.Int("x", bd.X+xoff),
						zap.Int("y", bd.Y+yoff),
						zap.Int("z", bd.Z),
						zap.Error(err))

					return err
				}

				if cb != nil {
					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisNeighboring:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (scene *Scene) analysisDepth(mapBI *BlockInfoMap, arr []*BlockData, bd *BlockData, level int, depth int) error {
	// arr := []*BlockData{bd}

	if bd.Z > 0 {
		err := scene.analysisNeighboring(mapBI, arr, bd, level, depth)
		if err != nil {
			goutils.Warn("Scene.analysisDepth:analysisNeighboring",
				zap.Error(err))
		}

		if bd.Z%2 == 0 {
			if scene.CanClickEx(bd.X, bd.Y, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X, bd.Y, bd.Z-1, scene.Arr[bd.Z-1][bd.Y][bd.X], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					// scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X+scene.XOff, bd.Y, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X+scene.XOff, bd.Y, bd.Z-1, scene.Arr[bd.Z-1][bd.Y][bd.X+scene.XOff], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X+scene.XOff),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					// scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X, bd.Y+scene.YOff, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X, bd.Y+scene.YOff, bd.Z-1, scene.Arr[bd.Z-1][bd.Y+scene.YOff][bd.X], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y+scene.YOff),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					// scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X+scene.XOff, bd.Y+scene.YOff, bd.Z-1, arr) {
				cb, err := mapBI.AddBlockDataEx2(scene, bd.X+scene.XOff, bd.Y+scene.YOff, bd.Z-1, scene.Arr[bd.Z-1][bd.Y+scene.YOff][bd.X+scene.XOff], arr)

				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X+scene.XOff),
						zap.Int("y", bd.Y+scene.YOff),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					// scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}
		} else {
			if scene.CanClickEx(bd.X, bd.Y, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X, bd.Y, bd.Z-1, scene.Arr[bd.Z-1][bd.Y][bd.X], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					// scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X-scene.XOff, bd.Y, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X-scene.XOff, bd.Y, bd.Z-1, scene.Arr[bd.Z-1][bd.Y][bd.X-scene.XOff], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X-scene.XOff),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X, bd.Y-scene.YOff, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X, bd.Y-scene.YOff, bd.Z-1, scene.Arr[bd.Z-1][bd.Y-scene.YOff][bd.X], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y-scene.YOff),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}

			if scene.CanClickEx(bd.X-scene.XOff, bd.Y-scene.YOff, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X-scene.XOff, bd.Y-scene.YOff, bd.Z-1, scene.Arr[bd.Z-1][bd.Y-scene.YOff][bd.X-scene.XOff], arr)
				if err != nil {
					goutils.Warn("Scene.analysisDepth:AddBlockDataEx2",
						zap.Int("x", bd.X-scene.XOff),
						zap.Int("y", bd.Y-scene.YOff),
						zap.Int("z", bd.Z-1),
						zap.Error(err))

					return err
				}

				if cb != nil {
					scene.ProcParent(cb, arr)
					// cb.AddParent(bd)
					// bd.AddChild(cb)

					if depth > 0 {
						arr1 := append([]*BlockData{}, arr...)
						arr1 = append(arr1, cb)

						err = scene.analysisDepth(mapBI, arr1, cb, level+1, depth-1)
						if err != nil {
							goutils.Warn("Scene.analysisDepth:analysisDepth",
								zap.Int("x", cb.X),
								zap.Int("y", cb.Y),
								zap.Int("z", cb.Z),
								zap.Error(err))

							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (scene *Scene) Analysis() *BlockInfoMap {
	mapBI := NewBlockInfoMap(DefaultMaxBlockLevel)

	for x := 0; x < scene.Width; x++ {
		for y := 0; y < scene.Height; y++ {
			mz := scene.GetMaxZ(x, y)
			if scene.CanClick(x, y, mz) {
				bd := NewBlockData(x, y, mz, scene.Arr[mz][y][x])
				mapBI.AddBlockData(bd, 0)

				arr := []*BlockData{bd}

				err := scene.analysisDepth(mapBI, arr, bd, 1, 1)
				if err != nil {
					goutils.Warn("Scene.Analysis:analysisDepth",
						zap.Int("x", bd.X),
						zap.Int("y", bd.Y),
						zap.Int("z", bd.Z),
						zap.Error(err))

					return nil
				}
			}
		}
	}

	mapBI.Format()
	for _, v := range scene.Block {
		mapBI.InsBlockSymbol((v.Symbol))
	}

	return mapBI
}

// Click - return gamestate, isok
func (scene *Scene) Click(x, y, z int) (int, bool) {
	if !scene.CanClick(x, y, z) {
		return GameStateRunning, false
	}

	if len(scene.Block) >= scene.MaxBlockNums {
		return GameStateFail, false
	}

	if HasBlockData(scene.Block, x, y, z) {
		return GameStateRunning, false
	}

	scene.ClickValues += len(scene.Block)
	scene.History = append(scene.History, []int{x, y, z, len(scene.Block)})

	b := NewBlockData(x, y, z, scene.Arr[z][y][x])
	scene.Arr[z][y][x] = 0

	scene.Block = insBlockDataAndProc(scene.Block, b)

	if !scene.IsValid() {
		goutils.Warn("Scene.Click:IsValid",
			zap.Int("blocks", len(scene.Block)),
			zap.Int("lastSymbols", scene.CountSymbols()))
	}

	if len(scene.Block) >= scene.MaxBlockNums {
		return GameStateFail, true
	}

	if scene.CountSymbols() == 0 {
		return GameStateSucess, true
	}

	return GameStateRunning, true
}

// CountBlockSymbols - return gamestate, isok
func (scene *Scene) CountBlockSymbols(symbol int) int {
	n := 0
	for _, v := range scene.Block {
		if v.Symbol == symbol {
			n++
		}
	}

	return n
}

func (scene *Scene) Save(fn string) error {
	scene.IsOutputScene = true

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	buf, err := json.Marshal(scene)
	if err != nil {
		goutils.Error("Scene.Save:Marshal",
			zap.Error(err))

		return err
	}

	err = ioutil.WriteFile(fn, buf, os.ModePerm)
	if err != nil {
		goutils.Error("Scene.Save:WriteFile",
			zap.Error(err))

		return err
	}

	return nil
}

func (scene *Scene) IsParent(bd *BlockData, pbd *BlockData) bool {
	if pbd.Z == bd.Z+1 {
		if bd.Z%2 == 0 {
			if (pbd.X == bd.X && pbd.Y == bd.Y) ||
				(pbd.X == bd.X+scene.XOff && pbd.Y == bd.Y) ||
				(pbd.X == bd.X && pbd.Y == bd.Y+scene.YOff) ||
				(pbd.X == bd.X+scene.XOff && pbd.Y == bd.Y+scene.YOff) {
				return true
			}
		} else {
			if (pbd.X == bd.X && pbd.Y == bd.Y) ||
				(pbd.X == bd.X-scene.XOff && pbd.Y == bd.Y) ||
				(pbd.X == bd.X && pbd.Y == bd.Y-scene.YOff) ||
				(pbd.X == bd.X-scene.XOff && pbd.Y == bd.Y-scene.YOff) {
				return true
			}
		}
	}

	return false
}

func (scene *Scene) IsParentEx(bd *BlockData, pbd *BlockData) bool {
	return scene.IsParent(bd, pbd) && !HasBlockData(bd.Parent, pbd.X, pbd.Y, pbd.Z)
}

// 判断是否可能是parent的parent
func (scene *Scene) IsParent2(bd *BlockData, pbd *BlockData, funcHasBlock FuncHasBlock) bool {
	if pbd.Z == bd.Z+1 {
		return scene.IsParent(bd, pbd)
	} else if pbd.Z > bd.Z+1 {
		if bd.Z%2 == 0 {
			if funcHasBlock(bd.X, bd.Y, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X,
					Y: bd.Y,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X+scene.XOff, bd.Y, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X + scene.XOff,
					Y: bd.Y,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X, bd.Y+scene.YOff, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X,
					Y: bd.Y + scene.YOff,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X+scene.XOff, bd.Y+scene.YOff, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X + scene.XOff,
					Y: bd.Y + scene.YOff,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}
		} else {
			if funcHasBlock(bd.X, bd.Y, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X,
					Y: bd.Y,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X-scene.XOff, bd.Y, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X - scene.XOff,
					Y: bd.Y,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X, bd.Y-scene.YOff, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X,
					Y: bd.Y - scene.YOff,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}

			if funcHasBlock(bd.X-scene.XOff, bd.Y-scene.YOff, bd.Z+1) {
				isp2 := scene.IsParent2(&BlockData{
					X: bd.X - scene.XOff,
					Y: bd.Y - scene.YOff,
					Z: bd.Z + 1,
				}, pbd, funcHasBlock)

				if isp2 {
					return true
				}
			}
		}
	}

	return false
}

// 统计子节点数量
func (scene *Scene) GetChildren(lst []int, x, y, z int, funcHasBlock FuncHasBlock) []int {
	if z == 0 {
		return lst
	}

	if goutils.FindInt3(lst, x, y, z) >= 0 {
		return lst
	}

	if z%2 == 0 {
		lst = scene.GetChildren(lst, x, y, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x-scene.XOff, y, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x, y-scene.YOff, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x-scene.XOff, y-scene.YOff, z-1, funcHasBlock)
	} else {
		lst = scene.GetChildren(lst, x, y, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x+scene.XOff, y, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x, y+scene.YOff, z-1, funcHasBlock)
		lst = scene.GetChildren(lst, x+scene.XOff, y+scene.YOff, z-1, funcHasBlock)
	}

	return lst
}

// 统计子节点数量
func (scene *Scene) CountChildrenNums(x, y, z int, funcHasBlock FuncHasBlock) int {
	lst := []int{}

	lst = scene.GetChildren(lst, x, y, z, funcHasBlock)

	return len(lst) / 3
}

// 统计子节点数量
func (scene *Scene) CountChildrenNumsEx(x, y, z int, w, h int, funcHasBlock FuncHasBlock) int {
	lst := []int{}

	for ox := 0; ox < w; ox++ {
		for oy := 0; oy < h; oy++ {
			lst = scene.GetChildren(lst, x+ox, y+oy, z, funcHasBlock)
		}
	}

	return len(lst) / 3
}

// 统计子节点数量
func (scene *Scene) CountNums(funcHasBlock FuncHasBlock) int {
	nums := 0

	for z, arr2 := range scene.InitArr {
		for y, arr1 := range arr2 {
			for x := range arr1 {
				if funcHasBlock(x, y, z) {
					nums++
				}
			}
		}
	}

	return nums
}

func (scene *Scene) ProcParent(bd *BlockData, arr []*BlockData) {
	for _, v := range arr {
		if scene.IsParentEx(bd, v) {
			bd.AddParent(v)
			v.AddChild(bd)
		}
	}
}

func (scene *Scene) ToScenePB() (*block7pb.Scene, error) {
	pbScene := &block7pb.Scene{
		StageID2: int32(scene.StageID),
		MapID2:   int32(scene.MapID),
		Version:  int32(scene.Version),
		SceneID:  scene.SceneID,
		Width:    int32(scene.Width),
		Height:   int32(scene.Height),
		Layers:   int32(scene.Layers),
		XOff:     int32(scene.XOff),
		YOff:     int32(scene.YOff),
		Offset:   scene.Offset,
	}

	arr, x, y, z := goutils.IntArr3ToInt32Arr(scene.InitArr)
	if x != scene.Width || y != scene.Height || z != scene.Layers {
		goutils.Error("Scene.ToScenePB:IntArr3ToInt32Arr",
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Int("z", z),
			zap.Error(ErrInvalidSceneWHL))

		return nil, ErrInvalidSceneWHL
	}

	pbScene.InitArr2 = arr

	// for _, arr2 := range scene.InitArr {
	// 	pblayer := &block7pb.SceneLayer{}

	// 	for _, arr1 := range arr2 {
	// 		pbcolumn := &block7pb.Column{}

	// 		for _, s := range arr1 {
	// 			pbcolumn.Values = append(pbcolumn.Values, int32(s))
	// 		}

	// 		pblayer.Values = append(pblayer.Values, pbcolumn)
	// 	}

	// 	pbScene.InitArr = append(pbScene.InitArr, pblayer)
	// }

	for _, sl := range scene.SpecialLayers {
		pbsl := &block7pb.SpecialLayer{
			Special:   int32(sl.Special.GetSpecialID()),
			Layer:     int32(sl.Layer),
			LayerType: int32(sl.LayerType),
		}

		vals, _, _ := goutils.IntArr2ToInt32Arr(sl.Pos)
		pbsl.Values = vals

		pbScene.SpecialLayer = append(pbScene.SpecialLayer, pbsl)
	}

	return pbScene, nil
}

func (scene *Scene) ToHistoryPB() (*block7pb.Scene, error) {
	pbScene := &block7pb.Scene{
		StageID2:        int32(scene.StageID),
		MapID2:          int32(scene.MapID),
		Version:         int32(scene.Version),
		SceneID:         scene.SceneID,
		UserID:          scene.UserID,
		RngData:         scene.RngData,
		GameState:       scene.GameState,
		BlockNums:       int32(scene.BlockNums),
		ClientMissionID: int32(scene.ClientMissionID),
		ClientStageType: int32(scene.ClientStageType),
		FirstItem:       int32(scene.FirstItem),
		ClientVersion:   scene.ClientVersion,
		LastHP:          int32(scene.LastHP),
		LastCoin:        int32(scene.LastCoin),
		RefreshTimes:    int32(scene.RefreshTimes),
		BackTimes:       int32(scene.BackTimes),
		BombTimes:       int32(scene.BombTimes),
		RebirthTimes:    int32(scene.RebirthTimes),
	}

	if pbScene.StageID2 == 0 && pbScene.ClientMissionID > 0 {
		pbScene.StageID2 = pbScene.ClientMissionID
	}

	if scene.IsFullHistoryData {
		pbScene.Width = int32(scene.Width)
		pbScene.Height = int32(scene.Height)
		pbScene.Layers = int32(scene.Layers)
		pbScene.XOff = int32(scene.XOff)
		pbScene.YOff = int32(scene.YOff)
		pbScene.Offset = scene.Offset

		arr, x, y, z := goutils.IntArr3ToInt32Arr(scene.InitArr)
		if x != scene.Width || y != scene.Height || z != scene.Layers {
			goutils.Error("Scene.ToScenePB:IntArr3ToInt32Arr",
				zap.Int("x", x),
				zap.Int("y", y),
				zap.Int("z", z),
				zap.Error(ErrInvalidSceneWHL))

			return nil, ErrInvalidSceneWHL
		}

		pbScene.InitArr2 = arr

		// for _, arr2 := range scene.InitArr {
		// 	pblayer := &block7pb.SceneLayer{}

		// 	for _, arr1 := range arr2 {
		// 		pbcolumn := &block7pb.Column{}

		// 		for _, s := range arr1 {
		// 			pbcolumn.Values = append(pbcolumn.Values, int32(s))
		// 		}

		// 		pblayer.Values = append(pblayer.Values, pbcolumn)
		// 	}

		// 	pbScene.InitArr = append(pbScene.InitArr, pblayer)
		// }

		for _, sl := range scene.SpecialLayers {
			pbsl := &block7pb.SpecialLayer{
				Special:   int32(sl.Special.GetSpecialID()),
				Layer:     int32(sl.Layer),
				LayerType: int32(sl.LayerType),
			}

			vals, _, _ := goutils.IntArr2ToInt32Arr(sl.Pos)
			pbsl.Values = vals

			pbScene.SpecialLayer = append(pbScene.SpecialLayer, pbsl)
		}
	}

	if len(scene.History) == 0 {
		return pbScene, nil
	}

	arr, x, _ := goutils.IntArr2ToInt32Arr(scene.History)
	if x != 4 {
		goutils.Error("Scene.ToHistoryPB:IntArr2ToInt32Arr",
			zap.Int("x", x),
			zap.Error(ErrInvalidHistoryWidth))

		return nil, ErrInvalidHistoryWidth
	}

	pbScene.History2 = arr

	// for _, arr := range scene.History {
	// 	pbcolumn := &block7pb.Column{}

	// 	for _, s := range arr {
	// 		pbcolumn.Values = append(pbcolumn.Values, int32(s))
	// 	}

	// 	pbScene.History = append(pbScene.History, pbcolumn)
	// }

	return pbScene, nil
}

func (scene *Scene) ReadyToClient() {
	// scene.SpecialLayersData = nil
	// for _, v := range scene.SpecialLayers {
	// 	arr := []int{v.LayerType}

	// 	for _, ca := range v.Pos {
	// 		arr = append(arr, ca...)
	// 	}

	// 	scene.SpecialLayersData = append(scene.SpecialLayersData, arr)
	// }
}

func (scene *Scene) HasSpecialLayer(x, y, z int, layer int) bool {
	for _, v := range scene.SpecialLayers {
		if v.Layer == layer {
			if goutils.FindIntArr(v.Pos, []int{x, y, z}) >= 0 {
				return true
			}
		}
	}

	return false
}

func (scene *Scene) AddSpecialLayers(spl *SpecialLayer) error {
	for _, v := range scene.SpecialLayers {
		if v.LayerType != 0 && v.LayerType == spl.LayerType {
			v.Pos = append(v.Pos, spl.Pos...)

			return nil
		}
	}

	scene.SpecialLayers = append(scene.SpecialLayers, spl)

	return nil
}

func (scene *Scene) resetInitLayerArr(stage *Stage) error {
	scene.InitLayerArr = [][][]int{}

	for z, arr0 := range stage.Layer {
		scene.InitLayerArr = append(scene.InitLayerArr, [][]int{})

		for y, arr1 := range arr0 {
			scene.InitLayerArr[z] = append(scene.InitLayerArr[z], []int{})

			for _, v := range arr1 {
				area := getBlockArea(v)
				block := getBlockSpecialBlock(v)

				if block > 0 {
					scene.InitLayerArr[z][y] = append(scene.InitLayerArr[z][y], 0)
				} else {
					scene.InitLayerArr[z][y] = append(scene.InitLayerArr[z][y], area)
				}
			}
		}
	}

	return nil
}

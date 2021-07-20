package block7game

import (
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhs007/block7/block7pb"
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

// Scene - scene
type Scene struct {
	StageID       int             `json:"stageid"` // 对应missionid，就是关卡id，版本不同，可能没有对比价值
	MapID         int             `json:"mapid"`   // 实际的mapid，有对比价值
	Version       int             `json:"version"`
	SceneID       int64           `json:"sceneid"` // 关卡的动态id，同一个地图，可能随机出不同的scene，这就是随机后的id
	UserID        int64           `json:"userid"`
	Width         int             `json:"width"`
	Height        int             `json:"height"`
	Layers        int             `json:"layers"`
	XOff          int             `json:"xoff"`
	YOff          int             `json:"yoff"`
	Arr           [][][]int       `json:"-"`
	Block         []*BlockData    `json:"-"`
	BlockEx       []*BlockData    `json:"-"`
	MaxBlockNums  int             `json:"-"`
	InitArr       [][][]int       `json:"layer"`
	History       [][]int         `json:"history"`
	ClickValues   int             `json:"clickValues"`
	FinishedPer   float32         `json:"finishedPer"`
	Offset        string          `json:"offset"`
	IsOutputScene bool            `json:"isOutputScene"`
	SpecialLayers []*SpecialLayer `json:"specialLayers"` // 这个是自己用的
	// SpecialLayersData [][]int         `json:"specialLayersData"` // 这个给前端用的
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
		block7utils.Warn("NewScene:MgrSpecial.GenSymbols",
			zap.Error(err))

		return nil, err
	}

	if len(ss) > stage.IconNums {
		block7utils.Warn("NewScene:IconNums",
			zap.Error(ErrInvalidSpecialNums))

		return nil, ErrInvalidSpecialNums
	}

	if len(ss) < stage.IconNums {
		ss1, err := genSymbols(rng, symbols, stage.IconNums-len(ss))
		if err != nil {
			block7utils.Warn("NewScene:genSymbols",
				zap.Error(err))

			return nil, err
		}

		ss = append(ss, ss1...)
	}

	// block7utils.Debug("NewScene",
	// 	block7utils.JSON("symbols", ss))

	scene := &Scene{
		Width:        stage.Width,
		Height:       stage.Height,
		Layers:       len(stage.Layer),
		XOff:         stage.XOff,
		YOff:         stage.YOff,
		MaxBlockNums: blockNums,
		Offset:       stage.Offset,
	}

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
				}
			}

			arrslayer = append(arrslayer, arrsrow)
		}

		scene.Arr = append(scene.Arr, arrslayer)
	}

	scene.InitArr = cloneArr3(scene.Arr)

	err = MgrSpecial.OnFixScene(ld2, scene)
	if err != nil {
		block7utils.Warn("NewScene:OnFixScene",
			zap.Error(err))

		return nil, err
	}

	err = MgrSpecial.GenSymbolLayers(rng, ld2, scene)
	if err != nil {
		block7utils.Warn("NewScene:GenSymbolLayers",
			zap.Error(err))

		return nil, err
	}

	return scene, nil
}

// NewSceneFromPB - new a scene
func NewSceneFromPB(pbscene *block7pb.Scene) (*Scene, error) {
	scene := &Scene{
		MapID:   int(pbscene.MapID2),
		Version: int(pbscene.Version),
		SceneID: pbscene.SceneID,
		Width:   int(pbscene.Width),
		Height:  int(pbscene.Height),
		Layers:  int(pbscene.Layers),
		XOff:    int(pbscene.XOff),
		YOff:    int(pbscene.YOff),
		Offset:  pbscene.Offset,
	}

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

	scene.InitArr = cloneArr3(scene.Arr)

	for _, pbsl := range pbscene.SpecialLayer {
		cs, isok := MgrSpecial.MapSpecial[int(pbsl.Special)]
		if isok {
			sl := &SpecialLayer{
				Layer:     int(pbsl.Layer),
				LayerType: int(pbsl.LayerType),
				Special:   cs,
			}

			arr, err := block7utils.Int32ArrToIntArr2(pbsl.Values, 3, len(pbsl.Values)/3)
			if err != nil {
				block7utils.Warn("NewSceneFromPB:Int32ArrToIntArr2",
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
	scene.Arr = cloneArr3(scene.InitArr)
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
					block7utils.Warn("Scene.analysisNeighboring:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisNeighboring:analysisDepth",
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
			block7utils.Warn("Scene.analysisDepth:analysisNeighboring",
				zap.Error(err))
		}

		if bd.Z%2 == 0 {
			if scene.CanClickEx(bd.X, bd.Y, bd.Z-1, arr) {

				cb, err := mapBI.AddBlockDataEx2(scene, bd.X, bd.Y, bd.Z-1, scene.Arr[bd.Z-1][bd.Y][bd.X], arr)
				if err != nil {
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.analysisDepth:AddBlockDataEx2",
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
							block7utils.Warn("Scene.analysisDepth:analysisDepth",
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
					block7utils.Warn("Scene.Analysis:analysisDepth",
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
		block7utils.Warn("Scene.Click:IsValid",
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
		block7utils.Error("Scene.Save:Marshal",
			zap.Error(err))

		return err
	}

	err = ioutil.WriteFile(fn, buf, os.ModePerm)
	if err != nil {
		block7utils.Error("Scene.Save:WriteFile",
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
		StageID: int32(scene.SceneID),
		MapID2:  int32(scene.MapID),
		Version: int32(scene.Version),
		SceneID: scene.SceneID,
		Width:   int32(scene.Width),
		Height:  int32(scene.Height),
		Layers:  int32(scene.Layers),
		XOff:    int32(scene.XOff),
		YOff:    int32(scene.YOff),
		Offset:  scene.Offset,
	}

	for _, arr2 := range scene.InitArr {
		pblayer := &block7pb.SceneLayer{}

		for _, arr1 := range arr2 {
			pbcolumn := &block7pb.Column{}

			for _, s := range arr1 {
				pbcolumn.Values = append(pbcolumn.Values, int32(s))
			}

			pblayer.Values = append(pblayer.Values, pbcolumn)
		}

		pbScene.InitArr = append(pbScene.InitArr, pblayer)
	}

	for _, sl := range scene.SpecialLayers {
		pbsl := &block7pb.SpecialLayer{
			Special:   int32(sl.Special.GetSpecialID()),
			Layer:     int32(sl.Layer),
			LayerType: int32(sl.LayerType),
		}

		vals, _, _ := block7utils.IntArr2ToInt32Arr(sl.Pos)
		pbsl.Values = vals

		pbScene.SpecialLayer = append(pbScene.SpecialLayer, pbsl)
	}

	return pbScene, nil
}

func (scene *Scene) ToHistoryPB() (*block7pb.Scene, error) {
	pbScene := &block7pb.Scene{
		StageID: int32(scene.SceneID),
		MapID2:  int32(scene.MapID),
		Version: int32(scene.Version),
		SceneID: scene.SceneID,
		UserID:  scene.UserID,
	}

	for _, arr := range scene.History {
		pbcolumn := &block7pb.Column{}

		for _, s := range arr {
			pbcolumn.Values = append(pbcolumn.Values, int32(s))
		}

		pbScene.History = append(pbScene.History, pbcolumn)
	}

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

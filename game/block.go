package block7game

import goutils "github.com/zhs007/goutils"

type BlockData struct {
	X        int          `json:"x"`
	Y        int          `json:"y"`
	Z        int          `json:"z"`
	Symbol   int          `json:"symbol"`
	Parent   []*BlockData `json:"parent"`
	Children []*BlockData `json:"children"`
}

type BlockInfo struct {
	LevelList [][]*BlockData `json:"levellist"`
}

func NewBlockData(x, y, z int, s int) *BlockData {
	return &BlockData{
		X:      x,
		Y:      y,
		Z:      z,
		Symbol: s,
	}
}

func (bd *BlockData) AddParent(p *BlockData) bool {
	if HasBlockData(bd.Parent, p.X, p.Y, p.Z) {
		return false
	}

	bd.Parent = append(bd.Parent, p)

	return true
}

func (bd *BlockData) AddChild(c *BlockData) bool {
	if HasBlockData(bd.Children, c.X, c.Y, c.Z) {
		return false
	}

	bd.Children = append(bd.Children, c)

	return true
}

func NewBlockInfo(maxlevel int) *BlockInfo {
	bi := &BlockInfo{}

	for i := 0; i < maxlevel; i++ {
		bi.LevelList = append(bi.LevelList, nil)
	}

	return bi
}

func HasBlockData(lst []*BlockData, x, y, z int) bool {
	for _, v := range lst {
		if v.X == x && v.Y == y && v.Z == z {
			return true
		}
	}

	return false
}

func DelBlockData(lst []*BlockData, x, y, z int) []*BlockData {
	for i, v := range lst {
		if v.X == x && v.Y == y && v.Z == z {
			return append(lst[0:i], lst[i+1:]...)
		}
	}

	return lst
}

func CountBlockData(lst []*BlockData, symbol int) int {
	n := 0
	for _, v := range lst {
		if v.Symbol == symbol {
			n++
		}
	}

	return n
}

func RemoveBlockData(lst []*BlockData, symbol int, nums int) []*BlockData {
	n := 0
	for i := 0; i < len(lst); {
		if lst[i].Symbol == symbol {
			lst = append(lst[0:i], lst[i+1:]...)

			n++

			if n >= nums {
				return lst
			}
		} else {
			i++
		}
	}

	return lst
}

func RandBlockData(rng IRng, lst []*BlockData, nums int) ([]*BlockData, []*BlockData, error) {
	if nums <= 0 {
		return nil, nil, ErrInvalidParams
	}

	arr := []*BlockData{}
	dst := append([]*BlockData{}, lst...)

	for i := 0; i < nums; i++ {
		cr, err := rng.Rand(len(dst))
		if err != nil {
			return nil, nil, err
		}

		arr = append(arr, dst[cr])
		dst = append(dst[0:cr], dst[cr+1:]...)
	}

	return dst, arr, nil
}

func GetBlockDataList(lst []*BlockData, nums int) ([]*BlockData, []*BlockData, error) {
	if nums <= 0 {
		return nil, nil, ErrInvalidParams
	}

	arr := []*BlockData{}
	dst := append([]*BlockData{}, lst...)

	for i := 0; i < nums; i++ {
		cr := 0

		arr = append(arr, dst[cr])
		dst = append(dst[0:cr], dst[cr+1:]...)
	}

	return dst, arr, nil
}

func FindAllSymbols(arr [][][]int, symbol int) []*BlockData {
	lst := []*BlockData{}

	for z, arr1 := range arr {
		for y, arr2 := range arr1 {
			for x, s := range arr2 {
				if s == symbol {
					lst = append(lst, &BlockData{
						X:      x,
						Y:      y,
						Z:      z,
						Symbol: symbol,
					})
				}
			}
		}
	}

	return lst
}

func FindAllSymbolsEx(arr [][][]int, symbols []int) [][]*BlockData {
	lst := [][]*BlockData{}
	for range symbols {
		lst = append(lst, []*BlockData{})
	}

	for z, arr1 := range arr {
		for y, arr2 := range arr1 {
			for x, s := range arr2 {
				i := goutils.FindInt(symbols, s)
				if i >= 0 {
					lst[i] = append(lst[i], &BlockData{
						X:      x,
						Y:      y,
						Z:      z,
						Symbol: s,
					})
				}
			}
		}
	}

	return lst
}

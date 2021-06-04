package block7

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

func RandBlockData(rng Rng, lst []*BlockData, nums int) ([]*BlockData, []*BlockData, error) {
	if nums <= 0 {
		return nil, nil, ErrInvalidParams
	}

	arr := []*BlockData{}

	for i := 0; i < nums; i++ {
		cr, err := rng.Rand(len(lst))
		if err != nil {
			return nil, nil, err
		}

		arr = append(arr, lst[cr])
		lst = append(lst[0:cr], lst[cr+1:]...)
	}

	return lst, arr, nil
}

package block7

type BlockData struct {
	X, Y, Z int
	Symbol  int
}

type BlockInfo struct {
	LevelList [][]*BlockData
}

func NewBlockData(x, y, z int, s int) *BlockData {
	return &BlockData{
		X:      x,
		Y:      y,
		Z:      z,
		Symbol: s,
	}
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

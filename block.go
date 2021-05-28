package block7

type BlockData struct {
	X, Y, Z int
	Symbol  int
}

type BlockInfo struct {
	L0List []*BlockData
}

func NewBlockData(x, y, z int, s int) *BlockData {
	return &BlockData{
		X:      x,
		Y:      y,
		Z:      z,
		Symbol: s,
	}
}

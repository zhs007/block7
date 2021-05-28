package block7

type BlockData struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Z      int `json:"z"`
	Symbol int `json:"symbol"`
}

type BlockInfo struct {
	L0List []*BlockData `json:"l0list"`
}

func NewBlockData(x, y, z int, s int) *BlockData {
	return &BlockData{
		X:      x,
		Y:      y,
		Z:      z,
		Symbol: s,
	}
}

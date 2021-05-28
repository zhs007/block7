package block7

type BlockList struct {
	arr []*BlockData
}

func NewBlockList() *BlockList {
	return &BlockList{}
}

func (bl *BlockList) AddBlockData(x, y, z int, s int) {
	bl.arr = append(bl.arr, NewBlockData(x, y, z, s))
}

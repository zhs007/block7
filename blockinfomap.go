package block7

type BlockInfoMap struct {
	MapBlockInfo map[int]*BlockInfo `json:"mapBlockInfo"`
}

func NewBlockInfoMap() *BlockInfoMap {
	return &BlockInfoMap{
		MapBlockInfo: make(map[int]*BlockInfo),
	}
}

func (m *BlockInfoMap) AddBlockData(block *BlockData) {
	_, isok := m.MapBlockInfo[block.Symbol]
	if !isok {
		m.MapBlockInfo[block.Symbol] = &BlockInfo{}
	}

	m.MapBlockInfo[block.Symbol].L0List = append(m.MapBlockInfo[block.Symbol].L0List, block)
}

func (m *BlockInfoMap) AddBlockDataEx(x, y, z int, s int) {
	_, isok := m.MapBlockInfo[s]
	if !isok {
		m.MapBlockInfo[s] = &BlockInfo{}
	}

	m.MapBlockInfo[s].L0List = append(m.MapBlockInfo[s].L0List, NewBlockData(x, y, z, s))
}

func (m *BlockInfoMap) OutputLog(msg string) {
	Info(msg,
		JSON("BlockInfoMap", m))
}

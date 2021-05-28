package block7

type BlockInfoMap struct {
	mapBlockInfo map[int]*BlockInfo
}

func NewBlockInfoMap() *BlockInfoMap {
	return &BlockInfoMap{
		mapBlockInfo: make(map[int]*BlockInfo),
	}
}

func (m *BlockInfoMap) AddBlockData(block *BlockData) {
	_, isok := m.mapBlockInfo[block.Symbol]
	if !isok {
		m.mapBlockInfo[block.Symbol] = &BlockInfo{}
	}

	m.mapBlockInfo[block.Symbol].L0List = append(m.mapBlockInfo[block.Symbol].L0List, block)
}

func (m *BlockInfoMap) AddBlockDataEx(x, y, z int, s int) {
	_, isok := m.mapBlockInfo[s]
	if !isok {
		m.mapBlockInfo[s] = &BlockInfo{}
	}

	m.mapBlockInfo[s].L0List = append(m.mapBlockInfo[s].L0List, NewBlockData(x, y, z, s))
}

func (m *BlockInfoMap) OutputLog(msg string) {
	Info(msg,
		JSON("BlockInfoMap", m))
}

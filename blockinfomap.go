package block7

type BlockInfoMap struct {
	maxLevel     int
	mapBlockInfo map[int]*BlockInfo
}

func NewBlockInfoMap(maxlevel int) *BlockInfoMap {
	return &BlockInfoMap{
		maxLevel:     maxlevel,
		mapBlockInfo: make(map[int]*BlockInfo),
	}
}

func (m *BlockInfoMap) AddBlockData(block *BlockData, level int) error {
	if level < 0 || level >= m.maxLevel {
		return ErrInvalidLevel
	}

	_, isok := m.mapBlockInfo[block.Symbol]
	if !isok {
		m.mapBlockInfo[block.Symbol] = NewBlockInfo(m.maxLevel)
	}

	m.mapBlockInfo[block.Symbol].LevelList[level] = append(m.mapBlockInfo[block.Symbol].LevelList[level], block)

	return nil
}

func (m *BlockInfoMap) AddBlockDataEx(x, y, z int, s int, level int) error {
	if level < 0 || level >= m.maxLevel {
		return ErrInvalidLevel
	}

	_, isok := m.mapBlockInfo[s]
	if !isok {
		m.mapBlockInfo[s] = NewBlockInfo(m.maxLevel)
	}

	m.mapBlockInfo[s].LevelList[level] = append(m.mapBlockInfo[s].LevelList[level], NewBlockData(x, y, z, s))

	return nil
}

func (m *BlockInfoMap) OutputLog(msg string) {
	Info(msg,
		JSON("BlockInfoMap", m))
}

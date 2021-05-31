package block7

type BlockInfoMap struct {
	MaxLevel     int                `json:"maxLevel"`
	MapBlockInfo map[int]*BlockInfo `json:"mapBlockInfo"`
}

func NewBlockInfoMap(maxlevel int) *BlockInfoMap {
	return &BlockInfoMap{
		MaxLevel:     maxlevel,
		MapBlockInfo: make(map[int]*BlockInfo),
	}
}

func (m *BlockInfoMap) AddBlockData(block *BlockData, level int) error {
	if level < 0 || level >= m.MaxLevel {
		return ErrInvalidLevel
	}

	_, isok := m.MapBlockInfo[block.Symbol]
	if !isok {
		m.MapBlockInfo[block.Symbol] = NewBlockInfo(m.MaxLevel)
	}

	m.MapBlockInfo[block.Symbol].LevelList[level] = append(m.MapBlockInfo[block.Symbol].LevelList[level], block)

	return nil
}

func (m *BlockInfoMap) AddBlockDataEx(x, y, z int, s int, level int) error {
	if level < 0 || level >= m.MaxLevel {
		return ErrInvalidLevel
	}

	_, isok := m.MapBlockInfo[s]
	if !isok {
		m.MapBlockInfo[s] = NewBlockInfo(m.MaxLevel)
	}

	m.MapBlockInfo[s].LevelList[level] = append(m.MapBlockInfo[s].LevelList[level], NewBlockData(x, y, z, s))

	return nil
}

func (m *BlockInfoMap) OutputLog(msg string) {
	Info(msg,
		JSON("BlockInfoMap", m))
}

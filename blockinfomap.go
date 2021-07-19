package block7

import (
	block7utils "github.com/zhs007/block7/utils"
)

type BlockInfoMap struct {
	MaxLevel     int                `json:"maxLevel"`
	MapBlockInfo map[int]*BlockInfo `json:"mapBlockInfo"`
	BlockSymbols []int              `json:"blockSymbols"`
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

	for _, v := range m.MapBlockInfo[block.Symbol].LevelList {
		if HasBlockData(v, block.X, block.Y, block.Z) {
			return nil
		}
	}

	m.MapBlockInfo[block.Symbol].LevelList[level] = append(m.MapBlockInfo[block.Symbol].LevelList[level], block)

	return nil
}

func (m *BlockInfoMap) AddBlockDataEx(x, y, z int, s int, level int) (*BlockData, error) {
	if level < 0 || level >= m.MaxLevel {
		return nil, ErrInvalidLevel
	}

	_, isok := m.MapBlockInfo[s]
	if !isok {
		m.MapBlockInfo[s] = NewBlockInfo(m.MaxLevel)
	}

	for _, v := range m.MapBlockInfo[s].LevelList {
		if HasBlockData(v, x, y, z) {
			return nil, nil
		}
	}

	b := NewBlockData(x, y, z, s)
	m.MapBlockInfo[s].LevelList[level] = append(m.MapBlockInfo[s].LevelList[level], b)

	return b, nil
}

func (m *BlockInfoMap) AddBlockDataEx2(scene *Scene, x, y, z int, s int, arr []*BlockData) (*BlockData, error) {
	_, isok := m.MapBlockInfo[s]
	if !isok {
		m.MapBlockInfo[s] = NewBlockInfo(m.MaxLevel)
	}

	for _, v := range m.MapBlockInfo[s].LevelList {
		if HasBlockData(v, x, y, z) {
			return nil, nil
		}
	}

	b := NewBlockData(x, y, z, s)

	scene.ProcParent(b, arr)
	level := len(b.Parent)

	m.MapBlockInfo[s].LevelList[level] = append(m.MapBlockInfo[s].LevelList[level], b)

	return b, nil
}

func (m *BlockInfoMap) HasBlockDataEx(x, y, z int, s int, level int) bool {
	if level < 0 || level >= m.MaxLevel {
		return false
	}

	_, isok := m.MapBlockInfo[s]
	if !isok {
		return false
	}

	return HasBlockData(m.MapBlockInfo[s].LevelList[level], x, y, z)
}

func (m *BlockInfoMap) Format() {
	for _, v := range m.MapBlockInfo {
		for i, l := range v.LevelList {
			for _, d := range l {
				m.delSymbol(d.X, d.Y, d.Z, d.Symbol, i+1)
			}
		}
	}
}

func (m *BlockInfoMap) delSymbol(x, y, z int, symbol int, level int) {
	v, isok := m.MapBlockInfo[symbol]
	if isok {
		for i := level; i < m.MaxLevel; i++ {
			for {
				if HasBlockData(v.LevelList[i], x, y, z) {
					v.LevelList[i] = DelBlockData(v.LevelList[i], x, y, z)
				} else {
					break
				}
			}

		}
	}
}

func (m *BlockInfoMap) InsBlockSymbol(s int) {
	for i, v := range m.BlockSymbols {
		if v == s {
			if i > 0 {
				arr := []int{s}
				arr = append(arr, m.BlockSymbols[:i]...)
				arr = append(arr, m.BlockSymbols[i+1:]...)
				m.BlockSymbols = arr
			}

			return
		}
	}

	m.BlockSymbols = append(m.BlockSymbols, s)
}

func (m *BlockInfoMap) OutputLog(msg string) {
	block7utils.Info(msg,
		block7utils.JSON("BlockInfoMap", m))
}

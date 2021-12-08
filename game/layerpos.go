package block7game

type LayerPos struct {
	X, Y, Z int
}

type LayerPosList struct {
	Pos []*LayerPos
}

func (lpl *LayerPosList) AddPos(x, y, z int) {
	for _, v := range lpl.Pos {
		if v.X == x && v.Y == y && v.Z == z {
			return
		}
	}

	lpl.Pos = append(lpl.Pos, &LayerPos{
		X: x,
		Y: y,
		Z: z,
	})
}

type LayerPosMap struct {
	MapLayerPos map[int]*LayerPosList
}

func NewLayerPosMap() *LayerPosMap {
	return &LayerPosMap{
		MapLayerPos: make(map[int]*LayerPosList),
	}
}

func (mapLayerPos *LayerPosMap) AddPos(x, y, z int, area int) {
	lpl, isok := mapLayerPos.MapLayerPos[area]
	if !isok {
		lpl = &LayerPosList{}
		mapLayerPos.MapLayerPos[area] = lpl
	}

	lpl.AddPos(x, y, z)
}

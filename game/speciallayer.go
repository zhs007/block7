package block7game

type SpecialLayer struct {
	Layer     int      `json:"layer"`
	LayerType int      `json:"layerType"`
	Special   ISpecial `json:"-"`
	Pos       [][]int  `json:"pos"`
}

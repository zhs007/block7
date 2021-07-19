package block7game

var MgrSpecial *SpecialMgr

func GenSpecialSymbols(lststd []SpecialTypeData) []int {
	arr := []int{}

	return arr
}

func init() {
	MgrSpecial = NewSpecialMgr()

	MgrSpecial.RegSpecial(8, NewBomb(8, 303))
	MgrSpecial.RegSpecial(9, NewBomb(9, 306))
	MgrSpecial.RegSpecial(10, NewBomb(10, 307))
	MgrSpecial.RegSpecial(11, NewBomb(11, 308))
}

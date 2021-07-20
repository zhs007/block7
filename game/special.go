package block7game

var MgrSpecial *SpecialMgr

func GenSpecialSymbols(lststd []SpecialTypeData) []int {
	arr := []int{}

	return arr
}

func init() {
	MgrSpecial = NewSpecialMgr()

	MgrSpecial.RegSpecial(3, NewQuestion(1, 10300))

	MgrSpecial.RegSpecial(1, NewIce(1, 10101, 1))
	MgrSpecial.RegSpecial(14, NewIce(14, 10102, 2))
	MgrSpecial.RegSpecial(15, NewIce(15, 10103, 3))
	MgrSpecial.RegSpecial(16, NewIce(16, 10104, 4))

	MgrSpecial.RegSpecial(2, NewCake(2, 403, 301))

	MgrSpecial.RegSpecial(8, NewBomb(8, 303))
	MgrSpecial.RegSpecial(9, NewBomb(9, 306))
	MgrSpecial.RegSpecial(10, NewBomb(10, 307))
	MgrSpecial.RegSpecial(11, NewBomb(11, 308))
}

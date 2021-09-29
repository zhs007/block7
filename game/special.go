package block7game

var MgrSpecial *SpecialMgr

func GenSpecialSymbols(lststd []SpecialTypeData) []int {
	arr := []int{}

	return arr
}

func init() {
	MgrSpecial = NewSpecialMgr()

	MgrSpecial.RegSpecial(7, NewIceCream(7, 10703, 302, 3, 2))

	MgrSpecial.RegSpecial(17, NewReverse(17, 10200))

	MgrSpecial.RegSpecial(12, NewTeleport(12, 305))

	MgrSpecial.RegSpecial(13, NewRainbow(13, 304))

	MgrSpecial.RegSpecial(6, NewWeeds(6, 10600))

	MgrSpecial.RegSpecial(5, NewLadybug(5, 10503, 3))

	MgrSpecial.RegSpecial(4, NewCurtain(4, 10400))

	MgrSpecial.RegSpecial(3, NewQuestion(3, 10300))

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

func SpecialType2SpecialID(st int) int {
	if st == 10703 {
		return 7
	} else if st == 10200 {
		return 17
	} else if st == 10600 {
		return 6
	} else if st == 10503 {
		return 5
	} else if st == 10400 {
		return 4
	} else if st == 10300 {
		return 3
	} else if st == 10101 {
		return 1
	} else if st == 10102 {
		return 14
	} else if st == 10103 {
		return 15
	} else if st == 10104 {
		return 16
	}

	return 0
}

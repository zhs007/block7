package block7

func genSymbols(rng Rng, symbols []int, nums int) ([]int, error) {
	if nums%BlockNums > 0 {
		return nil, ErrInvalidSymbolsLength
	}

	sn := len(symbols)
	sn1 := nums / BlockNums

	arr := []int{}

	sn2 := sn1 / sn
	sn3 := sn1 - sn*sn2

	for i := 0; i < sn2; i++ {
		for j := 0; j < sn; j++ {
			for k := 0; k < BlockNums; k++ {
				arr = append(arr, symbols[j])
			}
		}
	}

	if sn3 >= 0 {
		for i := 0; i < sn3; i++ {
			j, err := rng.Rand(len(symbols))
			if err != nil {
				return nil, err
			}

			for k := 0; k < BlockNums; k++ {
				arr = append(arr, symbols[j])
			}

			symbols = append(symbols[:j], symbols[j+1:]...)
		}
	}

	return arr, nil
}

func randSymbols(rng Rng, symbols []int) ([]int, int, error) {
	if len(symbols) <= 0 {
		return nil, 0, ErrInvalidSymbolsLength
	}

	si, err := rng.Rand(len(symbols))
	if err != nil {
		return nil, 0, err
	}

	c := symbols[si]
	symbols = append(symbols[:si], symbols[si+1:]...)

	return symbols, c, nil
}

func countSymbols(symbols []int, symbol int) int {
	n := 0
	for _, v := range symbols {
		if v == symbol {
			n++
		}
	}

	return n
}

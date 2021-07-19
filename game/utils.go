package block7game

func GenSymbols(nums int) []int {
	arr := []int{}

	for i := 1; i <= nums; i++ {
		arr = append(arr, i)
	}

	return arr
}

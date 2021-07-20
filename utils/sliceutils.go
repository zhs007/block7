package block7utils

// FindInt - find a int into []int
func FindInt(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}

	return -1
}

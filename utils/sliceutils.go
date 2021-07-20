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

// FindInt3 - find 3 ints into []int
func FindInt3(arr []int, x, y, z int) int {
	if len(arr)%3 == 0 {
		for i := 0; i < len(arr)/3; i++ {
			if arr[i*3] == x && arr[i*3+1] == y && arr[i*3+2] == z {
				return i * 3
			}
		}
	}

	return -1
}

// FindIntArr - find a int array into [][]int
func FindIntArr(arr [][]int, vals []int) int {
	for i, arr1 := range arr {
		if len(arr1) == len(vals) {
			isok := true
			for j, av := range arr1 {
				if av != vals[j] {
					isok = false
					break
				}
			}

			if isok {
				return i
			}
		}
	}

	return -1
}

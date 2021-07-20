package block7utils

import "go.uber.org/zap"

// Int32ArrToIntArr2 - []int32 to [][]int
func Int32ArrToIntArr2(arr []int32, x, y int) ([][]int, error) {
	arr2 := [][]int{}

	if len(arr) != x*y {
		Error("Int32ArrToIntArr2",
			zap.Int("len", len(arr)),
			zap.Int("x", x),
			zap.Int("y", y),
			zap.Error(ErrInvalidArrayLength))

		return nil, ErrInvalidArrayLength
	}

	for i := 0; i < len(arr)/x; i++ {
		carr := []int{}

		for j := 0; j < x; j++ {
			carr = append(carr, int(arr[i*x+j]))
		}

		arr2 = append(arr2, carr)
	}

	return arr2, nil
}

// IntArr2ToInt32Arr - [][]int to []int32
func IntArr2ToInt32Arr(arr [][]int) ([]int32, int, int) {
	arr2 := []int32{}

	for _, arr1 := range arr {
		for _, v := range arr1 {
			arr2 = append(arr2, int32(v))
		}
	}

	return arr2, len(arr[0]), len(arr)
}

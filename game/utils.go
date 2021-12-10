package block7game

import "strings"

func GenSymbols(nums int) []int {
	arr := []int{}

	for i := 1; i <= nums; i++ {
		arr = append(arr, i)
	}

	return arr
}

func genSymbols(rng IRng, symbols []int, nums int) ([]int, error) {
	if nums%BlockNums > 0 {
		return nil, ErrInvalidSymbolsLength
	}

	sn := len(symbols)
	sn1 := nums / BlockNums
	if sn1 < sn {
		sn = sn1
		symbols = GenSymbols(sn1)
	}

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

func randSymbols(rng IRng, symbols []int) ([]int, int, error) {
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

// func insBlockData(arr []*BlockData, b *BlockData) []*BlockData {
// 	arr = append(arr, b)

// 	return arr
// }

func insBlockDataAndProc(arr []*BlockData, b *BlockData) []*BlockData {
	arr = append(arr, b)
	cn := CountBlockData(arr, b.Symbol)
	if cn >= BlockNums {
		arr = RemoveBlockData(arr, b.Symbol, BlockNums*cn/BlockNums)
	}

	return arr
}

func CheckScene(scene *Scene, isMyBlock FuncHasBlock) [][]int {
	lst := [][]int{}

	for z, arr2 := range scene.InitArr {
		for y, arr1 := range arr2 {
			for x := range arr1 {
				if isMyBlock(x, y, z) {
					lst = append(lst, []int{x, y, z})
				}
			}
		}
	}

	return lst
}

// GetChildrenEx - 统计子节点数量，pos 是 []int{x,y,z} 的数组
func GetChildrenEx(scene *Scene, pos [][]int, w, h int, funcHasBlock FuncHasBlock) []int {
	lst := []int{}

	for _, cp := range pos {
		for ox := 0; ox < w; ox++ {
			for oy := 0; oy < h; oy++ {
				lst = scene.GetChildren(lst, cp[0]+ox, cp[1]+oy, cp[2], funcHasBlock)
			}
		}
	}

	return lst
}

func OffsetStringToXYOff(offset string) (int, int) {
	if len(offset) > 0 {
		arr := strings.Split(offset, ",")
		if len(arr) == 3 {
			if arr[0] == "0" {
				return 1, -1
			} else {
				return -1, 1
			}
		}
	}

	return 1, -1
}

// func cloneArr3(src [][][]int) [][][]int {
// 	arr := [][][]int{}

// 	for _, src2 := range src {
// 		arr2 := [][]int{}

// 		for _, src1 := range src2 {
// 			arr1 := append([]int{}, src1[0:]...)
// 			arr2 = append(arr2, arr1)
// 		}

// 		arr = append(arr, arr2)
// 	}

// 	return arr
// }

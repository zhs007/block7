package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

func GenBlocks(rng IRng, scene *Scene, nums int, funcHasBlock FuncHasBlock, funcCanAdd FuncHasBlock) ([][]int, error) {
	lst := [][]int{}
	lstpos := []int{}

	for z, arr1 := range scene.InitArr {
		for y, arr2 := range arr1 {
			for x := range arr2 {
				if funcHasBlock(x, y, z) {
					lstpos = append(lstpos, x, y, z)
				}
			}
		}
	}

	if len(lstpos)/3 < nums {
		block7utils.Error("GenBlocks",
			zap.Int("validpos", len(lstpos)/3),
			zap.Int("nums", nums),
			zap.Error(ErrInvalidGenBrotherBlocksNums))

		return nil, ErrInvalidGenBrotherBlocksNums
	}

	for i := 0; i < nums; i++ {
		cr, err := rng.Rand(len(lstpos) / 3)
		if err != nil {
			block7utils.Error("GenBlocks",
				zap.Int("validpos", len(lstpos)/3),
				zap.Int("nums", nums),
				zap.Int("i", i),
				zap.Error(err))

			return nil, err
		}

		cx := lstpos[cr*3]
		cy := lstpos[cr*3+1]
		cz := lstpos[cr*3+2]

		lstpos = append(lstpos[0:cr*3], lstpos[(cr+1)*3:]...)
		if len(lstpos)/3 < nums {
			block7utils.Error("GenBlocks",
				zap.Int("validpos", len(lstpos)/3),
				zap.Int("nums", nums),
				zap.Int("i", i),
				zap.Error(ErrInvalidGenBrotherBlocksNums))

			return nil, ErrInvalidGenBrotherBlocksNums
		}

		if !funcCanAdd(cx, cy, cz) {
			i--

			continue
		}

		lst = append(lst, []int{
			cx,
			cy,
			cz,
		})
	}

	return lst, nil
}

func GetAllBlocksEx(scene *Scene, w, h int, funcCanAdd FuncHasBlock, funcCanAddEx FuncHasBlockEx) ([][]int, error) {
	lst := [][]int{}

	for z, arr1 := range scene.InitArr {
		for y, arr2 := range arr1 {
			for x := range arr2 {
				isok := true
				for ox := 0; ox < w; ox++ {
					for oy := 0; oy < h; oy++ {
						if !funcCanAdd(x+ox, y+oy, z) {
							isok = false
							break
						}
					}

					if !isok {
						break
					}
				}

				if isok && funcCanAddEx(x, y, z, w, h) {
					lst = append(lst, []int{x, y, z})
				}
			}
		}
	}

	return lst, nil
}

func GetAllBlocks(scene *Scene, funcCanAdd FuncHasBlock) ([][]int, error) {
	lst := [][]int{}

	for z, arr1 := range scene.InitArr {
		for y, arr2 := range arr1 {
			for x := range arr2 {
				if funcCanAdd(x, y, z) {
					lst = append(lst, []int{x, y, z})
				}
			}
		}
	}

	return lst, nil
}

func AddBlocksVal(lst [][]int, countBlockValue FuncCountBlockValue) ([][]int, error) {
	narr := [][]int{}
	for i, pos := range lst {
		if len(pos) < 3 {
			block7utils.Error("GenBlocks",
				block7utils.JSON("pos", pos),
				zap.Int("i", i),
				zap.Error(ErrInvalidPositionList))

			return nil, ErrInvalidPositionList
		}

		v := countBlockValue(pos[0], pos[1], pos[2])

		carr1 := append(pos, v)
		narr = append(narr, carr1)
	}

	return narr, nil
}

func SortBlocks(lst [][]int, isLess FuncIsLess) {
}

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

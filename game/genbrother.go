package block7game

import (
	block7utils "github.com/zhs007/block7/utils"
	"go.uber.org/zap"
)

func GenBrotherBlocks(rng IRng, scene *Scene, brother int, nums int, funcHasBlock FuncHasBlock, funcCanAdd FuncHasBlock) ([][]int, error) {
	lst := [][]int{}
	lstpos := []int{}

	for z, arr1 := range scene.InitArr {
		for y, arr2 := range arr1 {
			for x := range arr2 {
				cv := 0

				if funcHasBlock(x-1, y, z) {
					cv++
				}

				if funcHasBlock(x+1, y, z) {
					cv++
				}

				if funcHasBlock(x, y-1, z) {
					cv++
				}

				if funcHasBlock(x, y+1, z) {
					cv++
				}

				if cv >= brother {
					lstpos = append(lstpos, x, y, z)
				}
			}
		}
	}

	if len(lstpos)/3 < nums {
		block7utils.Error("GenBrotherBlocks",
			zap.Int("validpos", len(lstpos)/3),
			zap.Int("nums", nums),
			zap.Error(ErrInvalidGenBrotherBlocksNums))

		return nil, ErrInvalidGenBrotherBlocksNums
	}

	for i := 0; i < nums; i++ {
		cr, err := rng.Rand(len(lstpos) / 3)
		if err != nil {
			block7utils.Error("GenBrotherBlocks",
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
			block7utils.Error("GenBrotherBlocks",
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

		ci := block7utils.FindInt3(lstpos, cx+1, cy, cz)
		if ci >= 0 {
			lstpos = append(lstpos[0:ci], lstpos[ci+3:]...)
		}

		ci = block7utils.FindInt3(lstpos, cx-1, cy, cz)
		if ci >= 0 {
			lstpos = append(lstpos[0:ci], lstpos[ci+3:]...)
		}

		ci = block7utils.FindInt3(lstpos, cx, cy-1, cz)
		if ci >= 0 {
			lstpos = append(lstpos[0:ci], lstpos[ci+3:]...)
		}

		ci = block7utils.FindInt3(lstpos, cx, cy+1, cz)
		if ci >= 0 {
			lstpos = append(lstpos[0:ci], lstpos[ci+3:]...)
		}

		if len(lstpos)/3 < nums {
			block7utils.Error("GenBrotherBlocks",
				zap.Int("validpos", len(lstpos)/3),
				zap.Int("nums", nums),
				zap.Int("i", i),
				zap.Error(ErrInvalidGenBrotherBlocksNums))

			return nil, ErrInvalidGenBrotherBlocksNums
		}
	}

	return lst, nil
}

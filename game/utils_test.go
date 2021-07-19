package block7game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_genSymbols(t *testing.T) {
	rng := NewRngNormal()

	_, err := genSymbols(rng, []int{1, 2, 3, 4, 5, 6, 7, 8}, 73)
	assert.NotNil(t, err)

	symbols, err := genSymbols(rng, []int{1, 2, 3, 4, 5, 6, 7, 8}, 72)
	assert.Nil(t, err)
	assert.Equal(t, countSymbols(symbols, 1), 9)
	assert.Equal(t, countSymbols(symbols, 2), 9)
	assert.Equal(t, countSymbols(symbols, 3), 9)
	assert.Equal(t, countSymbols(symbols, 4), 9)
	assert.Equal(t, countSymbols(symbols, 5), 9)
	assert.Equal(t, countSymbols(symbols, 6), 9)
	assert.Equal(t, countSymbols(symbols, 7), 9)
	assert.Equal(t, countSymbols(symbols, 8), 9)

	symbols, err = genSymbols(rng, []int{1, 2, 3, 4, 5, 6, 7, 8}, 75)
	assert.Nil(t, err)
	n9 := 0
	n12 := 0
	nother := 0
	for s := 1; s < 9; s++ {
		cn := countSymbols(symbols, s)
		if cn == 9 {
			n9++
		} else if cn == 12 {
			n12++
		} else {
			nother++
		}
	}
	assert.Equal(t, n9, 7)
	assert.Equal(t, n12, 1)
	assert.Equal(t, nother, 0)

	symbols, err = genSymbols(rng, []int{1, 2, 3, 4, 5, 6, 7, 8}, 93)
	assert.Nil(t, err)
	n9 = 0
	n12 = 0
	nother = 0
	for s := 1; s < 9; s++ {
		cn := countSymbols(symbols, s)
		if cn == 9 {
			n9++
		} else if cn == 12 {
			n12++
		} else {
			nother++
		}
	}
	assert.Equal(t, n9, 1)
	assert.Equal(t, n12, 7)
	assert.Equal(t, nother, 0)

	t.Logf("Test_genSymbols OK")
}

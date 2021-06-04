package block7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InsBlockSymbol(t *testing.T) {
	mapBI := NewBlockInfoMap(2)

	mapBI.InsBlockSymbol(1)
	mapBI.InsBlockSymbol(2)
	mapBI.InsBlockSymbol(3)
	mapBI.InsBlockSymbol(2)

	assert.Equal(t, len(mapBI.BlockSymbols), 3)
	assert.Equal(t, mapBI.BlockSymbols[0], 2)
	assert.Equal(t, mapBI.BlockSymbols[1], 1)
	assert.Equal(t, mapBI.BlockSymbols[2], 3)

	mapBI.BlockSymbols = nil

	mapBI.InsBlockSymbol(1)
	mapBI.InsBlockSymbol(1)
	mapBI.InsBlockSymbol(2)
	mapBI.InsBlockSymbol(2)
	mapBI.InsBlockSymbol(3)
	mapBI.InsBlockSymbol(3)

	assert.Equal(t, len(mapBI.BlockSymbols), 3)
	assert.Equal(t, mapBI.BlockSymbols[0], 3)
	assert.Equal(t, mapBI.BlockSymbols[1], 2)
	assert.Equal(t, mapBI.BlockSymbols[2], 1)

	mapBI.BlockSymbols = nil

	mapBI.InsBlockSymbol(1)
	mapBI.InsBlockSymbol(2)
	mapBI.InsBlockSymbol(2)
	mapBI.InsBlockSymbol(3)

	assert.Equal(t, len(mapBI.BlockSymbols), 3)
	assert.Equal(t, mapBI.BlockSymbols[0], 2)
	assert.Equal(t, mapBI.BlockSymbols[1], 1)
	assert.Equal(t, mapBI.BlockSymbols[2], 3)

	t.Logf("Test_InsBlockSymbol OK")
}

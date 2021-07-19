package block7game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewScene(t *testing.T) {
	rng := NewRngNormal()

	stage, err := LoadStage("./cfg/level_0100.json")
	assert.Nil(t, err)

	scene, err := NewScene(rng, stage, []int{1, 2, 3, 4, 5, 6, 7, 8}, DefaultMaxBlockNums, nil)
	assert.Nil(t, err)

	assert.Equal(t, scene.CountSymbols(), 120)
	assert.Equal(t, len(scene.Arr), 4)

	for _, arrlayer := range scene.Arr {
		assert.Equal(t, len(arrlayer), 10)
		for _, arrrow := range arrlayer {
			assert.Equal(t, len(arrrow), 9)
		}
	}

	for s := 1; s <= 8; s++ {
		assert.Equal(t, scene.CountSymbol(s), 15)
	}

	t.Logf("Test_NewScene OK")
}

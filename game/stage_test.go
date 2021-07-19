package block7game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadStage(t *testing.T) {
	stage, err := LoadStage("../cfg/level_0100.json")
	assert.Nil(t, err)

	assert.Equal(t, stage.Width, 9)
	assert.Equal(t, stage.Height, 10)
	assert.Equal(t, len(stage.Layer), 4)

	for _, arrlayer := range stage.Layer {
		assert.Equal(t, len(arrlayer), 10)
		for _, arrrow := range arrlayer {
			assert.Equal(t, len(arrrow), 9)
		}
	}

	assert.Equal(t, stage.CountSymbols(), 120)

	t.Logf("Test_LoadStage OK")
}

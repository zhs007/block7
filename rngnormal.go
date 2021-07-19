package block7

import (
	"math/rand"
	"time"
)

var isBasicPluginInited = false

// RngNormal -
type RngNormal struct {
}

// NewRngNormal - new RngNormal
func NewRngNormal() Rng {
	if !isBasicPluginInited {
		rand.Seed(time.Now().UnixNano())

		isBasicPluginInited = true
	}

	return &RngNormal{}
}

// Rand - rand
func (rng *RngNormal) Rand(r int) (int, error) {
	ci := rand.Int()
	return ci % r, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())

	isBasicPluginInited = true
}

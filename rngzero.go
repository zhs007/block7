package block7

// RngZero -
type RngZero struct {
}

// NewRngZero - new RngZero
func NewRngZero() Rng {
	return &RngZero{}
}

// Rand - rand
func (rng *RngZero) Rand(r int) (int, error) {
	return 0, nil
}

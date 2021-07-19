package block7game

// RngZero -
type RngZero struct {
}

// NewRngZero - new RngZero
func NewRngZero() IRng {
	return &RngZero{}
}

// Rand - rand
func (rng *RngZero) Rand(r int) (int, error) {
	return 0, nil
}

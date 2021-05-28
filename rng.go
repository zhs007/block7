package block7

type Rng interface {
	Rand(r int) (int, error)
}

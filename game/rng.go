package block7game

type IRng interface {
	Rand(r int) (int, error)
}

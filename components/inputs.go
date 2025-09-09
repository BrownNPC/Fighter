package c


type Input uint8
type Inputs []Input

const None Input = 0
const (
	Up Input = 1 << iota
	Down
	Left
	Right
)

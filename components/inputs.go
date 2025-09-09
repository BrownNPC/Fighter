package c

import "slices"

type Input uint8
type Inputs []Input

const None Input = 0
const (
	Up Input = 1 << iota
	Down
	Left
	Right
)

func (inp Inputs) Match(with ...Input) bool {
	if len(with) > len(inp) {
		return false
	}
	inp[:len(with)]
	slices.Equal()
	for _, input := range inp {
	}
}

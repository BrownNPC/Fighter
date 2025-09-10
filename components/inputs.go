package c

type Input uint8

const INeutral Input = 0
const (
	IUp Input = 1 << iota
	IDown
	IForward
	IBack
)

// If the input Contains the input you want.
// eg. if this input is Forward | Down and you want to check
// if forward was pressed, you can use this function.
func (i Input) Contains(j Input) bool {
	return i&j == j
}

package input

type Input uint8

const Neutral Input = 0
const (
	Up Input = 1 << iota
	Down
	Forward
	Back
	Attack
	Block
)

// If the input Contains the input you want.
// eg. if this input is Forward | Down and you want to check
// if forward was pressed, you can use this function.
//
//  NOTE: Neutral is strictly matched.
func (i Input) Contains(j Input) bool {
	if j == Neutral {
		return i == Neutral
	} else {
		return i&j == j
	}
}

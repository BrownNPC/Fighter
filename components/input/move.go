package input


type Move struct {
	Sequence []Input
	strict   bool
	leniency uint8
}

// NewMove creates a new Move from the sequence.
//
// leniency is the number of frames allowed to perform this move
//
// strict = true will disallow cases like
//
//	Down | Forward == Forward
func NewMove(strict bool, leniency uint8, sequence ...Input) Move {
	return Move{
		Sequence: sequence,
		strict:   strict,
		leniency: uint8(leniency),
	}
}

// A MoveGroup is useful to check different inputs for each move. allowing for leniancy.
//
// eg. Down, Forward | Down, Forward, Punch
//
// and Down, Forward, Forward | Punch
// would both perform a haduken.
type MoveGroup []Move

// Check if one of the Move sequences were performed.
func (g MoveGroup) Check(buf *InputBuffer) bool {
	for _, move := range g {
		if buf.CheckSequence(move) {
			return true
		}
	}
	return false
}

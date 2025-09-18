package input

import "slices"

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
	leniency = max(uint8(len(sequence)), leniency)
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
// And automatically clear the performed move from the buffer
func (g MoveGroup) Check(buf *InputBuffer) bool {
	return slices.ContainsFunc(g, func(E Move) bool {
		if buf.CheckSequence(E) {
			buf.ClearSequence(E)
			return true
		}
		return false
	})
}

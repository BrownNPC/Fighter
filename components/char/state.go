package char

import "GameFrameworkTM/components/input"

type StateType uint8

const (
	StateIdle StateType = iota
	StateJumping
	StateWalking
)
// CharacterState modifies a character object in some way by reading input and
// returns a CharacterState, that could be itself.
// This is supposed to be a method on a Character object.
type CharacterState interface {
	HandleInput(*input.InputBuffer) CharacterState
}

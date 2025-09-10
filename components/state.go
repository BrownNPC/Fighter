package c

type StateType uint8

const (
	StateIdle StateType = iota
	StateJumping
	StateWalking
)

// CharacterState modifies a character object in some way by reading input and
// returns a CharacterState, that could be itself
type CharacterState[CharacterType any] interface {
	HandleInput(Input) CharacterState[CharacterType]
}

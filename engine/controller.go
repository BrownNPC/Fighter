package engine

import (
	"GameFrameworkTM/components/input"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Control provides easy to type aliases for raylib gamepad inputs
type Control = int32
type Direction bool

const (
	// Left means character is looking left.
	DLeft = Direction(true)
	// Right means character is looking right.
	DRight = Direction(false)
)

const (
	None Control = rl.GamepadButtonUnknown
	// D-Pad
	Up    = rl.GamepadButtonLeftFaceUp
	Down  = rl.GamepadButtonLeftFaceDown
	Left  = rl.GamepadButtonLeftFaceLeft
	Right = rl.GamepadButtonLeftFaceRight
	// Light attack
	Light
	// Heavy attack
	Heavy
	Block

	TotalInputs
)

// Maps Keyboard,and Gamepad input to virtual input
var Keymap map[int32]Control

func init() {
	Keymap = map[int32]Control{
		rl.GamepadButtonLeftFaceUp:    Up,
		rl.GamepadButtonLeftFaceDown:  Down,
		rl.GamepadButtonLeftFaceLeft:  Left,
		rl.GamepadButtonLeftFaceRight: Right,
	}
}

var Player1, Player2 input.InputBuffer

func UpdatePlayer1(facing Direction) {
	checkGamepad(0, facing)
}
func checkGamepad(id int32, facing Direction) input.Input {
	var inputs []input.Input = make([]input.Input, TotalInputs)
	if rl.IsGamepadButtonDown(id, Left) {
		inputs = append(inputs, checkDirectionalInput(Left, facing))
	}
	if rl.IsGamepadButtonDown(id, Right) {
		inputs = append(inputs, checkDirectionalInput(Right, facing))
	}
	if rl.IsGamepadButtonDown(id, Up) {
		inputs = append(inputs, Up)
	}
	if rl.IsGamepadButtonDown(id, Up) {
		inputs = append(inputs, Up)
	}
	var allInputs = input.INeutral
	for _, inp := range inputs {
		allInputs |= inp
	}
	return allInputs
}
func checkDirectionalInput(button Control, facing Direction) input.Input {
	switch button {
	case Right:
		if facing == DRight {
			return input.IForward
		} else {
			return input.IBack
		}
	case Left:
		if facing == DLeft {
			return input.IForward
		} else {
			return input.IBack
		}
	}
	panic("invalid button passed")
}

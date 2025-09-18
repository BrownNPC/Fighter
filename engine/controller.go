package engine

import (
	"GameFrameworkTM/components/input"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Control provides easy to type aliases for raylib gamepad inputs
type Control int32
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
	Up    Control = rl.GamepadButtonLeftFaceUp
	Down  Control = rl.GamepadButtonLeftFaceDown
	Left  Control = rl.GamepadButtonLeftFaceLeft
	Right Control = rl.GamepadButtonLeftFaceRight
	// Square / X
	// Attack
	Attack Control = rl.GamepadButtonRightFaceLeft
	// Triangle / Y
	Block Control = rl.GamepadButtonRightFaceUp
)

// Controller input -> raylib Keyboard const
type Keymap map[Control]int32

var Keymap_Player1, Keymap_Player2 Keymap

// NOTE: this does not account for Left or Right.
// use checkDirectionalInput for that
var controlToInput = map[Control]input.Input{
	Up:     input.Up,
	Down:   input.Down,
	None:   input.Neutral,
	Attack: input.Attack,
	Block:  input.Block,
}

func init() {
	Keymap_Player1 = map[Control]int32{
		Up:     rl.KeyW,
		Down:   rl.KeyS,
		Left:   rl.KeyA,
		Right:  rl.KeyD,
		Attack: rl.KeyC,
		Block:  rl.KeyV,
	}
	Keymap_Player2 = map[Control]int32{
		Up:     rl.KeyI,
		Down:   rl.KeyK,
		Left:   rl.KeyJ,
		Right:  rl.KeyL,
		Attack: rl.KeyPeriod,
		Block:  rl.KeySlash,
	}
}

var Player1, Player2 input.InputBuffer

func UpdatePlayers(facingP1, facingP2 Direction) {
	var inp input.Input = input.Neutral
	// Player 1
	inp |= checkInput(Left, false, 0, Keymap_Player1, facingP1)
	inp |= checkInput(Right, false, 0, Keymap_Player1, facingP1)
	inp |= checkInput(Up, false, 0, Keymap_Player1, facingP1)
	inp |= checkInput(Down, false, 0, Keymap_Player1, facingP1)
	inp |= checkInput(Attack, true, 0, Keymap_Player1, facingP1)
	inp |= checkInput(Block, false, 0, Keymap_Player1, facingP1)
	Player1.Add(inp)

	inp = input.Neutral
	inp |= checkInput(Left, false, 1, Keymap_Player2, facingP2)
	inp |= checkInput(Right, false, 1, Keymap_Player2, facingP2)
	inp |= checkInput(Up, false, 1, Keymap_Player2, facingP2)
	inp |= checkInput(Down, false, 1, Keymap_Player2, facingP2)
	inp |= checkInput(Attack, true, 1, Keymap_Player2, facingP2)
	inp |= checkInput(Block, false, 1, Keymap_Player2, facingP2)
	Player2.Add(inp)
}

// pressed is whether to check if button is held down, or pressed.
func checkInput(button Control, pressed bool, padId int32, keyMap Keymap, facing Direction) input.Input {
	var down bool
	var keyFunc, gamePadFunc = rl.IsKeyDown, rl.IsGamepadButtonDown
	// pressed is whether to check if button is held down, or pressed.
	if pressed {
		keyFunc = rl.IsKeyPressed
		gamePadFunc = rl.IsGamepadButtonPressed
	}
	if keyFunc(keyMap[button]) {
		down = true
	}
	if gamePadFunc(padId, int32(button)) {
		down = true
	}
	if down {
		if button == Left || button == Right {
			return checkDirectionalInput(button, facing)
		}
		if ctrl, ok := controlToInput[button]; ok {
			return ctrl
		} else {
			panic("asked to check a control not found in control map " + fmt.Sprint(button))
		}
	}
	// OR with 0 is a no-op. INeutral is 0.
	return input.Neutral
}

func checkDirectionalInput(button Control, facing Direction) input.Input {
	switch button {
	case Right:
		if facing == DRight {
			return input.Forward
		} else {
			return input.Backward
		}
	case Left:
		if facing == DLeft {
			return input.Forward
		} else {
			return input.Backward
		}
	}
	panic("invalid button passed")
}

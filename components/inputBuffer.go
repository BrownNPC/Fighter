package c

import "GameFrameworkTM/components/frame"

type Input uint8

const (
	// buffer last 10 inputs
	INPUT_BUFFER_SIZE = frame.Frame(10)
	// clear after 6 frames
	INPUT_BUFFER_FRAMES = frame.Frame(6)
)

const (
	None Input = iota
	Up
	Down
	Left
	Right
)

type FrameInput struct {
	Frame frame.Frame
}

type InputBuffer struct {
}

package c

import (
	"GameFrameworkTM/components/frame"
	"slices"
)


const (
	// circular buffer size, how many inputs it can store.
	InputBufferSize = 32
	// clear after 6 frames
	InputBufferClearAfter = frame.Frame(6)
)

type FrameInput struct {
	Input Input
	Frame frame.Frame
}

// InputBuffer is a circular buffer of inputs.
// A zero value InputBuffer is ready to use.
type InputBuffer struct {
	// circular buffer
	buf [InputBufferSize]FrameInput
	// where to add the input next
	cursor int
}

func (b *InputBuffer) Add(i Input) {
	b.buf[b.cursor] = FrameInput{
		Frame: frame.Now(),
		Input: i,
	}
	b.cursor = (b.cursor + 1) % InputBufferSize
}

// GetPrevious inputs from now till however many frames ago
func (b *InputBuffer) GetPrevious(tillFramesAgo frame.Frame) Inputs {
	var totalInputs = make([]Input, 0, InputBufferSize)
	for _, frameInput := range b.buf {
		if frame.Since(frameInput.Frame) > tillFramesAgo {
			continue
		}
		totalInputs = append(totalInputs, frameInput.Input)
	}
	slices.Reverse(totalInputs)
	return totalInputs
}

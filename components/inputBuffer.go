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
	// start at the last write, and then walk backwards until
	// we get an input that is too old.
	var cursor int
	for range InputBufferSize {
		// we move back from here, because the cursor does not represent the last write
		// but the next write.
		cursor = (b.cursor - 1) % InputBufferSize
		frameInput := b.buf[cursor]

		// too old or uninitialized, break.
		if frame.Since(frameInput.Frame) > tillFramesAgo {
			break
		}

		totalInputs = append(totalInputs, frameInput.Input)
	}
	// newest to oldest sort
	slices.Reverse(totalInputs)
	return totalInputs
}

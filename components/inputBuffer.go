package c

import (
	"GameFrameworkTM/components/frame"
)

const (
	// circular buffer size, how many inputs it can store.
	InputBufferSize = 60
	// clear after 6 frames
	InputBufferClearAfter = frame.Frame(6)
)

type FrameInput struct {
	Input Input
}

// InputBuffer is a circular buffer of inputs.
// A zero value InputBuffer is ready to use.
type InputBuffer struct {
	// circular buffer
	buf [InputBufferSize]FrameInput
	// where to add the input next
	nextWrite int
	// index for the current input in the buffer.
	CurrentTick int
}

func (b *InputBuffer) Add(inputs ...Input) {
	for _, input := range inputs {
		b.buf[b.nextWrite] = FrameInput{
			Input: input,
		}
		b.CurrentTick = b.nextWrite
		b.nextWrite = (b.nextWrite + 1) % InputBufferSize
	}
}

// https://gamedev.stackexchange.com/a/68134
// The algorithm simply checks each input state from the
// CurrentTick offset inside the input buffer back until maxDuration.
func (b *InputBuffer) CheckSequence(maxDuration frame.Frame, sequence ...Input) bool {
	if len(sequence) == 0 {
		return true
	}
	w := len(sequence) - 1
	for i := range int(maxDuration) {
		frameInput := b.buf[(b.CurrentTick-i+InputBufferSize)%InputBufferSize]
		if frameInput.Input == sequence[w] {
			w--
		}
		if w == -1 {
			return true
		}
	}
	return false
}

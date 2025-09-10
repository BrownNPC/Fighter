package c

import (
	"GameFrameworkTM/components/frame"
)

const (
	// circular buffer size, how many inputs it can store.
	// 1 input per frame.
	InputBufferSize = 60
)

// InputBuffer is a circular buffer of inputs.
// A zero value InputBuffer is ready to use.
type InputBuffer struct {
	// circular buffer
	buf [InputBufferSize]Input
	// where to add the input next
	nextWrite int
	// index for the current input in the buffer.
	CurrentTick int
}

// Add adds an input to the circle buffer.
// the engine should only add 1 input per frame.
// Inputs can be combined using the | (OR) operator.
// eg. downForward := IDown|IForward
func (b *InputBuffer) Add(input Input) {
	b.buf[b.nextWrite] = input
	b.CurrentTick = b.nextWrite
	b.nextWrite = (b.nextWrite + 1) % InputBufferSize
}

// https://gamedev.stackexchange.com/a/68134
// The algorithm simply checks each input state from the
// CurrentTick offset inside the input buffer back until maxDuration.
// leniency is the number of frames the user has to perform this input
// and how many of those frames can be sloppy.
// eg. for the sequence [Down, Down|Forward, Forward] with leniency of 5 would allow for
// [Down, Up, Down|Forward, Back, Forward] to be matched, as long as they occured in the window of 5 frames.
func (b *InputBuffer) CheckSequence(leniency frame.Frame, sequence ...Input) bool {
	if len(sequence) == 0 {
		return true
	}
	w := len(sequence) - 1
	for i := range int(leniency) {
		// walk backwards from current frame for maxDuration
		frameInput := b.buf[Modulo(b.CurrentTick-i, InputBufferSize)]
		// if any of the inputs this frame matches with the sequence
		if frameInput == sequence[w] {
			w--
		}
		if w == -1 {
			return true
		}
	}
	return false
}

// Get the latest input
func (b *InputBuffer) Latest() Input {
	return b.buf[b.CurrentTick]
}

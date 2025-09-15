package input

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/frame"
)

const (
	// circular buffer size, how many inputs it can store.
	// 1 input per frame.
	InputBufferSize = 60 * 10
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
// CurrentTick offset inside the input buffer back until "limit" frames.
// NOTE:
// You must clean the move from the buffer BEFORE it's performed.
func (b *InputBuffer) CheckSequence(limit frame.Frame, sequence ...Input) bool {
	if len(sequence) == 0 {
		return true
	}
	w := len(sequence) - 1
	for i := range int(limit) {
		// walk backwards from current frame for maxDuration
		frameInput := b.buf[c.Modulo(b.CurrentTick-i, InputBufferSize)]
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

// Clear a move from the buffer. Use this to stop a move from being performed more than once.
// You must clean the move from the buffer BEFORE it's performed, and in the same tick as it's read.
func (b *InputBuffer) ClearMove(limit frame.Frame, sequence ...Input) {
	if len(sequence) == 0 {
		return
	}

	w := len(sequence) - 1
	for i := range limit {
		// walk backwards from current frame for maxDuration
		frameInput := &b.buf[c.Modulo(b.CurrentTick-i, InputBufferSize)]
		// if any of the inputs this frame matches with the sequence
		if *frameInput == sequence[w] {
			*frameInput = INeutral
			w--
		}
		if w == -1 {
			return
		}
	}
}

// Get the latest input
func (b *InputBuffer) Latest() Input {
	return b.buf[b.CurrentTick]
}

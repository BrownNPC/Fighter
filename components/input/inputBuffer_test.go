package input_test

import (
	"GameFrameworkTM/components/input"
	"testing"
)

func TestInputBuffer(t *testing.T) {
	var b input.InputBuffer
	haduken := input.NewMove(true, 12,
		input.Down, input.Down|input.Forward, input.Forward)
	for _, i := range haduken.Sequence {
		b.Add(i)
	}

	if !b.CheckSequence(haduken) {
		t.Error("failed to check haduken")
	}
}

package input_test

import (
	"GameFrameworkTM/components/input"
	"testing"
)

func TestInputBuffer(t *testing.T) {
	var b input.InputBuffer

	haduken := []input.Input{input.IDown, input.IDown | input.IForward, input.IForward}
	for _, i := range haduken {
		b.Add(i)
	}

	if !b.CheckSequence(12, haduken...) {
		t.Error("failed to check haduken")
	}
}

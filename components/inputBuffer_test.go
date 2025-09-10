package c_test

import (
	c "GameFrameworkTM/components"
	"testing"
)

func TestInputBuffer(t *testing.T) {
	var b c.InputBuffer

	haduken := []c.Input{c.IDown, c.IDown | c.IForward, c.IForward}
	b.Add(c.IDown, c.IDown|c.IForward, c.IForward)

	if !b.CheckSequence(12, haduken...) {
		t.Error("failed to check haduken")
	}
}

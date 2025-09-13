// package frame provides utilities for tracking frame times.
// It's basically the time package but for frames
package frame

// Frame is a number representing a point in time.
type Frame int

// the frame the game currently is on
var currentFrame Frame

// Now returns the current frame the game is on
func Now() Frame {
	return currentFrame
}

// Increment the internal frame counter. This must be called once each frame.
func Increment() {
	currentFrame++
}

// Since returns how many frames have passed Since the frame
func Since(f Frame) Frame {
	return currentFrame - f
}

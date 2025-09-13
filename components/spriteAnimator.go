package c

// SpriteAnimator provides a basic ticker that tells you what frame to draw in your animation.
// This version supports any animation FPS (including >60) and any number of frames.
type SpriteAnimator struct {
	currentFrame int
	TotalFrames  int
	// desired animation frames per second
	animFPS float64
	// accumulated fractional frames (ticks)
	accum float64
}

// NewSpriteAnimator provides a basic ticker that tells you what frame to draw in your animation.
// fps = desired animation frames-per-second (how many animation frames advance per second).
// totalFrames = number of frames in the animation.
func NewSpriteAnimator(fps, totalFrames int) SpriteAnimator {
	if fps <= 0 {
		fps = 1
	}
	a := SpriteAnimator{
		TotalFrames:  totalFrames,
		currentFrame: 0,
		animFPS:      float64(fps),
		accum:        0.0,
	}
	return a
}

// Reset returns to the first frame and clears internal accumulators.
func (a *SpriteAnimator) Reset() {
	a.currentFrame = 0
	a.accum = 0.0
}

// GetCurrentFrame updates the internal ticker and returns the current frame you should draw.
// Call this once per *game* frame (the game is assumed to run at 60 FPS).
func (a *SpriteAnimator) GetCurrentFrame() int {
	// Add the fraction of animation-frames produced this game tick.
	// Example: animFPS=30 -> add 30/60 = 0.5 per tick (so 2 ticks -> advance 1 frame).
	// If animFPS=120 -> add 120/60 = 2.0 per tick (advance 2 frames per tick).
	a.accum += a.animFPS / 60.0

	if a.accum >= 1.0 && a.TotalFrames > 0 {
		advance := int(a.accum) // how many animation frames to advance now
		a.accum -= float64(advance)     // remove the consumed whole frames
		a.currentFrame = (a.currentFrame + advance) % a.TotalFrames
	}

	return a.currentFrame
}

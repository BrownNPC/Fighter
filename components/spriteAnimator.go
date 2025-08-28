package c

type SpriteAnimator struct {
	currentFrame int
	TotalFrames  int
	Counter      int
	UpdateEvery  int // how many game frames to wait before advancing
}

// NewSpriteAnimator provides a basic ticker that tells you what frame to draw in your animation
// Note that max frames are 60.
func NewSpriteAnimator(fps int, totalFrames int) SpriteAnimator {
	return SpriteAnimator{
		TotalFrames:  totalFrames,
		currentFrame: 0,
		Counter:      0,
		// game is fixed at 60 fps
		UpdateEvery: 60 / fps,
	}
}

// Forcibly set currentFrame to first frame
func (a *SpriteAnimator) Reset() {
	a.currentFrame = 0
	a.Counter = 0
}

// GetCurrentFrame updates the internal ticker and returns the current frame you should draw.
// It must be called every frame.
func (a *SpriteAnimator) GetCurrentFrame() int {
	a.Counter++
	if a.Counter >= a.UpdateEvery {
		a.Counter = 0
		a.currentFrame = (a.currentFrame + 1) % a.TotalFrames
	}
	return a.currentFrame
}

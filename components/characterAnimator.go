package c

// AnimationType represent different animations for character states.
type AnimationType uint8

const (
	AnimationIdle AnimationType = iota
	AnimationWalk
	TotalAnimations
)

// CharacterAnimator contains different animations that can be played.
type CharacterAnimator struct {
	animations      [TotalAnimations]SpriteAnimator
	ActiveAnimation AnimationType
}

// Switch to an animation.
func (a *CharacterAnimator) Switch(to AnimationType) {
	a.animations[to].Reset()
	a.ActiveAnimation = to
}

// Draw the current frame of the animation
func (a *CharacterAnimator) Draw(x, y int32) {
	// TODO support drawing flipped by recieving a bool
}

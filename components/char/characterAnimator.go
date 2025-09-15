package char

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/render"
	"io/fs"
	"log/slog"
)

// AnimationType represent different animations for character states.
type AnimationType uint8

const (
	Idle AnimationType = iota
	Walk
	TotalAnimations
)

// CharacterAnimator contains different animations that can be played.
type CharacterAnimator struct {
	animations      [TotalAnimations]render.BaseSprite
	ActiveAnimation AnimationType
}
type AnimationConfig struct {
	// Prefix is the prefix for the animation spritesheet.
	// eg. idle for idle_5x5_23frames.png
	Prefix string
	Type   AnimationType
	FPS    int
}

func NewCharacterAnimator(characterName string, resolution c.Vec2, as fs.FS, supportedAnimations []AnimationConfig) CharacterAnimator {
	var ch CharacterAnimator
	for _, anim := range supportedAnimations {
		sprite, err := render.LoadCharacterAnimation(characterName, anim.Prefix, anim.FPS, resolution, as)
		if err != nil {
			slog.Error("Failed to load character animation, error occured while loading sprite", "character name", characterName, "animation prefix", anim.Prefix, "error", err)
			continue
		}
		ch.animations[anim.Type] = sprite
	}
	return ch
}

// Switch to an animation.
func (a *CharacterAnimator) Switch(to AnimationType) {
	a.animations[to].Reset()
	a.ActiveAnimation = to
}

// Draw the current frame of the animation
func (a *CharacterAnimator) Draw(x, y float32, facing Direction) {
	var anim render.BaseSprite = a.animations[a.ActiveAnimation]
	if facing == Left {
		anim.DrawFlipped(x, y)
	}else{
		anim.Draw(x,y)
	}
}

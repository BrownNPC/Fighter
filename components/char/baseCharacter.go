package char

import (
	"GameFrameworkTM/components/fixed"
	"GameFrameworkTM/engine"
)

type BaseCharacter struct {
	Position fixed.Vector2
	Velocity fixed.Vector2
	Anim     CharacterAnimator
	Facing   engine.Direction
}

// Initialize BaseCharacter
func (b *BaseCharacter) Init(Anim CharacterAnimator, Facing engine.Direction) {
	b.Anim = Anim
	b.Facing = Facing
}

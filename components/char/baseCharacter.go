package char

import (
	"GameFrameworkTM/components/fixed"
)

type BaseCharacter struct {
	Position fixed.Vector2
	Velocity fixed.Vector2
	Anim     CharacterAnimator
	Facing   Direction
}

// Initialize BaseCharacter
func (b *BaseCharacter) Init(Anim CharacterAnimator, Facing Direction) {
	b.Anim = Anim
	b.Facing = Facing
}

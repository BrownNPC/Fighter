package render

import (
	c "GameFrameworkTM/components"
	"io/fs"
	"path"
)

// usage : LoadCharacterAnimation("steve","walk")
func LoadCharacterAnimation(character, sprite string, fps int, resolution c.Vec2, as fs.FS) (BaseSprite, error) {
	characterPath := path.Join("assets", "characters", character)
	baseSprite, err := loadBaseSprite(as, characterPath, sprite, fps, resolution)
	return baseSprite, err
}

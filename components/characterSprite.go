package c

import (
	"io/fs"
	"path"
)

// usage : LoadCharacterSprite("steve","walk")
func LoadCharacterSprite(character, sprite string, fps int, resolution Vec2,as fs.FS) (BaseSprite, error) {
	characterPath := path.Join("assets", "characters", character)
	baseSprite, err := loadBaseSprite(as, characterPath, sprite, fps, resolution)
	return baseSprite, err
}

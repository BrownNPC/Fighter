package c

import rl "github.com/gen2brain/raylib-go/raylib"

type CharacterSprite struct {
	SpriteAnimator
	// sprite sheet containing all the frames for the animation
	Resource   rl.Texture2D
	SpritePath string `json:"SpritePath"`
	// number of frames in the sprite sheet
	TotalFrames int `json:"TotalFrames"`
}

package c

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var StageResolution = V2(512, 288)

type StageBackground struct {
	Frames         []rl.Texture2D
	SpriteAnimator SpriteAnimator
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, rd fs.FS) (StageBackground, error) {
	const stageBackgroundPattern = "background_%d.png"
	// paths of background frames
	var framePaths []string
	// find all files matching stageBackgroundPattern
	// stageBackgroundPattern must enumerate
	var lastFrame int = 1
	err := fs.WalkDir(rd, path.Join("assets", stageName),
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				panic(err)
			}
			if !d.IsDir() {
				ok, err := filepath.Match(fmt.Sprintf(stageBackgroundPattern, lastFrame), d.Name())
				if err != nil {
					panic(err)
				}
				if ok {
					lastFrame++
					framePaths = append(framePaths, path)
				}
			}
			return nil
		})
	if err != nil {
		return StageBackground{}, err
	}
	if len(framePaths) == 0 {
		return StageBackground{}, fmt.Errorf("No stage frames found for stage assets/%s. Make sure to have at least 1 file with the name background_1.png", stageName)
	}
	frames := make([]rl.Texture2D, len(framePaths))

	for i, tex := range framePaths {
		frames[i] = rl.LoadTexture(tex)
	}
	return StageBackground{
		Frames:         frames,
		SpriteAnimator: NewSpriteAnimator(len(frames), len(frames)),
	}, nil
}
func (s *StageBackground) Draw() {
	currentFrame := s.SpriteAnimator.GetCurrentFrame()
	frame := s.Frames[currentFrame]

	// we dont care what size the frame is. Just draw it as Stage Resolution.
	// Strech it, scale it down, whatever.
	rl.DrawTexturePro(frame,
		rl.NewRectangle(0, 0, float32(frame.Width), float32(frame.Height)),
		rl.NewRectangle(0, 0, StageResolution.X, StageResolution.Y),
		V2Z.R(),
		0,
		rl.White,
	)
}
// Unload frees gpu resources
func (s *StageBackground) Unload() {
	for _, frame := range s.Frames {
		if rl.IsTextureValid(frame) {
			rl.UnloadTexture(frame)
		}
	}
}

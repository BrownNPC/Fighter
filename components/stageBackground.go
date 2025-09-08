package c

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StageBackground struct {
	// single texture containing frames, horizontally.
	Frames         rl.Texture2D
	SpriteAnimator SpriteAnimator
	Resolution     Vec2
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, resolution Vec2, as fs.FS) (StageBackground, error) {
	stagePath := path.Join("assets", "stages", stageName)
	entries, err := os.ReadDir(stagePath)
	if err != nil {
		return StageBackground{}, err
	}
	var (
		// frame_192.png
		frameName string
		//192
		numFrames int
	)

	re := regexp.MustCompile(`^frames_(\d+)\.png$`)
	for _, entry := range entries {
		if !entry.IsDir() {
			if m := re.FindStringSubmatch(entry.Name()); m != nil {
				frameName = entry.Name()
				numFrames, err = strconv.Atoi(m[1])
				if err != nil {
					return StageBackground{}, fmt.Errorf("failed to convert string to int %w %s", err, frameName)
				}
			}
		}
	}
	if frameName == "" {
		return StageBackground{}, fmt.Errorf("no frames found. make sure there is a frames_<numberOfFrames>.png in %s for example frames_12.png that contains all the frames for the background", stagePath)
	}
	return StageBackground{
		Frames:         rl.LoadTexture(filepath.Join(stagePath, frameName)),
		SpriteAnimator: NewSpriteAnimator(60, numFrames),
		Resolution:     resolution,
	}, nil
}
func (s *StageBackground) Draw() {
	currentFrame := s.SpriteAnimator.GetCurrentFrame()
	srcX := float32(currentFrame) * s.Resolution.X
	// we dont care what size the frame is. Just draw it as Stage Resolution.
	// Strech it, scale it down, whatever.
	rl.DrawTexturePro(s.Frames,
		rl.NewRectangle(srcX, 0, s.Resolution.X, s.Resolution.Y),
		rl.NewRectangle(0, 0, s.Resolution.X, s.Resolution.Y),
		V2Z.R(),
		0,
		rl.White,
	)
}

// Unload frees gpu resources
func (s *StageBackground) Unload() {
	rl.UnloadTexture(s.Frames)
}

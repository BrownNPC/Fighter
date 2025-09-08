package c

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StageBackground struct {
	// single texture containing frames, horizontally.
	Frames        rl.Texture2D
	Rows, Columns int

	SpriteAnimator SpriteAnimator
	Resolution     Vec2
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, resolution Vec2, as fs.FS) (StageBackground, error) {
	stagePath := path.Join("assets", "stages", stageName)
	entries, err := fs.ReadDir(as, stagePath)
	if err != nil {
		return StageBackground{}, err
	}
	var (
		frameName string
		rows      int
		columns   int
		//192
		numFrames int
	)

	re := regexp.MustCompile(`^stage_(\d+)x(\d+)_(\d+)frames\.png$`)
	for _, entry := range entries {
		if !entry.IsDir() {
			if m := re.FindStringSubmatch(entry.Name()); m != nil {
				frameName = entry.Name()
				rows, _ = strconv.Atoi(m[1])
				columns, _ = strconv.Atoi(m[2])
				numFrames, _ = strconv.Atoi(m[3])
			}
		}
	}
	if frameName == "" {
		return StageBackground{}, fmt.Errorf("no frames found. make sure there is a frames_<numberOfFrames>.png in %s for example frames_12.png that contains all the frames for the background", stagePath)
	}
	return StageBackground{
		Frames: rl.LoadTexture(filepath.Join(stagePath, frameName)),
		Rows:   rows, Columns: columns,
		SpriteAnimator: NewSpriteAnimator(30, numFrames),
		Resolution:     resolution,
	}, nil
}
func (s *StageBackground) Draw(x, y float32) {
	currentFrame := s.SpriteAnimator.GetCurrentFrame()

	column := currentFrame % s.Columns
	row := currentFrame / s.Columns

	srcX := float32(column) * s.Resolution.X
	srcY := float32(row) * s.Resolution.Y

	// we dont care what size the frame is. Just draw it as Stage Resolution.
	// Strech it, scale it down, whatever.
	rl.DrawTexturePro(s.Frames,
		rl.NewRectangle(srcX, srcY, s.Resolution.X, s.Resolution.Y),
		rl.NewRectangle(x, y, s.Resolution.X, s.Resolution.Y),
		V2Z.R(),
		0,
		rl.White,
	)
}

// Unload frees gpu resources
func (s *StageBackground) Unload() {
	rl.UnloadTexture(s.Frames)
}

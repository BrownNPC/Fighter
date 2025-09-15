package render

import (
	c "GameFrameworkTM/components"
	"fmt"
	"io/fs"
	"math"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StageBackground struct {
	BaseSprite
	VirtualSize c.Vec2
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, resolution c.Vec2, virtualSize c.Vec2, as fs.FS) (StageBackground, error) {
	stagePath := path.Join("assets", "stages", stageName)
	baseSprite, err := loadBaseSprite(as, stagePath, "stage", 80, resolution)
	if err != nil {
		return StageBackground{}, fmt.Errorf("failed to load stage: %w", err)
	}
	return StageBackground{
		BaseSprite:  baseSprite,
		VirtualSize: virtualSize,
	}, nil
}

// Get the frame to draw for the camera virtual X coordinate.
func (s *StageBackground) GetFrameForCameraX(camX float32) int {
	if s.VirtualSize.X <= 0 {
		panic("Virtual size cannot be less than 0")
	}
	// middle / center frame
	mid := (s.TotalFrames - 1) / 2
	// camera starts from the middle of the screen
	start := s.VirtualSize.X / 2

	// offset in pixels from the start center +1 px maps to +1 frame
	offset := int(math.Round(float64(camX - start)))

	// clamp
	idx := max(mid+offset, 0)
	if idx >= s.TotalFrames {
		idx = s.TotalFrames - 1
	}

	return idx
}
func (s *StageBackground) DrawFrame(f int, x, y float32) {
	srcRec := s.BaseSprite.GetRectForFrame(f)
	rl.DrawTexturePro(s.Resource,
		srcRec,
		rl.NewRectangle(x, y, s.Resolution.X, s.Resolution.Y),
		c.V2Z.R(),
		0,
		rl.White,
	)
}
func (s *StageBackground) Draw(x, y float32) {
	s.DrawFrame(s.TotalFrames/2, x, y)
}

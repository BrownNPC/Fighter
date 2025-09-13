package c

import (
	"fmt"
	"io/fs"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StageBackground struct {
	BaseSprite
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, resolution Vec2, as fs.FS) (StageBackground, error) {
	stagePath := path.Join("assets", "stages", stageName)
	baseSprite, err := loadBaseSprite(as, stagePath, "stage", 80, resolution)
	if err != nil {
		return StageBackground{}, fmt.Errorf("failed to load stage: %w", err)
	}
	return StageBackground{
		BaseSprite: baseSprite,
	}, nil
}

func (s *StageBackground) Draw(x, y float32) {
	srcRec := s.BaseSprite.GetRectForFrame((s.TotalFrames / 2))
	rl.DrawTexturePro(s.Resource,
		srcRec,
		rl.NewRectangle(x, y, s.Resolution.X, s.Resolution.Y),
		V2Z.R(),
		0,
		rl.White,
	)

}

package c

import (
	"fmt"
	"io/fs"
	"path"
)

type StageBackground struct {
	BaseSprite
}

// The filesystem is only used for reading how many frames there are for a stage.
func LoadStage(stageName string, resolution Vec2, as fs.FS) (StageBackground, error) {
	stagePath := path.Join("assets", "stages", stageName)
	baseSprite, err := loadBaseSprite(as, stagePath, "stage", 40, resolution)
	if err != nil {
		return StageBackground{}, fmt.Errorf("failed to load stage: %w", err)
	}
	return StageBackground{
		BaseSprite: baseSprite,
	}, nil
}

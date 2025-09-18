package start

import (
	"GameFrameworkTM/engine"
)

type Scene struct {
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx *engine.Context) {
}

// update is called every frame
func (scene *Scene) Update(ctx *engine.Context) (unload bool) {
	return true // if true is returned, Unload is called
}

// called after Update returns true
func (scene *Scene) Unload(ctx *engine.Context) (nextSceneID string) {
	return "play" // the engine will switch to the scene that is registered with this id
}

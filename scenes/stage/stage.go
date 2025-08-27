package stage

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	// draw to render texture
	Screen c.RenderTexture
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.Screen = c.NewRenderTexture(c.RecommendedVirtualResolution)
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "start" // the engine will switch to the scene that is registered with this id
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	scene.Screen.BeginDrawing()
	rl.ClearBackground(rl.White)
	scene.Screen.EndDrawing()
	return false // if true is returned, Unload is called
}

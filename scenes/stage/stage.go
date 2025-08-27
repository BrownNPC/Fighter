package stage

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	// draw to render texture
	Screen       c.RenderTexture
	Background   rl.Texture2D
	DrawTemplate bool
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.Background = rl.LoadTexture("assets/stage1/background.png")
	scene.Screen = c.NewRenderTexture(c.RecommendedVirtualResolution)
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	scene.Screen.Unload()
	return "start" // the engine will switch to the scene that is registered with this id
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	if rl.IsKeyPressed(rl.KeyF3) {
		scene.DrawTemplate = !scene.DrawTemplate
	}
	scene.Screen.BeginDrawing()
	rl.DrawTexture(scene.Background, 0, 0, rl.White)
	scene.Screen.EndDrawing()
	return false // if true is returned, Unload is called
}

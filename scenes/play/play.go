package play

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"
	"log"
)

type Scene struct {
	// draw to render texture
	Screen c.Screen
	Stage  c.StageBackground

	// slice of unloader functions
	Unloader c.Stack[func()]
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	var err error
	scene.Screen = c.NewRenderTexture(ctx.VirtualResolution)
	defer scene.Unloader.Add(scene.Screen.Unload)

	scene.Stage, err = c.LoadStage("stage1", ctx.StageResolution, ctx.Assets)
	if err != nil {
		log.Fatalln("failed to load stage", err)
	}
	defer scene.Unloader.Add(scene.Stage.Unload)
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	for _, unloadFunc := range scene.Unloader.Items {
		unloadFunc()
	}
	return "start" // the engine will switch to the scene that is registered with this id
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	scene.Screen.BeginDrawing()
	scene.Stage.Draw()
	scene.Screen.EndDrawing()
	return false // if true is returned, Unload is called
}

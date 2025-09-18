package play

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/input"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	// draw to render texture
	Screen       render.Screen
	Stage        render.StageBackground
	Steve        render.BaseSprite
	Shadow       rl.Texture2D
	currentFrame int
	cam          c.Vec2
	// slice of unloader functions
	Unloader c.Stack[func()]
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx *engine.Context) {
	var err error
	scene.Screen = render.NewScreen(ctx.VirtualResolution)
	defer scene.Unloader.Add(scene.Screen.Unload)

	scene.Stage, err = render.LoadStage("stage1", ctx.StageFrameResolution, ctx.VirtualStageSize, ctx.Assets)
	if err != nil {
		log.Fatalln("failed to load stage", err)
	}
	// center
	scene.cam.X = scene.Stage.VirtualSize.X / 2

	scene.Shadow = rl.LoadTexture("assets/misc/shadow.png")
	scene.Steve, err = render.LoadCharacterAnimation("steve", "idle", 11, c.V2(256, 256), ctx.Assets)
	scene.currentFrame = 29
	defer scene.Unloader.Add(scene.Stage.Unload)
}

// called after Update returns true
func (scene *Scene) Unload(ctx *engine.Context) (nextSceneID string) {
	for _, unloadFunc := range scene.Unloader.Items {
		unloadFunc()
	}
	return "start" // the engine will switch to the scene that is registered with this id
}

// update is called every frame
func (scene *Scene) Update(ctx *engine.Context) (unload bool) {
	scene.Screen.BeginDrawing()
	rl.ClearBackground(rl.White)
	if rl.IsKeyDown(rl.KeyD) {
		scene.cam.X++
	}
	if rl.IsKeyDown(rl.KeyA) {
		scene.cam.X--
	}
	haduken := input.MoveGroup{
		input.NewMove(true, 12, input.Down, input.Down|input.Forward, input.Forward, input.Attack),
		input.NewMove(true, 10, input.Down, input.Down|input.Forward, input.Forward|input.Attack),
	}
	if haduken.Check(&engine.Player2) {
		fmt.Println("HADUKEN")
	}
	scene.cam.X = min(ctx.VirtualStageSize.X, scene.cam.X)
	scene.cam.X = max(0, scene.cam.X)
	scene.currentFrame = scene.Stage.GetFrameForCameraX(scene.cam.X)

	scene.Stage.DrawFrame(scene.currentFrame, 0, -600+576)
	scene.Steve.Draw(0, 97)
	scene.Screen.EndDrawing()
	rl.DrawText(fmt.Sprint(scene.currentFrame), 0, 0, 22, rl.White)
	return false // if true is returned, Unload is called
}

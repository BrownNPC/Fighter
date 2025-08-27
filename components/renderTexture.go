package c

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// RenderTexture provides a wrapper around a raylib render texture
type RenderTexture struct {
	screen        rl.RenderTexture2D
	Width, Height int
}

// retro resolution
var RecommendedVirtualResolution = V2(320, 240)

// Recommended 320×240 resolution for retro feel
func NewRenderTexture(resolution Vec2) RenderTexture {
	virtualWidth, virtualHeight := resolution.ToInt()

	r := RenderTexture{
		screen: rl.LoadRenderTexture(int32(virtualWidth), int32(virtualHeight)),
		Width:  virtualWidth,
		Height: virtualHeight,
	}
	return r
}
func (r *RenderTexture) Unload() {
	rl.UnloadRenderTexture(r.screen)
}

// render the render texture with black bars to keep the aspect ratio
func (r *RenderTexture) Render() {
	target := r.screen
	scale := r.Scale()
	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.Rectangle{
			X:      (float32(rl.GetScreenWidth()) - float32(r.Width)*scale) * 0.5,
			Y:      (float32(rl.GetScreenHeight()) - float32(r.Height)*scale) * 0.5,
			Width:  float32(r.Width) * scale,
			Height: float32(r.Height) * scale,
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}

// render the render texture with black bars to keep the aspect ratio
func (r *RenderTexture) BeginDrawing() {
	rl.BeginTextureMode(r.screen)
}

// render the render texture with black bars to keep the aspect ratio
func (r *RenderTexture) EndDrawing() {
	rl.EndTextureMode()
	r.Render()
}
func (r *RenderTexture) Scale() float32 {
	return min(float32(rl.GetScreenWidth())/float32(r.Width),
		float32(rl.GetScreenHeight())/float32(r.Height))
}

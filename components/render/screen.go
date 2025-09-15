package render

import (
	c "GameFrameworkTM/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Screen provides a wrapper around a raylib render texture
type Screen struct {
	renderTexture rl.RenderTexture2D
	Width, Height int
}

// Does letterboxing (black bars) to maintain the resolution.
func NewScreen(resolution c.Vec2) Screen {
	virtualWidth, virtualHeight := resolution.ToInt()
	r := Screen{
		renderTexture: rl.LoadRenderTexture(int32(virtualWidth), int32(virtualHeight)),
		Width:         virtualWidth,
		Height:        virtualHeight,
	}
	return r
}
func (r *Screen) Unload() {
	rl.UnloadRenderTexture(r.renderTexture)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) Render() {
	target := r.renderTexture
	scale := r.Scale()
	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.Rectangle{
			X:      (float32(rl.GetRenderWidth()) - float32(r.Width)*scale) * 0.5,
			Y:      (float32(rl.GetRenderHeight()) - float32(r.Height)*scale) * 0.5,
			Width:  float32(r.Width) * scale,
			Height: float32(r.Height) * scale,
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) BeginDrawing() {
	rl.BeginTextureMode(r.renderTexture)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) EndDrawing() {
	rl.EndTextureMode()
	r.Render()
}
func (r *Screen) Scale() float32 {
	return min(float32(rl.GetRenderWidth())/float32(r.Width),
		float32(rl.GetRenderHeight())/float32(r.Height))
}

//go:build js

package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed assets
var Assets embed.FS

func init() {
	ASSETS = Assets
	rl.AddFileSystem(Assets)
}

package main

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes"
	"fmt"
	"io/fs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ASSETS fs.FS

// You can edit the window title in this file.
func main() {
	rl.SetTraceLogLevel(rl.LogDebug)
	err := engine.Run(scenes.Registered, engine.Config{
		WindowTitle:       "Fighter",
		Assets:            ASSETS,
		VirtualResolution: c.V2(384, 288),
		StageResolution:   c.V2(1024, 288),
	})
	if err != nil {
		fmt.Println(err)
	}
}

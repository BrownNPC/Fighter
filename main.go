package main

import (
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes"
	"fmt"
	"io/fs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ASSETS fs.FS

// You can edit the window title in this file.
func main() {
	rl.SetTraceLogLevel(rl.LogError)
	err := engine.Run(scenes.Registered, engine.Config{
		WindowTitle: "Fighter",
		Assets:      ASSETS,
	})
	if err != nil {
		fmt.Println(err)
	}
}

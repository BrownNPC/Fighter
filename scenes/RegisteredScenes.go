package scenes

import (
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes/stage"
	"GameFrameworkTM/scenes/start"
)

// register all your scenes in here
var Registered = engine.Scenes{
	"start": &start.Scene{},
	"stage": &stage.Scene{},
}

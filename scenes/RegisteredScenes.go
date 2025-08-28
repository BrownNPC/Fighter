package scenes

import (
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes/play"
	"GameFrameworkTM/scenes/start"
)

// register all your scenes in here
var Registered = engine.Scenes{
	"start": &start.Scene{},
	"play": &play.Scene{},
}

//go:build js

package main

import "embed"

//go:embed assets
var Assets embed.FS

func init() {
	ASSETS = Assets
}

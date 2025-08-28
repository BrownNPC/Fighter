//go:build !js
package main

import "os"

func init() {
	ASSETS = os.DirFS(".")
}

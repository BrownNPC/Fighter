package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	os.Mkdir("build/cache", 0o755)
	var gitRepo = filepath.FromSlash("build/cache/Raylib-Go-Wasm")
	var distPath = filepath.FromSlash("build/dist/web")

	// clone git repo
	if !DirExists(gitRepo) {
		cmd := exec.Command("git", "clone", "--depth", "1",
			"https://github.com/BrownNPC/Raylib-Go-Wasm.git",
			gitRepo)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	// set env for wasm build
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	// call the Go compiler.
	// make sure to add a replace directive to use the web build.
	WithReplace(
		"github.com/gen2brain/raylib-go/raylib",
		"./build/cache/Raylib-Go-Wasm/raylib",
		func() error {
			// call the Go compiler
			buildOutput := filepath.Join(gitRepo, "index", "main.wasm")
			cmd := exec.Command("go", "build", "-o", buildOutput, ".")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	)

	os.MkdirAll(distPath, 0o755)
	wasmRuntimePath := filepath.Join(build.Default.GOROOT, "lib/wasm/wasm_exec.js")
	// copy wasm runtime
	if err := CopyToDir(wasmRuntimePath, distPath); err != nil {
		log.Panicln("failed to copy wasm runtime", err)
	}

	if err := CopyToDir(filepath.Join(gitRepo, "index"), distPath); err != nil {
		log.Panicln("failed to copy wasm runtime", err)
	}
	fmt.Println("build artifacts are in:", distPath)
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

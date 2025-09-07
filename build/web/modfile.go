package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/mod/modfile"
)

// WithReplace adds the replace directive, runs action(), and then restores the original go.mod.
// oldVersion/newVersion can be empty strings when not applicable.
func WithReplace(oldPath, newPath string, action func() error) error {
	// Read original go.mod
	orig, err := os.ReadFile("go.mod")
	if err != nil {
		return fmt.Errorf("read go.mod: %w", err)
	}

	// Write a backup so we can always restore exactly
	backup := "go.mod.bak"
	if err := os.WriteFile(backup, orig, 0644); err != nil {
		return fmt.Errorf("write backup: %w", err)
	}

	// Ensure we restore the original go.mod no matter what.
	defer func() {
		_ = os.WriteFile("go.mod", orig, 0644)
		_ = os.Remove(backup)
		cmd := exec.Command("go", "mod", "tidy")
		cmd.Run()
	}()

	// Parse and add the replace
	mf, err := modfile.Parse("go.mod", orig, nil)
	if err != nil {
		return fmt.Errorf("parse go.mod: %w", err)
	}

	if err := mf.AddReplace(oldPath, "", newPath, ""); err != nil {
		return fmt.Errorf("add replace: %w", err)
	}

	newBytes, err := mf.Format()
	if err != nil {
		return fmt.Errorf("format go.mod: %w", err)
	}

	if err := os.WriteFile("go.mod", newBytes, 0644); err != nil {
		return fmt.Errorf("write modified go.mod: %w", err)
	}
	exec.Command("go", "mod", "tidy").Run()

	// Run the user action while the replace is present
	if err := action(); err != nil {
		return fmt.Errorf("action failed: %w", err)
	}

	// success; defer will restore original go.mod
	return nil
}

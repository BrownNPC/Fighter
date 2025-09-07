package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyToDir copies either a file or a directory's contents into dstDir.
// - If src is a file, it copies the file to dstDir/<basename(src)>.
// - If src is a directory, it copies the directory's contents (recursively)
//   into dstDir (i.e. children of src go directly under dstDir).
func CopyToDir(src, dstDir string) error {
	src = filepath.Clean(src)
	dstDir = filepath.Clean(dstDir)

	info, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Ensure dstDir exists
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return err
	}

	if info.IsDir() {
		return copyDirContents(src, dstDir)
	}

	// src is a file or symlink
	destPath := filepath.Join(dstDir, filepath.Base(src))
	return copyFileOrSymlink(src, destPath)
}

func copyDirContents(srcDir, dstDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dstDir, rel)

		info, err := d.Info()
		if err != nil {
			return err
		}

		// If directory -> create it (preserve perms)
		if d.IsDir() {
			// Create directory with same perms (umask may reduce them)
			if err := os.MkdirAll(target, info.Mode().Perm()); err != nil {
				return err
			}
			return nil
		}

		// File or symlink
		if info.Mode()&os.ModeSymlink != 0 {
			// Read symlink target and recreate symlink
			linkDest, err := os.Readlink(path)
			if err != nil {
				return err
			}
			// remove any existing target first
			_ = os.Remove(target)
			if err := os.Symlink(linkDest, target); err != nil {
				return err
			}
			return nil
		}

		// Regular file: ensure parent dir exists
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		return copyFileWithInfo(path, target, info)
	})
}

func copyFileOrSymlink(src, dst string) error {
	info, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// If symlink, recreate
	if info.Mode()&os.ModeSymlink != 0 {
		linkDest, err := os.Readlink(src)
		if err != nil {
			return err
		}
		_ = os.Remove(dst)
		return os.Symlink(linkDest, dst)
	}

	// Regular file
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	return copyFileWithInfo(src, dst, info)
}

func copyFileWithInfo(src, dst string, info os.FileInfo) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Create/truncate destination with same permission bits
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
	if err != nil {
		// On some systems Create may succeed but with different perms; try Create then Chmod
		out, err = os.Create(dst)
		if err != nil {
			return err
		}
		if err := out.Chmod(info.Mode().Perm()); err != nil {
			// not fatal; continue
		}
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	// Ensure final file mode is set
	if err := out.Chmod(info.Mode().Perm()); err != nil {
		// non-fatal on some platforms
	}

	// Preserve modification & access times (best-effort)
	modTime := info.ModTime()
	if err := os.Chtimes(dst, modTime, modTime); err != nil {
		// ignore: not all platforms support it
	}

	return nil
}


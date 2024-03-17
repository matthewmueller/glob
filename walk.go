package glob

import (
	"io/fs"
	"path/filepath"
)

// Walk the files
func Walk(pattern string, fn fs.WalkDirFunc) error {
	// Expand the pattern
	patterns, err := expand(pattern)
	if err != nil {
		return err
	}
	for _, pattern := range patterns {
		// Get the base directory (non-glob part of the pattern)
		dir := base(pattern)
		// Compile the pattern into a matcher
		matcher, err := compile(pattern)
		if err != nil {
			return err
		}
		// Walk the directory, matching files
		if err := filepath.WalkDir(dir, func(path string, de fs.DirEntry, err error) error {
			if !matcher.Match(path) {
				return nil
			}
			return fn(path, de, err)
		}); err != nil {
			return err
		}
	}
	return nil
}

// WalkFS walks an abstract abstract filesystem, matching on the pattern
func WalkFS(fsys fs.FS, pattern string, fn fs.WalkDirFunc) error {
	// Expand the pattern
	patterns, err := expand(pattern)
	if err != nil {
		return err
	}
	for _, pattern := range patterns {
		// Get the base directory (non-glob part of the pattern)
		dir := base(pattern)
		// Compile the pattern into a matcher
		matcher, err := compile(pattern)
		if err != nil {
			return err
		}
		// Walk the directory, matching files
		if err := fs.WalkDir(fsys, dir, func(path string, de fs.DirEntry, err error) error {
			if !matcher.Match(path) {
				return nil
			}
			return fn(path, de, err)
		}); err != nil {
			return err
		}
	}
	return nil
}

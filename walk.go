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
	seen := make(map[string]struct{})
	for _, pattern := range patterns {
		// Get the base directory (non-glob part of the pattern)
		dir := Base(pattern)
		// Compile the pattern into a matcher
		matcher, err := compile(pattern, filepath.Separator)
		if err != nil {
			return err
		}
		// Walk the directory, matching files
		if err := filepath.WalkDir(dir, func(path string, de fs.DirEntry, err error) error {
			if !matcher.Match(path) {
				return nil
			}
			if _, ok := seen[path]; ok {
				return nil
			}
			err = fn(path, de, err)
			seen[path] = struct{}{}
			return err
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
	seen := make(map[string]struct{})
	for _, pattern := range patterns {
		// Get the base directory (non-glob part of the pattern)
		dir := Base(pattern)
		// Compile the pattern into a matcher
		matcher, err := compile(pattern, '/')
		if err != nil {
			return err
		}
		// Walk the directory, matching files
		if err := fs.WalkDir(fsys, dir, func(path string, de fs.DirEntry, err error) error {
			if !matcher.Match(path) {
				return nil
			}
			if _, ok := seen[path]; ok {
				return nil
			}
			err = fn(path, de, err)
			seen[path] = struct{}{}
			return err
		}); err != nil {
			return err
		}
	}
	return nil
}

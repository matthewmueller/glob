package glob

import (
	"io/fs"
	"path/filepath"
)

// Walk the files
func Walk(glob string, fn fs.WalkDirFunc) error {
	matcher, err := Compile(glob)
	if err != nil {
		return err
	}

	// get the base directory (non-magic part)
	base := Base(glob)

	// start walking, matching on the glob
	return filepath.WalkDir(base, func(path string, de fs.DirEntry, err error) error {
		if !matcher.Match(path) {
			return nil
		}
		return fn(path, de, err)
	})
}

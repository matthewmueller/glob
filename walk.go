package glob

import (
	"os"

	"github.com/matthewmueller/go-glob/internal/fastwalk"
)

// WalkFunc is the walk function for glob.Walk
type WalkFunc func(path string, mode os.FileMode) error

// Walk the files
func Walk(glob string, fn WalkFunc) error {
	matcher, err := Compile(glob)
	if err != nil {
		return err
	}

	// get the base directory (non-magic part)
	base := Base(glob)

	// start walking, matching on the glob
	return fastwalk.Walk(base, func(path string, mode os.FileMode) error {
		if !matcher.Match(path) {
			return nil
		}
		return fn(path, mode)
	})
}

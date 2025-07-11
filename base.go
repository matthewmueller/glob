package glob

import (
	"path/filepath"

	"github.com/matthewmueller/glob/internal"
)

// Base gets the non-magical part of the glob
func Base(patterns ...string) string {
	return internal.Base(filepath.Separator, patterns...)
}

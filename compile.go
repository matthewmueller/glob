package glob

import (
	"path/filepath"

	"github.com/matthewmueller/glob/internal"
)

type Matcher = internal.Matcher

// Comple a glob pattern into a matcher
func Compile(patterns ...string) (Matcher, error) {
	return internal.Compile(filepath.Separator, patterns...)
}

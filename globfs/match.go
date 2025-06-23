package globfs

import (
	"path/filepath"

	"github.com/matthewmueller/glob/internal"
)

// Match checks if the given name matches any of the provided glob patterns.
// Note: We've reversed the order of parameters from filepath.Match to allow
// for multiple patterns to be passed in at once.
func Match(name string, patterns ...string) (bool, error) {
	return internal.Match(filepath.Separator, name, patterns...)
}

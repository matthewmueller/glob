package globfs

import (
	"github.com/matthewmueller/glob/internal"
)

// Base gets the non-magical part of the glob
func Base(patterns ...string) string {
	return internal.Base('/', patterns...)
}

package globfs

import "github.com/matthewmueller/glob/internal"

func Is(pattern string) bool {
	return internal.Is(pattern)
}

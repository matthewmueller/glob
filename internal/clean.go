package internal

import (
	"path"
	"path/filepath"
)

func clean(sep rune, p string) string {
	if sep == '/' {
		return path.Clean(p)
	}
	return filepath.Clean(p)
}

package globfs

import "io/fs"

// Glob returns the union of files matching the patterns in fsys. Patterns use
// slash-separated paths relative to the filesystem root. Empty patterns are
// ignored, and no patterns returns no matches.
func Glob(fsys fs.FS, patterns ...string) (matches []string, err error) {
	for path, err := range Walk(fsys, patterns...) {
		if err != nil {
			return nil, err
		}
		matches = append(matches, path)
	}
	return matches, nil
}

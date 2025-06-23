package globfs

import "io/fs"

// Glob matches a pattern against an abstract filesystem
func Glob(fsys fs.FS, patterns ...string) (matches []string, err error) {
	for err, path := range Walk(fsys, patterns...) {
		if err != nil {
			return nil, err
		}
		matches = append(matches, path)
	}
	return matches, nil
}

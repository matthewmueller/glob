package globfs

import (
	"io/fs"
	"iter"

	"github.com/matthewmueller/glob/internal"
)

// Walk an abstract abstract filesystem, matching on the pattern
func Walk(fsys fs.FS, patterns ...string) iter.Seq2[error, string] {
	return func(yield func(error, string) bool) {
		var all []string
		for _, pattern := range patterns {
			subpatterns, err := internal.Split(pattern)
			if err != nil {
				if !yield(err, "") {
					return
				}
				continue
			}
			all = append(all, subpatterns...)
		}

		// Compile the patterns into a single matcher
		matcher, err := Compile(all...)
		if err != nil {
			yield(err, "")
			return
		}

		// Compute the greatest common ancestor to walk from
		baseDir := Base(all...)

		if err := fs.WalkDir(fsys, baseDir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				if !yield(err, "") {
					return err
				}
				return nil
			}
			if !matcher.Match(path) {
				return nil
			}
			if !yield(nil, path) {
				return nil
			}
			return nil
		}); err != nil {
			if !yield(err, "") {
				return
			}
		}
	}
}

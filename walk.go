package glob

import (
	"io/fs"
	"iter"
	"path/filepath"

	"github.com/matthewmueller/glob/internal"
)

// Walk a directory, matching on the patterns
func Walk(dir string, patterns ...string) iter.Seq2[error, string] {
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
			for _, subpattern := range subpatterns {
				all = append(all, filepath.Join(dir, subpattern))
			}
		}

		// Compile the patterns into a single matcher
		matcher, err := Compile(all...)
		if err != nil {
			yield(err, "")
			return
		}

		// Compute the greatest common ancestor to walk from
		baseDir := Base(all...)

		if err := filepath.WalkDir(baseDir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				if !yield(err, "") {
					return fs.SkipAll
				}
				return nil
			}
			if !matcher.Match(path) {
				return nil
			}
			if !yield(nil, path) {
				return fs.SkipAll
			}
			return nil
		}); err != nil {
			if !yield(err, "") {
				return
			}
		}
	}
}

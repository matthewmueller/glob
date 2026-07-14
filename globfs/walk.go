package globfs

import (
	"io/fs"
	"iter"
	"path"

	"github.com/matthewmueller/glob/internal"
)

// Walk yields each matching path followed by a nil error. Errors are yielded
// with an empty path. Patterns use slash-separated paths relative to the
// filesystem root. Empty patterns are ignored, and no patterns yields nothing.
func Walk(fsys fs.FS, patterns ...string) iter.Seq2[string, error] {
	patterns = append([]string(nil), patterns...)
	return func(yield func(string, error) bool) {
		var all []string
		for _, pattern := range patterns {
			if pattern == "" {
				continue
			}
			subpatterns, err := internal.Split(pattern)
			if err != nil {
				if !yield("", err) {
					return
				}
				continue
			}
			for _, subpattern := range subpatterns {
				all = append(all, path.Clean(subpattern))
			}
		}
		if len(all) == 0 {
			return
		}

		// Compile the patterns into a single matcher
		matcher, err := Compile(all...)
		if err != nil {
			yield("", err)
			return
		}

		// Compute the greatest common ancestor to walk from
		baseDir := Base(all...)

		if err := fs.WalkDir(fsys, baseDir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				if !yield("", err) {
					return fs.SkipAll
				}
				return nil
			}
			if !matcher.Match(path) {
				return nil
			}
			if !yield(path, nil) {
				return fs.SkipAll
			}
			return nil
		}); err != nil {
			if !yield("", err) {
				return
			}
		}
	}
}

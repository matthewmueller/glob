package glob

import (
	"io/fs"
)

// Match is conceptually similar to filepath.Glob, but it uses the glob package
// from github.com/gobwas/glob, which supports more advanced globbing patterns.
func Match(glob string) (matches []string, err error) {
	err = Walk(glob, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if de.IsDir() {
			return nil
		}
		matches = append(matches, path)
		return nil
	})
	return matches, err
}

// MatchFS matches a pattern against an abstract filesystem
func MatchFS(fsys fs.FS, glob string) (matches []string, err error) {
	err = WalkFS(fsys, glob, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if de.IsDir() {
			return nil
		}
		matches = append(matches, path)
		return nil
	})
	return matches, err
}

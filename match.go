package glob

import "io/fs"

// Match is conceptually similar to filepath.Glob, but it uses the glob package
// from github.com/gobwas/glob, which supports more advanced globbing patterns.
func Match(pattern string) (matches []string, err error) {
	err = Walk(pattern, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		matches = append(matches, path)
		return nil
	})
	return matches, err
}

package glob

// Glob is conceptually similar to filepath.Glob, but it uses the glob package
// from github.com/gobwas/glob, which supports more advanced globbing patterns.
func Glob(dir string, patterns ...string) (matches []string, err error) {
	for err, path := range Walk(dir, patterns...) {
		if err != nil {
			return nil, err
		}
		matches = append(matches, path)
	}
	return matches, nil
}

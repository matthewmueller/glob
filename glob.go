package glob

// Glob returns the union of files matching the complete patterns. It derives
// the directory to walk from the non-pattern prefix shared by the patterns.
// Empty patterns are ignored, and no patterns returns no matches.
func Glob(patterns ...string) (matches []string, err error) {
	for path, err := range Walk(patterns...) {
		if err != nil {
			return nil, err
		}
		matches = append(matches, path)
	}
	return matches, nil
}

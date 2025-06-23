package internal

// Match checks if the given name matches any of the provided glob patterns.
// Note: We've reversed the order of parameters from filepath.Match to allow
// for multiple patterns to be passed in at once.
func Match(sep rune, name string, patterns ...string) (bool, error) {
	matcher, err := Compile(sep, patterns...)
	if err != nil {
		return false, err
	}
	return matcher.Match(name), nil
}

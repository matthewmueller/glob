package internal

import (
	"strings"
)

// GCA returns the greatest common ancestor of a list of directories.
func GCA(sep rune, patterns ...string) string {
	if len(patterns) == 0 {
		return "."
	} else if len(patterns) == 1 {
		return patterns[0]
	}
	aParts := strings.Split(patterns[0], string(sep))
	aLen := len(aParts)
	for i := 1; i < len(patterns); i++ {
		bParts := strings.Split(patterns[i], string(sep))
		bLen := len(bParts)
		for i := 0; i < aLen && i < bLen; i++ {
			if aParts[i] != bParts[i] {
				aParts = aParts[:i]
				aLen = i
				break
			}
		}
	}
	if aLen == 0 {
		return "."
	} else if aLen == 1 && aParts[0] == "" {
		return string(sep)
	}
	return clean(sep, strings.Join(aParts, string(sep)))
}

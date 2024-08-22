package glob

import (
	"github.com/gobwas/glob"
)

type Matcher = glob.Glob

// Comple a glob pattern into a matcher
func Compile(pattern string, separators ...rune) (Matcher, error) {
	// Expand patterns like {a,b} into multiple globs a & b. This avoids an
	// infinite loop described in this comment:
	// https://github.com/gobwas/glob/issues/50#issuecomment-1330182417
	patterns, err := expand(pattern)
	if err != nil {
		return nil, err
	}
	matchers := make(matchers, len(patterns))
	for i, pattern := range patterns {
		matcher, err := compile(pattern, separators...)
		if err != nil {
			return nil, err
		}
		matchers[i] = matcher
	}
	return matchers, nil
}

// compile a glob pattern into a matcher
func compile(pattern string, separators ...rune) (Matcher, error) {
	return glob.Compile(pattern, separators...)
}

type matchers []glob.Glob

func (matchers matchers) Match(path string) bool {
	for _, matcher := range matchers {
		if matcher.Match(path) {
			return true
		}
	}
	return false
}

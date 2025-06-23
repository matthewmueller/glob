package internal

import (
	"github.com/gobwas/glob"
)

type Matcher = glob.Glob

// Comple a glob pattern into a matcher
func Compile(sep rune, patterns ...string) (Matcher, error) {
	var all []string
	for _, pattern := range patterns {
		// Expand patterns like {a,b} into multiple globs a & b. This avoids an
		// infinite loop described in this comment:
		// https://github.com/gobwas/glob/issues/50#issuecomment-1330182417
		subpatterns, err := Split(pattern)
		if err != nil {
			return nil, err
		}
		all = append(all, subpatterns...)
	}
	return compile(sep, all...)
}

// compile a glob pattern into a matcher
func compile(sep rune, patterns ...string) (Matcher, error) {
	matchers := make(matchers, len(patterns))
	for i, pattern := range patterns {
		matcher, err := glob.Compile(pattern, sep)
		if err != nil {
			return nil, err
		}
		matchers[i] = matcher
	}
	return matchers, nil
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

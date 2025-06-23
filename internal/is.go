package internal

import "github.com/gobwas/glob/syntax/lexer"

// Is checks if the given pattern is a glob
func Is(pattern string) bool {
	for i := range pattern {
		if lexer.Special(pattern[i]) {
			return true
		}
	}
	return false
}

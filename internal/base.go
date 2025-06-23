package internal

import (
	"strings"

	"github.com/gobwas/glob/syntax/lexer"
)

func Base(sep rune, patterns ...string) string {
	if len(patterns) == 0 {
		return "."
	}
	baseDirs := make([]string, len(patterns))
	for i, pattern := range patterns {
		baseDirs[i] = base(sep, pattern)
	}
	return GCA(sep, baseDirs...)
}

func base(sep rune, pattern string) string {
	parts := strings.Split(pattern, string(sep))
	var segments []string
outer:
	for _, part := range parts {
		lex := lexer.NewLexer(part)
	inner:
		for {
			token := lex.Next()
			switch token.Type {
			case lexer.Text:
				continue
			case lexer.EOF:
				break inner
			default:
				break outer
			}
		}
		segments = append(segments, part)
	}
	if len(segments) == 0 {
		return "."
	}
	base := strings.Join(segments, string(sep))
	if base == "" {
		return string(sep)
	}
	return clean(sep, base)
}

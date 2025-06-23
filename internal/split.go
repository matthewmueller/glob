package internal

import (
	"fmt"

	"github.com/gobwas/glob/syntax/ast"
	"github.com/gobwas/glob/syntax/lexer"
)

func Split(str string) ([]string, error) {
	lex := lexer.NewLexer(str)
	node, err := ast.Parse(lex)
	if err != nil {
		return nil, err
	}
	patterns, err := splitNode(node)
	if err != nil {
		return nil, err
	}
	return unique(patterns), nil
}

func splitNode(node *ast.Node) (patterns []string, err error) {
	prefix := ""
	write := func(value string) {
		prefix += value
		for i := range patterns {
			patterns[i] += value
		}
	}
	for _, child := range node.Children {
		switch child.Kind {
		case ast.KindText:
			text, ok := child.Value.(ast.Text)
			if !ok {
				return nil, fmt.Errorf("expected text value, got %T", child.Value)
			}
			write(text.Text)
		case ast.KindList:
			list, ok := child.Value.(ast.List)
			if !ok {
				return nil, fmt.Errorf("expected list value, got %T", child.Value)
			}
			write("[")
			if list.Not {
				write("^")
			}
			write(list.Chars)
			write("]")
		case ast.KindRange:
			rng, ok := child.Value.(ast.Range)
			if !ok {
				return nil, fmt.Errorf("expected rng value, got %T", child.Value)
			}
			write("[")
			if rng.Not {
				write("^")
			}
			write(string(rng.Lo))
			write("-")
			write(string(rng.Hi))
			write("]")
		case ast.KindAny:
			write("*")
		case ast.KindSuper:
			write("**")
		case ast.KindSingle:
			write("?")
		case ast.KindAnyOf:
			var newPatterns []string
			for _, child := range child.Children {
				results, err := splitNode(child)
				if err != nil {
					return nil, err
				}
				for _, result := range results {
					if len(patterns) == 0 {
						newPatterns = append(newPatterns, prefix+result)
					} else {
						for _, pattern := range patterns {
							newPatterns = append(newPatterns, pattern+result)
						}
					}
				}
			}
			patterns = newPatterns
		default:
			return nil, fmt.Errorf("unknown node kind: %v", child.Kind)
		}
	}
	if len(patterns) == 0 {
		patterns = append(patterns, prefix)
	}
	return patterns, nil
}

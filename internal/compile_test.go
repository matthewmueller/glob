package internal_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob/internal"
)

func TestCompileSinglePattern(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/bar")
	is.NoErr(err)
	is.True(m.Match("foo/bar"))
	is.True(!m.Match("foo/baz"))
}

func TestCompileMultiplePatterns(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/bar", "baz/qux")
	is.NoErr(err)
	is.True(m.Match("foo/bar"))
	is.True(m.Match("baz/qux"))
	is.True(!m.Match("foo/baz"))
}

func TestCompileWithBraces(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/{bar,baz}")
	is.NoErr(err)
	is.True(m.Match("foo/bar"))
	is.True(m.Match("foo/baz"))
	is.True(!m.Match("foo/qux"))
}

func TestCompileStar(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/*/bar")
	is.NoErr(err)
	is.True(m.Match("foo/x/bar"))
	is.True(m.Match("foo/123/bar"))
	is.True(!m.Match("foo/bar"))
}

func TestCompileDoubleStar(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/**/bar")
	is.NoErr(err)
	is.True(m.Match("foo/x/bar"))
	is.True(m.Match("foo/123/bar"))
	is.True(m.Match("foo/bar"))
}

func TestCompileEmptyPattern(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "")
	is.NoErr(err)
	is.True(m.Match(""))
	is.True(!m.Match("foo"))
}

func TestCompileSeparator(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('-', "foo-bar")
	is.NoErr(err)
	is.True(m.Match("foo-bar"))
	is.True(!m.Match("foo/bar"))
}

func TestCompileMultiple(t *testing.T) {
	is := is.New(t)
	m, err := internal.Compile('/', "foo/*/bar", "jan/*/bar")
	is.NoErr(err)
	is.True(m.Match("foo/x/bar"))
	is.True(m.Match("foo/123/bar"))
	is.True(!m.Match("foo/bar"))
	is.True(m.Match("jan/x/bar"))
	is.True(m.Match("jan/123/bar"))
	is.True(!m.Match("jan/bar"))
}

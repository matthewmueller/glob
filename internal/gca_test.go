package internal_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob/internal"
)

func TestGCANoPaths(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/')
	is.Equal(got, ".")
}

func TestGCASinglePath(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b/c")
	is.Equal(got, "/a/b/c")
}

func TestGCACommonAncestor(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b/c", "/a/b/d", "/a/b/e/f")
	is.Equal(got, "/a/b")
}

func TestGCANoCommonAncestor(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b/c", "/x/y/z")
	is.Equal(got, "/")
}

func TestGCARelativePaths(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "a/b/c", "a/b/d", "a/b/e/f")
	is.Equal(got, "a/b")
}

func TestGCAMixedAbsoluteAndRelative(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b/c", "a/b/c")
	is.Equal(got, ".")
}

func TestGCAIdenticalAbsPaths(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b/c", "/a/b/c")
	is.Equal(got, "/a/b/c")
}

func TestGCAIdenticalRelPaths(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "a/b/c", "a/b/c")
	is.Equal(got, "a/b/c")
}

func TestGCAAncestorOfOthers(t *testing.T) {
	is := is.New(t)
	got := internal.GCA('/', "/a/b", "/a/b/c/d", "/a/b/c")
	is.Equal(got, "/a/b")
}

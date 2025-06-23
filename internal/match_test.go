package internal_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob/internal"
)

func TestMatch(t *testing.T) {
	is := is.New(t)
	// Helper function for test cases
	test := func(input string, expect bool, patterns ...string) {
		is.Helper()
		ok, err := internal.Match('/', input, patterns...)
		is.NoErr(err)
		is.Equal(ok, expect)
	}

	// Initial test cases
	test("foo/bar", true, "foo/bar")
	test("foo/baz", false, "foo/bar")

	// TestMatchMultiplePatterns cases
	test("foo/bar", true, "foo/bar", "baz/qux")
	test("baz/qux", true, "foo/bar", "baz/qux")
	test("foo/baz", false, "foo/bar", "baz/qux")

	// TestMatchWithStar cases
	test("foo/x/bar", true, "foo/*/bar")
	test("foo/bar", false, "foo/*/bar")

	// TestMatchWithDoubleStar cases
	test("foo/x/bar", true, "foo/**/bar")
	test("foo/bar", true, "foo/**/bar")

	// TestMatchEmptyPattern cases
	test("", true, "")
	test("foo", false, "")

	// TestMatchNoPatterns cases
	test("foo", false)
}

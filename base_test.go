package glob

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestBase(t *testing.T) {
	is := is.New(t)
	test := func(input, expect string) {
		is.Helper()
		is.Equal(Base(input), expect)
	}
	test(".", ".")
	test(".*", ".")
	test("a/*/b", "a")
	test("a*/.*/b", ".")
	test("*/a/b/c", ".")
	test("*", ".")
	test("*/", ".")
	test("*/*", ".")
	test("*/*/", ".")
	test("**", ".")
	test("**/", ".")
	test("**/*", ".")
	test("**/*/", ".")
	test("/*.js", "/")
	test("*.js", ".")
	test("**/*.js", ".")
	test("{a,b}", ".")
	test("/{a,b}", "/")
	test("/{a,b}/", "/")
	test("{a,b}", ".")
	test("/{a,b}", "/")
	test("./{a,b}", ".")
	test("path/to/*.js", "path/to")
	test("/root/path/to/*.js", "/root/path/to")
	test("chapter/foo [bar]/", "chapter")
	test("path/[a-z]", "path")
	test("[a-z]", ".")
	test("path/{to,from}", "path")
	test("path/!/foo", "path/!/foo")
	test("path/?/foo", "path")
	test("path/+/foo", "path/+/foo")
	test("path/*/foo", "path")
	test("path/@/foo", "path/@/foo")
	test("path/!/foo/", "path/!/foo")
	test("path/?/foo/", "path")
	test("path/+/foo/", "path/+/foo")
	test("path/*/foo/", "path")
	test("path/@/foo/", "path/@/foo")
	test("path/**/*", "path")
	test("path/**/subdir/foo.*", "path")
	test("path/subdir/**/foo.js", "path/subdir")
	test("path/!subdir/foo.js", "path/!subdir/foo.js")
	test("path/{foo,bar}/", "path")
	test("{controller/**.go,view/**}", ".")
}

func TestBases(t *testing.T) {
	is := is.New(t)
	test := func(input string, expects ...string) {
		is.Helper()
		var dirs []string
		patterns, err := expand(input)
		if err != nil {
			is.True(len(expects) > 0) // missing expect
			is.Equal(err.Error(), expects[0])
			return
		}
		for _, pattern := range patterns {
			dirs = append(dirs, Base(pattern))
		}
		dirs = unique(dirs)
		is.Equal(strings.Join(dirs, ", "), strings.Join(expects, ", "))
	}
	test("path/{foo,bar}/", "path/foo", "path/bar")
	test("path/{to,from}", "path/to", "path/from")
	test("path/*/{to,from}", "path")
	test("{controller/**.go,view/**}", "controller, view")
	test("{controller/**.go,controller/**}", "controller")
}

package glob_test

import (
	"testing"

	"github.com/matthewmueller/go-glob"
	"github.com/tj/assert"
)

func TestBase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// should strip glob magic to return parent path
		{".", "."},
		{".*", "."},
		{"a/*/b", "a"},
		{"a*/.*/b", "."},
		{"*/a/b/c", "."},
		{"*", "."},
		{"*/", "."},
		{"*/*", "."},
		{"*/*/", "."},
		{"**", "."},
		{"**/", "."},
		{"**/*", "."},
		{"**/*/", "."},
		{"/*.js", "/"},
		{"*.js", "."},
		{"**/*.js", "."},
		{"{a,b}", "."},
		{"/{a,b}", "/"},
		{"/{a,b}/", "/"},
		{"{a,b}", "."},
		{"/{a,b}", "/"},
		{"./{a,b}", "."},
		{"path/to/*.js", "path/to"},
		{"/root/path/to/*.js", "/root/path/to"},
		{"chapter/foo [bar]/", "chapter"},
		{"path/[a-z]", "path"},
		{"[a-z]", "."},
		{"path/{to,from}", "path"},
		// {"path/!/foo", "path/!"},
		{"path/?/foo", "path"},
		// {"path/+/foo", "path/+"},
		{"path/*/foo", "path"},
		// {"path/@/foo", "path/@"},
		// {"path/!/foo/", "path/!/foo"},
		// {"path/?/foo/", "path"},
		// {"path/+/foo/", "path/+/foo"},
		// {"path/*/foo/", "path"},
		// {"path/@/foo/", "path/@/foo"},
		{"path/**/*", "path"},
		{"path/**/subdir/foo.*", "path"},
		{"path/subdir/**/foo.js", "path/subdir"},
		// {"path/!subdir/foo.js", "path/!subdir"},
		{"path/{foo,bar}/", "path"},

		// should respect escaped characters
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			base := glob.Base(test.input)
			assert.Equal(t, test.expected, base)
		})
	}
}

package glob_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob"
)

func TestMatch(t *testing.T) {
	is := is.New(t)
	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{
		"foo.md":            "foo",
		"bar.md":            "bar",
		"baz.txt":           "baz",
		"qux.markdown":      "qux",
		"quux.mdx":          "quux",
		"sub/topic.md":      "topic",
		"sub/deep/topic.md": "deep",
	})
	files, err := glob.Match(filepath.Join(dir, "**.{md,markdown}"))
	is.NoErr(err)
	is.Equal(len(files), 5)
	is.Equal(files[0], filepath.Join(dir, "bar.md"))
	is.Equal(files[1], filepath.Join(dir, "foo.md"))
	is.Equal(files[2], filepath.Join(dir, "qux.markdown"))
	is.Equal(files[3], filepath.Join(dir, "sub/deep/topic.md"))
	is.Equal(files[4], filepath.Join(dir, "sub/topic.md"))
}

func ExampleMatch() {
	files, err := glob.Match("[A-Z]*.md")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	// Output:
	// Changelog.md
	// License.md
	// Readme.md
}

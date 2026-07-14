package glob_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob"
)

func TestGlob(t *testing.T) {
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
	pattern := filepath.Join(dir, "*.md")
	want, err := filepath.Glob(pattern)
	is.NoErr(err)
	got, err := glob.Glob(pattern)
	is.NoErr(err)
	is.Equal(got, want)
}

func ExampleGlob() {
	files, err := glob.Glob("[A-Z]*.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file)
	}

	// Output:
	// Changelog.md
	// License.md
	// Readme.md
}

func TestGlobNginx(t *testing.T) {
	is := is.New(t)
	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{
		"nginx.http.conf":   "http {}",
		"nginx.stream.conf": "stream {} ",
	})
	pattern := filepath.Join(dir, "nginx.*.conf")
	want, err := filepath.Glob(pattern)
	is.NoErr(err)
	is.Equal(len(want), 2)
	got, err := glob.Glob(pattern)
	is.NoErr(err)
	is.Equal(got, want)
	multiple, err := glob.Glob(
		filepath.Join(dir, "nginx.http.conf"),
		filepath.Join(dir, "nginx.stream.conf"),
	)
	is.NoErr(err)
	is.Equal(multiple, want)
}

func TestGlobMultipleBases(t *testing.T) {
	is := is.New(t)
	dir := t.TempDir()
	writeFiles(t, dir, map[string]string{
		"left/one.md":   "one",
		"left/one.txt":  "one",
		"right/two.md":  "two",
		"right/two.txt": "two",
	})
	matches, err := glob.Glob(
		filepath.Join(dir, "left/*.md"),
		filepath.Join(dir, "right/*.txt"),
	)
	is.NoErr(err)
	is.Equal(matches, []string{
		filepath.Join(dir, "left/one.md"),
		filepath.Join(dir, "right/two.txt"),
	})
}

func TestGlobNoPatterns(t *testing.T) {
	is := is.New(t)
	matches, err := glob.Glob()
	is.NoErr(err)
	is.Equal(matches, nil)
	matches, err = glob.Glob("")
	is.NoErr(err)
	is.Equal(matches, nil)
	matches, err = filepath.Glob("")
	is.NoErr(err)
	is.Equal(matches, nil)
}

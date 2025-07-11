package glob_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob"
)

func writeFiles(t testing.TB, dir string, files map[string]string) error {
	for file, content := range files {
		path := filepath.Join(dir, file)
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			t.Errorf("error creating directory: %v", err)
			return err
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Errorf("error writing file: %v", err)
		}
	}
	return nil
}

func TestWalk(t *testing.T) {
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
	files := []string{}
	for err, file := range glob.Walk(dir, "**.{md,markdown}") {
		is.NoErr(err)
		rel, err := filepath.Rel(dir, file)
		is.NoErr(err)
		files = append(files, rel)
	}
	is.Equal(len(files), 5)
	is.Equal(files[0], "bar.md")
	is.Equal(files[1], "foo.md")
	is.Equal(files[2], "qux.markdown")
	is.Equal(files[3], "sub/deep/topic.md")
	is.Equal(files[4], "sub/topic.md")
}

func TestWalkBreak(t *testing.T) {
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
	for err, file := range glob.Walk(dir, "**.{md,markdown}") {
		is.NoErr(err)
		rel, err := filepath.Rel(dir, file)
		is.NoErr(err)
		if rel == "foo.md" {
			break
		}
	}
}

func TestWalkReturn(t *testing.T) {
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
	for err, file := range glob.Walk(dir, "**.{md,markdown}") {
		is.NoErr(err)
		rel, err := filepath.Rel(dir, file)
		is.NoErr(err)
		if rel == "foo.md" {
			return
		}
	}
}

func TestWalkNone(t *testing.T) {
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
	for err := range glob.Walk(dir, "notadir/**") {
		is.True(errors.Is(err, fs.ErrNotExist))
		break
	}
}

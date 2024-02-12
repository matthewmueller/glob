package glob_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob"
)

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
	err := glob.Walk(filepath.Join(dir, "**.{md,markdown}"), func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		files = append(files, rel)
		return nil
	})
	is.NoErr(err)
	is.Equal(len(files), 5)
	is.Equal(files[0], "bar.md")
	is.Equal(files[1], "foo.md")
	is.Equal(files[2], "qux.markdown")
	is.Equal(files[3], "sub/deep/topic.md")
	is.Equal(files[4], "sub/topic.md")
}

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

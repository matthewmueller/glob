package glob_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

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
	is.Equal(files[2], "sub/deep/topic.md")
	is.Equal(files[3], "sub/topic.md")
	is.Equal(files[4], "qux.markdown")
}

func TestWalkFS(t *testing.T) {
	is := is.New(t)
	fsys := fstest.MapFS{
		"foo.md":            &fstest.MapFile{Data: []byte("foo")},
		"bar.md":            &fstest.MapFile{Data: []byte("bar")},
		"baz.txt":           &fstest.MapFile{Data: []byte("baz")},
		"qux.markdown":      &fstest.MapFile{Data: []byte("qux")},
		"quux.mdx":          &fstest.MapFile{Data: []byte("quux")},
		"sub/topic.md":      &fstest.MapFile{Data: []byte("topic")},
		"sub/deep/topic.md": &fstest.MapFile{Data: []byte("deep")},
	}
	files := []string{}
	err := glob.WalkFS(fsys, "**.{md,markdown}", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	is.NoErr(err)
	is.Equal(len(files), 5)
	is.Equal(files[0], "bar.md")
	is.Equal(files[1], "foo.md")
	is.Equal(files[2], "sub/deep/topic.md")
	is.Equal(files[3], "sub/topic.md")
	is.Equal(files[4], "qux.markdown")
}

func TestWalkBaseFS(t *testing.T) {
	is := is.New(t)
	fsys := fstest.MapFS{
		"foo.md":            &fstest.MapFile{Data: []byte("foo")},
		"bar.md":            &fstest.MapFile{Data: []byte("bar")},
		"baz.txt":           &fstest.MapFile{Data: []byte("baz")},
		"qux.markdown":      &fstest.MapFile{Data: []byte("qux")},
		"quux.mdx":          &fstest.MapFile{Data: []byte("quux")},
		"sub/topic.md":      &fstest.MapFile{Data: []byte("topic")},
		"sub/deep/topic.md": &fstest.MapFile{Data: []byte("deep")},
	}
	files := []string{}
	err := glob.WalkFS(fsys, "sub/**.{md,markdown}", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		return nil
	})
	is.NoErr(err)
	is.Equal(len(files), 2)
	is.Equal(files[0], "sub/deep/topic.md")
	is.Equal(files[1], "sub/topic.md")
}

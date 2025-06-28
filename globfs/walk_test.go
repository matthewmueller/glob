package globfs_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob/globfs"
)

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
	for err, file := range globfs.Walk(fsys, "**.{md,markdown}") {
		is.NoErr(err)
		files = append(files, file)
	}
	is.Equal(len(files), 5)
	is.Equal(files[0], "bar.md")
	is.Equal(files[1], "foo.md")
	is.Equal(files[2], "qux.markdown")
	is.Equal(files[3], "sub/deep/topic.md")
	is.Equal(files[4], "sub/topic.md")
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
	for err, file := range globfs.Walk(fsys, "sub/**.{md,markdown}") {
		is.NoErr(err)
		files = append(files, file)
	}
	is.Equal(len(files), 2)
	is.Equal(files[0], "sub/deep/topic.md")
	is.Equal(files[1], "sub/topic.md")
}

func TestWalkBreak(t *testing.T) {
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
	for err, file := range globfs.Walk(fsys, "sub/**.{md,markdown}") {
		is.NoErr(err)
		if file == "foo.md" {
			break
		}
	}
}

func TestWalkReturn(t *testing.T) {
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
	for err, file := range globfs.Walk(fsys, "sub/**.{md,markdown}") {
		is.NoErr(err)
		if file == "foo.md" {
			return
		}
	}
}

func TestWalkNone(t *testing.T) {
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
	for err := range globfs.Walk(fsys, "notadir/**") {
		is.True(errors.Is(err, fs.ErrNotExist))
		break
	}
}

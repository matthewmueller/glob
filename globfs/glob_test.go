package globfs_test

import (
	"fmt"
	"testing"
	"testing/fstest"

	"github.com/matryer/is"
	"github.com/matthewmueller/glob/globfs"
)

func TestGlob(t *testing.T) {
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
	files, err := globfs.Glob(fsys, "**.{md,markdown}")
	is.NoErr(err)
	is.Equal(len(files), 5)
	is.Equal(files[0], "bar.md")
	is.Equal(files[1], "foo.md")
	is.Equal(files[2], "qux.markdown")
	is.Equal(files[3], "sub/deep/topic.md")
	is.Equal(files[4], "sub/topic.md")

	files, err = globfs.Glob(fsys, "foo.md", "bar.md")
	is.NoErr(err)
	is.Equal(files, []string{"bar.md", "foo.md"})
}

func TestGlobNoPatterns(t *testing.T) {
	is := is.New(t)
	fsys := fstest.MapFS{}
	files, err := globfs.Glob(fsys)
	is.NoErr(err)
	is.Equal(files, []string(nil))
	files, err = globfs.Glob(fsys, "")
	is.NoErr(err)
	is.Equal(files, []string(nil))
}

func TestNested(t *testing.T) {
	is := is.New(t)
	fsys := fstest.MapFS{
		"App.svelte":               &fstest.MapFile{Data: []byte("foo")},
		"lib/App.svelte":           &fstest.MapFile{Data: []byte("foo")},
		"lib/src/App.svelte":       &fstest.MapFile{Data: []byte("foo")},
		"lib/src/shared.svelte.js": &fstest.MapFile{Data: []byte("foo")},
		"lib/src/shared.js":        &fstest.MapFile{Data: []byte("foo")},
	}
	files, err := globfs.Glob(fsys, "**.svelte{,.js}")
	is.NoErr(err)
	is.Equal(len(files), 4)
	is.Equal(files[0], "App.svelte")
	is.Equal(files[1], "lib/App.svelte")
	is.Equal(files[2], "lib/src/App.svelte")
	is.Equal(files[3], "lib/src/shared.svelte.js")

	files, err = globfs.Glob(fsys, "{**,*}/*.svelte{,.js}")
	is.NoErr(err)
	is.Equal(len(files), 3)
	is.Equal(files[0], "lib/App.svelte")
	is.Equal(files[1], "lib/src/App.svelte")
	is.Equal(files[2], "lib/src/shared.svelte.js")
}

func ExampleGlob() {
	fsys := fstest.MapFS{
		"posts/first.md":  &fstest.MapFile{Data: []byte("first")},
		"posts/second.md": &fstest.MapFile{Data: []byte("second")},
		"posts/draft.txt": &fstest.MapFile{Data: []byte("draft")},
	}

	files, err := globfs.Glob(fsys, "posts/*.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println(file)
	}

	// Output:
	// posts/first.md
	// posts/second.md
}

# glob

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewmueller/glob.svg)](https://pkg.go.dev/github.com/matthewmueller/glob)

Glob utilities built on top of the well-tested [gobwas/glob](https://github.com/gobwas/glob).

## Features

- Supports advanced patterns such as `**` and `{md,markdown}`
- Accepts multiple patterns and returns their union without duplicate paths
- Smarter traversal by starting from the non-glob base
  - e.g. `/some/base/**.md`, starts traversing from `/some/base/`

## Install

```sh
go get github.com/matthewmueller/glob
```

## Example

`Glob` collects matching paths. `filepath.Glob` with extras:

```go
package main

import (
	"fmt"

	"github.com/matthewmueller/glob"
)

func main() {
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
```

`Walk` streams the same matches:

```go
for file, err := range glob.Walk("*.md", "docs/**/*.md") {
	if err != nil {
		fmt.Println(err)
		break
	}
	fmt.Println(file)
}
// Output:
// Changelog.md
// License.md
// Readme.md
```

## GlobFS

The `globfs` package provides the same API for an `fs.FS`. Patterns use
slash-separated paths relative to the filesystem root.

```go
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
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT

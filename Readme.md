# glob

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewmueller/glob.svg)](https://pkg.go.dev/github.com/matthewmueller/glob)

Glob utilities built on top of the well-tested [gobwas/glob](https://github.com/gobwas/glob).

## Features

- Supports wildcard globs like `**`
- Faster matching and walking by pulling out the non-wildcard parts of a glob (e.g. `/posts/` in `/posts/**`)
- Supports globs like `{posts/**,tags/**}` by pre-expanding globs.

## Install

```sh
go get github.com/matthewmueller/glob
```

## Example

```go
package main

import (
	"github.com/matthewmueller/glob"
)

func main() {
	files, _ := glob.Glob("[A-Z]*.md")
	for _, file := range files {
		fmt.Println(file)
	}
	// Output:
	// Changelog.md
	// License.md
	// Readme.md
}
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT

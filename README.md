# X11 Colors for golang

The package provides X11 colors names as variables, together with colors caption names and couple additional helper functions.

Use the following command to install the package:

```
go get -u github.com/vgarvardt/x11colors-go
```

## Features

All colors are available as package-level variables and can be used in the following manner:

```go
package main

import (
	"fmt"

	"github.com/vgarvardt/x11colors-go"
)

func main() {
	someColor := x11colors.RandomSeeded()

	fmt.Printf(
		`You are using "%s" color (slugified "%s") with RGB HEX #%02X%02X%02X that should use %s text color when used as background%s`,
		someColor.Name,
		someColor.Name.Slugify(),
		someColor.RGBA.R,
		someColor.RGBA.G,
		someColor.RGBA.B,
		map[bool]string{true: "BLACK", false: "WHITE"}[someColor.CaptionBlack],
		"\n",
	)
}

```

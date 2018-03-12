# X11 Colors for golang

The package provides X11 colors names as variables, together with colors caption names and couple additional helper functions.

Use the following command to install the package:

```
go get -u github.com/vgarvardt/x11colors-go
```

## Usage example

All colors are available as package-level variables and can be used in the following manner:

```go
package main

import (
	"fmt"

	"github.com/vgarvardt/x11colors-go"
)

func printColorDetails(color x11colors.X11Color) {
	fmt.Printf(
		`You are using "%s" color (slugified "%s") with RGB HEX #%02X%02X%02X that should use %s text color when used as background%s`,
		color.Name,
		color.Name.Slugify(),
		color.RGBA.R,
		color.RGBA.G,
		color.RGBA.B,
		map[bool]string{true: "BLACK", false: "WHITE"}[color.CaptionBlack],
		"\n",
	)
}

func main() {
	printColorDetails(x11colors.GreenX11)
	printColorDetails(x11colors.RandomSeeded())
}
```

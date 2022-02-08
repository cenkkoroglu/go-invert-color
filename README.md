# go-invert-color

[![GoDoc][godoc-image]][godoc-url]

> Generates inverted color of the given color in Go.

# Information

- Accepts 3 or 6 length of Hex color code (#FFF, #F0F0F0, etc.)
- Accepts color Hex code with # or without # (#FFF, FFF are accepted)

# Usage

The basic level example:

```go
package main

import (
	"fmt"

	"github.com/cenkkoroglu/go-invert-color"
)

func main() {
	color := "#FFF"
	invertedColor, err := invert.Invert(color)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("inverted color of %s is %s", color, invertedColor)
}
```


[godoc-image]: https://godoc.org/github.com/cenkkoroglu/go-invert-color?status.svg

[godoc-url]: https://godoc.org/github.com/cenkkoroglu/go-invert-color

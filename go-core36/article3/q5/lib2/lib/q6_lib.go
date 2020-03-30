package lib

import (
	"fmt"
	"io"
)

func Hello(w io.Writer, name string) {
	fmt.Fprintf(w, "Hi, %s!\n", name)
}


package lib

import (
	"os"
	in "gopl.io/go-core36/article3/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}

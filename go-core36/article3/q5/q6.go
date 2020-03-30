package main

import (
	"flag"
	"gopl.io/go-core36/article3/q5/lib"
	lib2 "gopl.io/go-core36/article3/q5/lib2/lib"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	lib.Hello(os.Stdout, name)
	lib2.Hello(os.Stdout, name)
}

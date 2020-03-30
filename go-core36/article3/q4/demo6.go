package main

import (
	"flag"
	"gopl.io/go-core36/article3/q4/lib"
	//in "gopl.io/go-core36/article3/q4/lib/internal" // 此行无法通过编译。
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}

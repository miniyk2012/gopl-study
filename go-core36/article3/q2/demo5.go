package main

import (
	"flag"
	"gopl.io/go-core36/article3/q2/lib"
)



var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib5.Hello(name)
}

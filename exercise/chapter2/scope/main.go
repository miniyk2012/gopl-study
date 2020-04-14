package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string


func init() {
	var err error
	cwd, err = os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println(cwd)
}
func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
	fmt.Printf("\ncwd=%s", cwd)
}

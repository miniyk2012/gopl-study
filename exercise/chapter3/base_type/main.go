package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	fmt.Println(x)
	var y uint8 = 11&^11
	fmt.Println(y)
}


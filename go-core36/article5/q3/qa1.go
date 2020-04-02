package main

import (
	. "flag"
	"fmt"
)

//var ErrHelp = 4
func main() {
	fmt.Println(ErrHelp)
	{
		ErrHelp := 6
		fmt.Println(ErrHelp)
	}
}

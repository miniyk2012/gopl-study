package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	fmt.Println(x)
	var y uint8 = 11 &^ 11
	fmt.Println(y)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%[2]d %[2]c %[2]q\n", "", ascii)   // "97 a 'a'"
	fmt.Printf("%[2]d %[2]c %[2]q\n", "", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)              // "10 '\n'"
	f := 65
	fmt.Printf("%0x\n", int(f))
	fmt.Printf("%o\n", int(f))
	fmt.Println(int(f))

	nan := math.NaN()
	fmt.Println(math.Inf(-1))
	fmt.Println(math.IsNaN(nan))
	fmt.Println(nan == nan)
}

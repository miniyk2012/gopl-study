// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/popcount"
	"gopl.io/ch2/tempconv"
)

var b = a + 1
var a = 10
func init() {fmt.Println(1)}
func init() {fmt.Println(2)}

func main() {
	fmt.Println(b)
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n",
			k, tempconv.KToC(k), c, tempconv.CToK(c))
	}
	fmt.Println(byte(0xff>>(0*8)))
	fmt.Println(popcount.PopCount(0xff))
}
func init() {fmt.Println(3)}


//!-

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

type A struct {
	i []int
	j int
}

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
	fmt.Println(f() == f())
	fmt.Println(f())
	a := A{[]int{10}, 5}
	fmt.Println(a)
	change(a)
	fmt.Println(a)
	z := []int{10}
	changZ(z)
	fmt.Println(z)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
func changZ(z []int) {
	z[0] = 5
	fmt.Println(z)
}

func change(a A) {
	a.i[0] = 100
	a.j = 500
	fmt.Println(a)
}
func f() *int {
	v := 1
	return &v
}

//!-

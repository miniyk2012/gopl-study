// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "", "three"}
	newData := nonempty(data)
	fmt.Printf("%q %d\n", newData, len(newData)) // `["one" "three"]`
	fmt.Printf("%q\n", data)                     // `["one" "three" "three"]`
	//!-main

	data2 := []string{"one", "", "three"}
	newData2 := nonempty2(data2)
	fmt.Printf("%q %d\n", newData2, len(newData2)) // `["one" "three"]`
	fmt.Printf("%q\n", data2)                      // `["one" "three" "three"]`

	// remove
	s := []int{5, 6, 7, 8, 9}
	s = remove(s, 2)
	fmt.Printf("cap(s)=%d, len(s)=%d, s=%v\n", cap(s), len(s), s) // "cap(s)=5, len(s)=4, s=[5 6 8 9]"
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	c3 := sha256.Sum256([]byte("x\n"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%x\n", c3)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// 73cb3858a687a8494ca3323053016282f3dad39d42cf62ca4e79dda2aac7d9ac
	// false
	// [32]uint8
	fmt.Printf("%x\n", c1)
	zero(&c1)
	fmt.Printf("%x\n", c1)
}

//!-

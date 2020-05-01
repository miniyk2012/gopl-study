package main

import (
	"fmt"
	"math/cmplx"
	"testing"
)

func TestX(t *testing.T) {
	var a float64 = 10
	fmt.Println(3+6i)
	fmt.Println((3+6i)*1i)
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)"
	fmt.Println(complex(a, 0)+8i)   // (10+8i)
	fmt.Printf("%T\n", complex(a, 0)+8i)   // complex128
	fmt.Printf("%T\n", 1+8.78i)  // complex128
}

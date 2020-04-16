package main

import "fmt"

func Shrink(s []int) []int {
	length := len(s)
	newS := make([]int, length)
	copy(newS, s)
	return newS
}
func main()  {
	s := make([]int, 0)
	fmt.Printf("The capacity of s: %d\n", cap(s))
	for i := 1; i <= 5; i++ {
		s = append(s, i)
		fmt.Printf("s(%d): len: %d, cap: %d\n", i, len(s), cap(s))
	}
	fmt.Printf("len=%d, cap=%d, s=%v\n", len(s), cap(s), s)
	newS := Shrink(s)
	fmt.Printf("len=%d, cap=%d, newS=%d\n", len(newS), cap(newS), newS)
}

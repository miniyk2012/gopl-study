package main

import (
	"fmt"
	"time"
)

func SameChars(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	if len(r1) != len(r2) {
		return false
	}
	for _, r := range r1 {
		var flag bool
		for _, rr := range r2 {
			if r == rr {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(SameChars("世界", "世"))
	fmt.Println(SameChars("世界", "世界"))
	fmt.Println(SameChars("a世界", "世a界"))
	fmt.Println(SameChars("a世界", "世a界世"))
	fmt.Println(SameChars("hello, 世界", "世he,界 llo"))
	fmt.Printf("%T %[1]v\n", time.Microsecond)
}

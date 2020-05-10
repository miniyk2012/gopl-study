package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s)
}

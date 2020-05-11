package main

import "fmt"

func removeSucEmpty(s []string) []string {
	last := s[0]
	j := 1
	for i := 1; i<len(s); i++ {
		if s[i] != last {
			s[j] = s[i]
			j++
		}
		last = s[i]
	}
	s = s[:j]
	return s
}
func main() {
	s := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	s = removeSucEmpty(s)
	fmt.Println(s)
}

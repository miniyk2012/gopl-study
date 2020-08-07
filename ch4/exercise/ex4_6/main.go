package main

import (
	"fmt"
	"unicode"
)

func replaceSpace(b []byte) []byte {
	s := string(b)
	r := ""
	flag := false
	for _, v := range s {
		if unicode.IsSpace(v) {
			if !flag {
				r += string(v)
			} 
			flag = true
		} else {
			r += string(v)
			flag = false
		}
	}

	return []byte(r)
}
func main() {
	s := []byte("a  abbb fdsa眼   科  ")
	got := replaceSpace(s)
	fmt.Println(string(got))

}

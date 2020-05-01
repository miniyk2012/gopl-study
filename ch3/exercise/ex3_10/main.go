package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		count++
		buf.WriteByte(s[i])
		if count != 1 && i != 0 && count%3 == 0 {
			buf.WriteByte(',')
		}
	}
	s2 := buf.String()
	var buf2 bytes.Buffer
	for i := len(s2) - 1; i >= 0; i-- {
		buf2.WriteByte(s2[i])
	}
	return buf2.String()
}

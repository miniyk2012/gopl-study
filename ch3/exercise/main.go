package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "y恺"
	fmt.Println(len(s))
	fmt.Println(s[0], s[1], s[2], s[3])
	fmt.Println(s[1:4])
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`
	fmt.Println(GoUsage)
	s1 := "Hello, 世界"
	fmt.Println(len(s1))                    // 13个字节
	fmt.Println(utf8.RuneCountInString(s1)) // 只对应9个Unicode字符
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	a := '时'
	fmt.Printf("%T\n", a)
	fmt.Println()
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for range s1 {
		n++
	}
	fmt.Println(n)  // 9个unicode字符
	
	r1 := []rune("Hello, 世界")
	r1[0] = '傻'
	fmt.Printf("%s\n", string(r1))

	fmt.Println(string(1234567)) // "?"
}

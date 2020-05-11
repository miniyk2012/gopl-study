package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func funcA() (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			debug.PrintStack()
			err = errors.New(p.(string))
		}
	}()
	//regexp.MustCompile("(?<a=pattern)")
	return funcB()
}

func funcB() error {
	// simulation

	panic("foo")
	return errors.New("success")
}

func test() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\n")
	} else {
		fmt.Printf("err is %v\n", err)
	}
	s := "abc"
	fmt.Println(s[1:])
}

func testMap() {
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
	fmt.Println(ages["abc"])
	delete(ages, "abc")
	//ages["abc"] = 1 // panic
}

func appendDemo() {
	var runes = []rune("Hello, 世界")

	//for _, r := range "Hello, 世界" {
	//	runes = append(runes, r)
	//}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}

var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func mapDemo() {
	a := []string{"12", "56"}
	Add(a)
	Add(a)
	Add(a)
	fmt.Println(m)
	Count(a)
}
func main() {
	//test()
	//appendDemo()
	//testMap()
	mapDemo()
}

package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sort"
	"time"
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

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
func capSlice() {
	var z []int = make([]int, 3, 10)
	fmt.Printf("%v\n", z)
	fmt.Printf("len=%d, cap=%d\n", len(z), cap(z))
	z = z[:6]
	z[5] = 100
	fmt.Printf("%v\n", z)
	fmt.Printf("len=%d, cap=%d", len(z), cap(z))
}
func mapDemo() {
	a := []string{"12", "56"}
	Add(a)
	Add(a)
	Add(a)
	fmt.Println(m)
	Count(a)
}

func sortMap() {
	ages := map[string]int{
		"b恺": 30,
		"a逼": 50,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println(len(ages) == 2) // "true"

	var empty map[string]int
	fmt.Println(empty["adf"]) // 0
	adf, ok := empty["adf"]
	fmt.Println(adf, ok) // 0 false

}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee {
	return &Employee{}
	/* ... */
}

type address struct {
	hostname string
	port     int
}



func structDemo() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Printf("%v\n", hits)
}

func main() {
	//test()
	//appendDemo()
	//testMap()
	//mapDemo()
	//capSlice()
	//sortMap()
	var dilbert Employee
	dilbert.Name = "yangk"
	dilbert.Salary -= 5000
	fmt.Println(dilbert)
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
	structDemo()
}

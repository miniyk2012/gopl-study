package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	INDEX                 = 1
	noDelay time.Duration = 3 * time.Minute
	timeout
	KB           = 1000
	MB           = 1000 * KB
	GB           = 1000 * MB
	PB           = 1000 * GB
	Pi64 float64 = math.Pi
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(KiB)
	fmt.Println(YiB / ZiB)
	//fmt.Println(YiB)       // overflows int
	var f float64 = 3 + 0i // untyped complex -> float64
	f = 2                  // untyped integer -> float64
	f = 1e123              // untyped floating-point -> float64
	f = 'a'                // untyped rune -> float64
	fmt.Printf("%T %[1]v\n", f)

	fmt.Printf("%T %[1]v\n", math.Pi)
	var yy complex128 = complex128(Pi64)
	fmt.Printf("%T %[1]v\n", 0i)
	var xy rune = '\000'
	fmt.Printf("%T %[1]v\n", xy)
	fmt.Println(yy)
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
	fmt.Println(n) // 9个unicode字符

	r1 := []rune("Hello, 世界")
	r1[0] = '傻'
	fmt.Printf("%s\n", string(r1))

	fmt.Println(string(1234567)) // "?"

	fmt.Println(strings.Count("afsdafad", "af"))
	fmt.Println(strings.Fields("adf afs"))
	fmt.Println(strings.Fields(""))
	fmt.Println(strings.Join([]string{"afd", "dff"}, " "))

	fmt.Println()
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	s3 := fmt.Sprintf("%b", x)
	fmt.Println(s3, strconv.FormatInt(int64(x), 2))
	z, _ := strconv.Atoi("123") // x is an int
	fmt.Println(z)

	fmt.Println()

	zz, _ := strconv.ParseInt("1238389393930", 10, 64) // base 10, up to 64 bits
	fmt.Printf("%T %[1]d\n", zz, zz)

	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 3m0s"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 3m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

}

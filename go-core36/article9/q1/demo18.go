package main

import "fmt"

func main() {
	// 示例1。
	//var badMap1 = map[[]int]int{} // 这里会引发编译错误。
	//_ = badMap1

	// 示例2。
	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, // 这里会引发panic。
	//	3:        3,
	//}
	//_ = badMap2

	// 示例3。
	//var badMap3 map[[1][]string]int // 这里会引发编译错误。
	//_ = badMap3

	// 示例4。
	//type BadKey1 struct {
	//	slice []string
	//}
	//var badMap4 map[BadKey1]int // 这里会引发编译错误。
	//_ = badMap4

	// 示例5。
	//var badMap5 map[[1][2][3][]string]int // 这里会引发编译错误。
	//_ = badMap5

	// 示例6。
	type BadKey2Field1 struct {
		slice [2]string
	}
	type BadKey2 struct {
		field BadKey2Field1
	}
	var goodMap6 = make(map[BadKey2]int, 3)
	_ = goodMap6

	var goodMap7 = make(map[*BadKey2]int, 3)
	_ = goodMap7
	fmt.Println(len(goodMap7))
	b1 := BadKey2{BadKey2Field1{[...]string{"1", "2"}}}
	b2 := BadKey2{BadKey2Field1{[...]string{"1", "2"}}}
	goodMap6[b1] = 100
	goodMap6[b2] = 300
	fmt.Println(goodMap6)
	goodMap7[&b1] = 5
	goodMap7[&b2] = 6
	fmt.Println(goodMap6)
	if v, ok := goodMap6[b1]; ok {
		fmt.Println(v)
	}
	if v, ok := goodMap6[b2]; ok {
		fmt.Println(v)
	}
	if v, ok := goodMap7[&b1]; ok {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(goodMap7)
	b1.field.slice[0] = "12"
	z := &b1
	fmt.Println(goodMap6)
	fmt.Println(goodMap7)
	if v, ok := goodMap6[b1]; ok {
		fmt.Println(v)  // 不打印
	}
	if v, ok := goodMap6[b2]; ok {
		fmt.Println(v)  // 打印, 因此map里真存了一个BadKey2
	}
	if v, ok := goodMap6[b2]; ok {
		fmt.Println(v)
	}
	if v, ok := goodMap7[z]; ok {
		fmt.Println(v)
	}
}

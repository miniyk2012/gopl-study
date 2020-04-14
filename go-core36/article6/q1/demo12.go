package main

import (
	"fmt"
	"strconv"
)

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// 方式1。
	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		container[1], container)

	// 方式2。
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)

	var srcInt = int16(-255)
	dstInt := int8(srcInt)
	fmt.Println(dstInt, srcInt)
	fmt.Println(string(97))
	fmt.Println(strconv.Itoa(97))
	fmt.Println(strconv.Atoi("10"))
	fmt.Println([]byte("你好"))
	fmt.Printf("%U\n", []rune("你好"))
	fmt.Printf("%q\n", []rune{'\u4F60'})
}

func getElement(containerI interface{}) (elem string, err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", containerI)
		return
	}
	return
}

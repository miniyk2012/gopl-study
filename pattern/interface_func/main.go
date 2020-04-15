package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func ForEach(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

// 这个就是接口型函数
// 接口型函数，指的是用函数实现接口，这样在调用的时候就会非常简便，我称这种函数，为接口型函数，这种方式适用于只有一个函数的接口。
type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}


func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}



func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	Each(persons, HandlerFunc(selfInfo))
	ForEach(persons, selfInfo)
}

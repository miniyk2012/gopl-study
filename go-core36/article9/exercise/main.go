package main

import (
	"fmt"
	"sync"
	"time"
)



func main() {
	m := map[int]string {
		1 : "haha",
	}
	//开箱即用
	var sm sync.Map
	sm.Store(1, "a")
	go syncRead(&sm)
	go read(m)
	time.Sleep(time.Second)
	go write(m)
	go syncWrite(&sm)
	time.Sleep(30*time.Second)
	fmt.Println(m)


}

func syncRead(sm *sync.Map) {
	for {
		if v,ok:=sm.Load(1);ok{
			fmt.Println(v)
		}
		time.Sleep(time.Second)

	}
}
func syncWrite(sm *sync.Map) {
	for {
		//store 方法,添加元素
		sm.Store(1, "b")
		time.Sleep(time.Second)
	}
}
func read(m map[int]string) {
	for {
		_ = m[1]
		fmt.Println(m[1])

		time.Sleep(time.Second)
	}
}

func write(m map[int]string) {
	for {
		m[1] = "write"
		//_ = m[1]
		fmt.Println(m[1])
		time.Sleep(2 * time.Second)
	}
}



// go run -race main.go

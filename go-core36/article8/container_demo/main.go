package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	link := list.New()

	// 循环插入到头部
	for i := 0; i <= 10; i++ {
		fmt.Println(link.PushBack(i).Value)
	}

	// 遍历链表
	for p := link.Front(); p != nil; p = p.Next() {
		fmt.Println("Number", p.Value)
	}
	fmt.Println()
	var l list.List
	fmt.Println(l.Len())
	fmt.Println(l.Front())
	l.PushBack(100)
	fmt.Println(l.Len())
	fmt.Println(l.Front().Value)
	println()
	RingDemo()
	
}

func RingDemo() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
	

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	println()
	// Create two rings, r and s, of size 2
	r = ring.New(3)
	s := ring.New(2)

	// Get the length of the ring
	lr := r.Len()
	ls := s.Len()

	// Initialize r with 0s
	for i := 0; i < lr; i++ {
		r.Value = i
		r = r.Next()
	}

	// Initialize s with 1s
	for j := 0; j < ls; j++ {
		s.Value = j + 10
		s = s.Next()
	}

	// Link ring r and ring s
	rs := r.Link(s)

	// Iterate through the combined ring and print its contents
	rs.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}

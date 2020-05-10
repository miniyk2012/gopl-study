package main

import (
	"fmt"
	"math"
)

// 向右旋转
func rotate(s []int, i int) {
	var temp []int
	i = i % len(s)
	temp = make([]int, int(math.Abs(float64(i))))
	if i > 0 {
		copy(temp, s[len(s)-i:])
		copy(s[i:], s[:len(s)-i])
		copy(s[:i], temp)
	} else {
		i = -i
		copy(temp, s[:i])
		copy(s[:len(s)-i], s[i:])
		copy(s[len(s)-i:], temp)
	}

}

func main() {
	s1 := []int{1, 2, 3, 4}
	rotate(s1, 6)
	fmt.Println(s1) // [3 4 1 2]

	s2 := []int{1, 2, 3, 4, 5}
	rotate(s2, -6)
	fmt.Println(s2) // [2 3 4 5 1]
}

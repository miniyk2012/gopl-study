package main

import (
	"crypto/sha256"
	"fmt"
)

// 返回bit为1的个数
func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func DiffSum256(c1, c2 []byte) int {
	diffCount := 0
	for i := 0; i < len(c1) || i < len(c2); i++ {
		switch {
		case i >= len(c1):
			diffCount += popCount(c2[i])
		case i >= len(c2):
			diffCount += popCount(c1[i])
		default:
			diffCount += popCount(c1[i] ^ c2[i])
		}
	}
	return diffCount
}
func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Println(DiffSum256(c1[:], c2[:]))
	// 0001,0010,0011
	// 0100,0101,1101
	fmt.Println(DiffSum256([]byte{1, 2, 3}, []byte{4, 5, 13}))
}

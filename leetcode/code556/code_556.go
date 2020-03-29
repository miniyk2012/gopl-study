package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	var dstR, dstC = -1, c - 1
	var ret [][]int
	for row, rows := range nums {
		for col := range rows {
			dstC++
			if dstC == c {
				dstR++
				dstC = 0
				ret = append(ret, make([]int, c, c))
			}
			ret[dstR][dstC] = nums[row][col]
		}
	}
	return ret
}

func main() {
	var nums = [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(nums, 1, 4))
	for _, i := range []int{12, 3} {
		fmt.Println(i)
	}
}

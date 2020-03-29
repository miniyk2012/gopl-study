package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var src_r, src_c = len(nums), len(nums[0])
	if src_r*src_c != r*c {
		return nums
	}
	var dst_r, dst_c = -1, c - 1
	var ret = [][]int{}
	for row := 0; row < src_r; row++ {
		for col := 0; col < src_c; col++ {
			dst_c++
			if dst_c == c {
				dst_r++
				dst_c = 0
				ret = append(ret, make([]int, c, c))
			}
			ret[dst_r][dst_c] = nums[row][col]
		}
	}
	return ret
}

func main() {
	var nums = [][]int{
		{1, 2},
		{3, 4},
	}
	matrixReshape(nums, 1, 4)
}

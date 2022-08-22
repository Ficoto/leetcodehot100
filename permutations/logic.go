package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

// https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	var res [][]int
	backtrack(&res, nums, 0)
	return res
}

func backtrack(res *[][]int, output []int, first int) {
	if first >= len(output) {
		*res = append(*res, output)
		return
	}
	for i := first; i < len(output); i++ {
		output[i], output[first] = output[first], output[i]
		tmp := copySlice(output)
		backtrack(res, tmp, first+1)
		output[i], output[first] = output[first], output[i]
	}
}

func copySlice(target []int) []int {
	var res = make([]int, len(target))
	for i, item := range target {
		res[i] = item
	}
	return res
}

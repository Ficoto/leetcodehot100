package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}

// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for a := 0; a < len(nums)-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		if nums[a] > 0 {
			break
		}
		c := len(nums) - 1
		for b := a + 1; b < len(nums)-1; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[a]+nums[b]+nums[c] > 0 {
				c--
			}
			if b >= c {
				break
			}
			if nums[a]+nums[b]+nums[c] != 0 {
				continue
			}
			res = append(res, []int{nums[a], nums[b], nums[c]})
		}
	}
	return res
}

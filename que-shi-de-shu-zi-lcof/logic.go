package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
}

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	return findMissingNumber(nums)
}

func findMissingNumber(nums []int) int {
	if len(nums) <= 3 {
		for index := 0; index < len(nums)-1; index++ {
			if nums[index+1]-nums[index] == 1 {
				continue
			}
			return nums[index] + 1
		}
		return nums[len(nums)-1] + 1
	}
	midIndex := (nums[0] + nums[len(nums)-1]) / 2
	if nums[midIndex-nums[0]] == midIndex {
		return findMissingNumber(nums[midIndex-nums[0]:])
	}
	return findMissingNumber(nums[:midIndex-nums[0]])
}

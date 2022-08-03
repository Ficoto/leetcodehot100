package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func search(nums []int, target int) int {
	targetIndex := findTargetIndex(nums, target)
	if targetIndex == -1 {
		return 0
	}
	var res = 1
	for tmp := targetIndex - 1; tmp >= 0; tmp-- {
		if nums[tmp] != target {
			break
		}
		res += 1
	}
	for tmp := targetIndex + 1; tmp < len(nums); tmp++ {
		if nums[tmp] != target {
			break
		}
		res += 1
	}
	return res
}

func findTargetIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}
	if nums[0] > target {
		return -1
	}
	if nums[len(nums)-1] < target {
		return -1
	}
	midIndex := len(nums) / 2
	if nums[midIndex] == target {
		return midIndex
	}
	if nums[midIndex] < target {
		targetIndex := findTargetIndex(nums[midIndex+1:], target)
		if targetIndex == -1 {
			return -1
		}
		return midIndex + targetIndex + 1
	}
	return findTargetIndex(nums[:midIndex], target)
}

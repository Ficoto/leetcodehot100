package main

import "fmt"

func main() {
	fmt.Println(search([]int{1}, 0))
}

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	maxIndex := searchMaxIndex(nums)
	if nums[maxIndex] == target {
		return maxIndex
	}
	if nums[maxIndex] < target {
		return -1
	}
	if nums[0] <= target {
		return searchTarget(nums[:maxIndex], target)
	}
	searchIndex := searchTarget(nums[maxIndex:], target)
	if searchIndex == -1 {
		return -1
	}
	return maxIndex + searchIndex
}

func searchMaxIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[len(nums)-1] >= nums[0] {
		return len(nums) - 1
	}
	index := len(nums) / 2
	if nums[index] >= nums[0] {
		return index + searchMaxIndex(nums[index:])
	}
	return searchMaxIndex(nums[:index])
}

func searchTarget(sub []int, target int) int {
	if len(sub) == 0 {
		return -1
	}
	if len(sub) == 1 {
		if sub[0] == target {
			return 0
		}
		return -1
	}
	midIndex := len(sub) / 2
	if sub[midIndex] == target {
		return midIndex
	}
	if sub[midIndex] < target {
		searchIndex := searchTarget(sub[midIndex+1:], target)
		if searchIndex == -1 {
			return -1
		}
		return midIndex + searchIndex + 1
	}
	return searchTarget(sub[:midIndex], target)
}

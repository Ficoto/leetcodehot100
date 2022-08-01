package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{1}, 1))
}

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	tIndex := findTargetIndex(nums, target)
	if tIndex == -1 {
		return []int{-1, -1}
	}
	var minIndex, maxIndex = tIndex, tIndex
	for {
		if minIndex < 0 {
			minIndex++
			break
		}
		if nums[minIndex] == target {
			minIndex--
			continue
		}
		minIndex++
		break
	}
	for {
		if maxIndex >= len(nums) {
			maxIndex--
			break
		}
		if nums[maxIndex] == target {
			maxIndex++
			continue
		}
		maxIndex--
		break
	}
	return []int{minIndex, maxIndex}
}

func findTargetIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	midIndex := len(nums) / 2
	if nums[midIndex] == target {
		return midIndex
	}
	if nums[midIndex] < target {
		tIndex := findTargetIndex(nums[midIndex+1:], target)
		if tIndex == -1 {
			return tIndex
		}
		return midIndex + findTargetIndex(nums[midIndex+1:], target) + 1
	}
	return findTargetIndex(nums[:midIndex], target)
}

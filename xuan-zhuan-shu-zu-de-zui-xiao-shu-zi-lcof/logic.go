package main

import "fmt"

func main() {
	fmt.Println(minArray([]int{10, 10, 10, 1, 10}))
}

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 2 {
		var res = numbers[0]
		if res > numbers[1] {
			res = numbers[1]
		}
		return res
	}
	if numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}
	midIndex := len(numbers) / 2
	if numbers[midIndex] < numbers[len(numbers)-1] {
		return minArray(numbers[:midIndex+1])
	}
	if numbers[midIndex] > numbers[len(numbers)-1] {
		return minArray(numbers[midIndex+1:])
	}
	if numbers[midIndex] > numbers[0] {
		return numbers[0]
	}
	if numbers[midIndex] < numbers[0] {
		return minArray(numbers[:midIndex+1])
	}
	lMin := minArray(numbers[:midIndex+1])
	rMin := minArray(numbers[midIndex+1:])
	if lMin < rMin {
		return lMin
	}
	return rMin
}

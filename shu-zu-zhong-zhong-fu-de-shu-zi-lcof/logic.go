package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	var existsMap = make(map[int]struct{})
	for _, item := range nums {
		if _, ok := existsMap[item]; ok {
			return item
		}
		existsMap[item] = struct{}{}
	}
	return 0
}

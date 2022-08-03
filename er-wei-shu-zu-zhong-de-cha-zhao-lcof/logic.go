package main

import "fmt"

func main() {
	var matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(findNumberIn2DArray(matrix, 20))
}

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	if matrix[0][0] > target {
		return false
	}
	var i, j = 0, len(matrix[0]) - 1
	for {
		if i >= len(matrix) {
			break
		}
		if j < 0 {
			break
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		i++
		continue
	}
	return false
}

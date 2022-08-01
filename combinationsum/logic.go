package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

// https://leetcode.cn/problems/combination-sum/submissions/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	for index := range candidates {
		res = append(res, findCombination(candidates, target, index)...)
	}
	return res
}

func findCombination(candidates []int, target int, index int) [][]int {
	if candidates[index] > target {
		return [][]int{}
	}
	if candidates[index] == target {
		return [][]int{{target}}
	}
	var res [][]int
	for tmpIndex := index; tmpIndex < len(candidates); tmpIndex++ {
		subList := findCombination(candidates, target-candidates[index], tmpIndex)
		if len(subList) == 0 {
			continue
		}
		for _, item := range subList {
			var elem = []int{candidates[index]}
			elem = append(elem, item...)
			res = append(res, elem)
		}
	}
	return res
}

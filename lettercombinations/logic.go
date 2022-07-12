package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var ns = []string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	var targetList []string
	for _, target := range digits {
		index := target - '2'
		if index < 0 {
			continue
		}
		targetList = append(targetList, ns[index])
	}
	if len(targetList) == 0 {
		return []string{}
	}
	return combination(targetList)
}

func combination(targetList []string) []string {
	if len(targetList) == 0 {
		return []string{}
	}
	var res []string
	if len(targetList) == 1 {
		for _, item := range targetList[0] {
			res = append(res, string(item))
		}
	}
	for _, item := range targetList[0] {
		for _, elem := range combination(targetList[1:]) {
			res = append(res, string(item)+elem)
		}
	}
	return res
}

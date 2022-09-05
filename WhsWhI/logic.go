package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

// https://leetcode.cn/problems/WhsWhI/
func longestConsecutive(nums []int) int {
	var m = make(map[int]struct{})
	for _, item := range nums {
		m[item] = struct{}{}
	}
	var res, tmp int
	for k := range m {
		if _, ok := m[k-1]; ok {
			continue
		}
		kTmp := k
		tmp = 1
		for {
			if _, ok := m[kTmp+1]; !ok {
				break
			}
			kTmp += 1
			tmp += 1
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}

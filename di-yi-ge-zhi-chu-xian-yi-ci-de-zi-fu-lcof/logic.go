package main

import "fmt"

func main() {
	fmt.Println(string(firstUniqChar("abaccdeff")))
}

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func firstUniqChar(s string) byte {
	var existsMap = make(map[byte]bool)
	for _, item := range s {
		bItem := byte(item)
		exists, ok := existsMap[bItem]
		if !ok {
			existsMap[bItem] = true
			continue
		}
		if !exists {
			continue
		}
		existsMap[bItem] = false
	}
	for _, item := range s {
		if existsMap[byte(item)] {
			return byte(item)
		}
	}
	return byte(' ')
}

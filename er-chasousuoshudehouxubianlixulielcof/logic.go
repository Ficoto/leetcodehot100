package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
}

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	var (
		sub1, sub2 []int
		isBigger   bool
	)
	for i := 0; i < len(postorder)-1; i++ {
		if isBigger && postorder[i] < postorder[len(postorder)-1] {
			return false
		}
		if isBigger {
			sub2 = append(sub2, postorder[i])
			continue
		}
		if postorder[i] > postorder[len(postorder)-1] {
			sub2 = append(sub2, postorder[i])
			isBigger = true
			continue
		}
		sub1 = append(sub1, postorder[i])
	}
	return verifyPostorder(sub1) && verifyPostorder(sub2)
}

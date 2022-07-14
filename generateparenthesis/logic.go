package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var cache = [100][]string{}

// https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	generate(n)
	return cache[n]
}

func generate(n int) {
	if cache[n] != nil {
		return
	}
	if n == 0 {
		cache[0] = []string{""}
		return
	}
	var res []string
	for i := 0; i != n; i++ {
		generate(i)
		generate(n - i - 1)
		for _, left := range cache[i] {
			for _, right := range cache[n-i-1] {
				res = append(res, "("+left+")"+right)
			}
		}
	}
	cache[n] = res
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	var res strings.Builder
	for _, item := range s {
		if item != ' ' {
			res.WriteByte(byte(item))
			continue
		}
		res.WriteString("%20")
	}
	return res.String()
}

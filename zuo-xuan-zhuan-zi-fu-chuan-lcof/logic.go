package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}

func reverseLeftWords(s string, n int) string {
	if n > len(s) {
		return s
	}
	var res strings.Builder
	res.WriteString(s[n:])
	res.WriteString(s[:n])
	return res.String()
}

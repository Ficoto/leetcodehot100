package main

import "fmt"

func main() {
	fmt.Println(isValid("("))
}

type node struct {
	value int32
	prev  *node
}

type stack struct {
	head *node
}

func (s *stack) add(v int32) {
	var n = new(node)
	n.value = v
	n.prev = s.head
	s.head = n
}

func (s *stack) pop() *node {
	if s.head == nil {
		return s.head
	}
	res := s.head
	s.head = s.head.prev
	return res
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	var (
		m = map[int32]int32{
			')': '(',
			'}': '{',
			']': '[',
		}
		sta stack
	)
	for _, item := range s {
		target, ok := m[item]
		if !ok {
			sta.add(item)
			continue
		}
		n := sta.pop()
		if n != nil && n.value == target {
			continue
		}
		return false
	}
	return sta.head == nil
}

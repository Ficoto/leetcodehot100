package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

type stock struct {
	s []int
}

func (s *stock) Push(i int) {
	s.s = append(s.s, i)
}

func (s *stock) Pop() int {
	if len(s.s) == 0 {
		return -1
	}
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res
}

type CQueue struct {
	s1 *stock
	s2 *stock
}

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
func Constructor() CQueue {
	var res = CQueue{
		s1: new(stock),
		s2: new(stock),
	}
	return res
}

func (this *CQueue) AppendTail(value int) {
	for {
		oldValue := this.s1.Pop()
		if oldValue == -1 {
			break
		}
		this.s2.Push(oldValue)
	}
	this.s1.Push(value)
	for {
		oldValue := this.s2.Pop()
		if oldValue == -1 {
			break
		}
		this.s1.Push(oldValue)
	}
}

func (this *CQueue) DeleteHead() int {
	return this.s1.Pop()
}

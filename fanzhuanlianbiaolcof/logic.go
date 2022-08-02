package main

import "fmt"

func main() {
	var head = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 3
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	res := reverseList(head)
	for {
		if res == nil {
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		res *ListNode
		tmp = head
	)
	for {
		if tmp == nil {
			break
		}
		n := new(ListNode)
		n.Val = tmp.Val
		n.Next = res
		res = n
		tmp = tmp.Next
	}
	return res
}

package main

import "fmt"

func main() {
	var head = new(ListNode)
	head.Val = 1
	head.Next = new(ListNode)
	head.Next.Val = 3
	head.Next.Next = new(ListNode)
	head.Next.Next.Val = 2
	fmt.Println(reversePrint(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	subList := reversePrint(head.Next)
	return append(subList, head.Val)
}

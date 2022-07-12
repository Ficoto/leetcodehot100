package main

import "fmt"

func main() {
	var (
		head = new(ListNode)
		tmp  = head
	)
	for i, item := range []int{1} {
		if i == 0 {
			head.Val = item
			tmp = head
			continue
		}
		elem := new(ListNode)
		elem.Val = item
		tmp.Next = elem
		tmp = elem
	}
	head = removeNthFromEnd(head, 1)
	for {
		if head == nil {
			break
		}
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	num := removeNth(head, n)
	if num > 0 {
		return head.Next
	}
	return head
}

func removeNth(node *ListNode, n int) int {
	if n == 0 {
		return -1
	}
	if node == nil {
		return 0
	}
	subNth := removeNth(node.Next, n)
	if subNth < 0 {
		return subNth
	}
	if subNth < n {
		return subNth + 1
	}
	node.Next = node.Next.Next
	return -1
}

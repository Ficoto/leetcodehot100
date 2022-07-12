package main

import "fmt"

func main() {
	var (
		l1, l2 = new(ListNode), new(ListNode)
		tmp    *ListNode
	)
	for i, item := range []int{1, 2, 4} {
		if i == 0 {
			l1.Val = item
			tmp = l1
			continue
		}
		elem := new(ListNode)
		elem.Val = item
		tmp.Next = elem
		tmp = tmp.Next
	}
	for i, item := range []int{1, 3, 4} {
		if i == 0 {
			l2.Val = item
			tmp = l2
			continue
		}
		elem := new(ListNode)
		elem.Val = item
		tmp.Next = elem
		tmp = tmp.Next
	}
	res := mergeTwoLists(l1, l2)
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

// https://leetcode.cn/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res = new(ListNode)
	merge(list1, list2, res)
	return res.Next
}

func merge(l1, l2 *ListNode, res *ListNode) {
	if l1 == nil && l2 == nil {
		return
	}
	if l1 == nil {
		res.Next = l2
		return
	}
	if l2 == nil {
		res.Next = l1
		return
	}
	if l1.Val < l2.Val {
		elem := new(ListNode)
		elem.Val = l1.Val
		res.Next = elem
		merge(l1.Next, l2, res.Next)
		return
	}
	elem := new(ListNode)
	elem.Val = l2.Val
	res.Next = elem
	merge(l1, l2.Next, res.Next)
	return
}

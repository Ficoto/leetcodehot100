package mergeklists

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode.cn/problems/merge-k-sorted-lists/
func mergeKLists(list []*ListNode) *ListNode {
	var res *ListNode
	for _, node := range list {
		if node == nil {
			continue
		}
		if res == nil {
			res = new(ListNode)
			res.Val = node.Val
			merge(res, node.Next)
			continue
		}
		if res.Val > node.Val {
			var tmp = new(ListNode)
			tmp.Val = node.Val
			tmp.Next = res
			res = tmp
			merge(res, node.Next)
			continue
		}
		merge(res, node)
	}
	return res
}

func merge(n1, n2 *ListNode) {
	if n2 == nil {
		return
	}
	if n1 == nil {
		return
	}
	if n2.Val < n1.Val {
		var tmp = new(ListNode)
		tmp.Val = n2.Val
		tmp.Next = n1
		merge(tmp, n2)
		return
	}
	if n1.Next == nil {
		var tmp = new(ListNode)
		tmp.Val = n2.Val
		n1.Next = tmp
		merge(tmp, n2.Next)
		return
	}
	if n1.Val <= n2.Val && n2.Val < n1.Next.Val {
		var tmp = new(ListNode)
		tmp.Val = n2.Val
		tmp.Next = n1.Next
		n1.Next = tmp
		merge(tmp, n2.Next)
		return
	}
	merge(n1.Next, n2)
}
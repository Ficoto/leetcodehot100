package addtwonumbers

// https://leetcode.cn/problems/add-two-numbers/
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = new(ListNode)
	addTwoNode(l1, l2, false, res)
	return res.Next
}

func addTwoNode(l1 *ListNode, l2 *ListNode, isNeedAdd bool, res *ListNode) {
	if l1 == nil && l2 == nil {
		if !isNeedAdd {
			return
		}
		res.Next = new(ListNode)
		res.Next.Val = 1
		return
	}
	if l1 == nil {
		if isNeedAdd {
			l2.Val += 1
		}
		res.Next = l2
		if l2.Val <= 9 {
			return
		}
		l2.Val %= 10
		addTwoNode(l1, l2.Next, true, l2)
		return
	}
	if l2 == nil {
		if isNeedAdd {
			l1.Val += 1
		}
		res.Next = l1
		if l1.Val <= 9 {
			return
		}
		l1.Val %= 10
		addTwoNode(l1.Next, l2, true, l1)
		return
	}
	var newNode = new(ListNode)
	newNode.Val = l1.Val + l2.Val
	if isNeedAdd {
		newNode.Val += 1
	}
	res.Next = newNode
	isNeedAdd = newNode.Val >= 10
	if isNeedAdd {
		newNode.Val %= 10
	}
	addTwoNode(l1.Next, l2.Next, isNeedAdd, newNode)
}
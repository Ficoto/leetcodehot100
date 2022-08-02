package main

import "fmt"

func main() {
	var (
		t = new(Node)
	)
	t.Val = 7
	//tmp := t
	tmp := new(Node)
	tmp.Val = 13
	t.Next = tmp
	tmp.Random = t
	tmp.Next = new(Node)
	tmp = tmp.Next
	tmp.Val = 11
	n4 := new(Node)
	n4.Val = 1
	tmp.Random = n4
	tmp.Next = new(Node)
	tmp.Next.Random = tmp
	tmp = tmp.Next
	tmp.Val = 10
	tmp.Next = n4
	tmp = t
	for {
		if tmp == nil {
			break
		}
		fmt.Printf("this is %d\n", tmp.Val)
		if tmp.Random != nil {
			fmt.Printf("random is %d\n", tmp.Random.Val)
		}
		tmp = tmp.Next
	}
	fmt.Println("分割线=======")
	c := copyRandomList(t)
	for {
		if c == nil {
			break
		}
		fmt.Printf("this is %d\n", c.Val)
		if c.Random != nil {
			fmt.Printf("random is %d\n", c.Random.Val)
		}
		c = c.Next
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var (
		nodeMap = make(map[*Node]*Node)
		res     *Node
		tmp     *Node
	)
	for {
		if head == nil {
			break
		}
		if res == nil {
			res = new(Node)
			res.Val = head.Val
			if head.Random == head {
				res.Random = res
			} else if head.Random != nil {
				rNode := new(Node)
				rNode.Val = head.Random.Val
				nodeMap[head.Random] = rNode
				res.Random = rNode
			}
			tmp = res
			nodeMap[head] = res
			head = head.Next
			continue
		}
		if n, ok := nodeMap[head]; ok {
			tmp.Next = n
		} else {
			tmp.Next = new(Node)
			tmp.Next.Val = head.Val
			nodeMap[head] = tmp.Next
		}
		if n, ok := nodeMap[head.Random]; ok {
			tmp.Next.Random = n
		} else if head.Random != nil {
			tmp.Next.Random = new(Node)
			tmp.Next.Random.Val = head.Random.Val
			nodeMap[head.Random] = tmp.Next.Random
		}
		tmp = tmp.Next
		head = head.Next
	}
	return res
}

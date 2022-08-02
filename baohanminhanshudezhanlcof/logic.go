package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())
}

// https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
type Node struct {
	Value    int
	NextNode *Node
	LastNode *Node
}

type MinStack struct {
	s      []*Node
	header *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	var n = new(Node)
	n.Value = x
	this.s = append(this.s, n)
	if this.header == nil {
		this.header = n
		return
	}
	var tmpNode = this.header
	for {
		if tmpNode.Value > x {
			n.LastNode = tmpNode.LastNode
			n.NextNode = tmpNode
			tmpNode.LastNode = n
			if n.LastNode == nil {
				this.header = n
			} else {
				n.LastNode.NextNode = n
			}
			break
		}
		if tmpNode.NextNode == nil {
			tmpNode.NextNode = n
			n.LastNode = tmpNode
			break
		}
		tmpNode = tmpNode.NextNode
	}
}

func (this *MinStack) Pop() {
	if len(this.s) == 0 {
		return
	}
	popNode := this.s[len(this.s)-1]
	if popNode.LastNode != nil {
		popNode.LastNode.NextNode = popNode.NextNode
	} else {
		this.header = popNode.NextNode
	}
	if popNode.NextNode != nil {
		popNode.NextNode.LastNode = popNode.LastNode
	}
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	if len(this.s) == 0 {
		return 0
	}
	topNode := this.s[len(this.s)-1]
	return topNode.Value
}

func (this *MinStack) Min() int {
	if this.header == nil {
		return 0
	}
	return this.header.Value
}

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	findValue(true, &res, root)
	return res
}

func findValue(isLeft bool, res *[][]int, nodes ...*TreeNode) {
	if len(nodes) == 0 {
		return
	}
	var (
		elem     []int
		subNodes []*TreeNode
	)
	for _, item := range nodes {
		if item == nil {
			continue
		}
		if isLeft {
			elem = append(elem, item.Val)
		} else {
			elem = append([]int{item.Val}, elem...)
		}
		if item.Left != nil {
			subNodes = append(subNodes, item.Left)
		}
		if item.Right != nil {
			subNodes = append(subNodes, item.Right)
		}
	}
	if len(elem) == 0 {
		return
	}
	*res = append(*res, elem)
	findValue(!isLeft, res, subNodes...)
	return
}

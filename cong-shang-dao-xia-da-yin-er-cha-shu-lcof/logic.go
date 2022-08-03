package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func levelOrder(root *TreeNode) []int {
	return findValue(root)
}

func findValue(nodes ...*TreeNode) []int {
	if len(nodes) == 0 {
		return []int{}
	}
	var (
		subNodes []*TreeNode
		res      []int
	)
	for _, item := range nodes {
		if item == nil {
			continue
		}
		res = append(res, item.Val)
		if item.Left != nil {
			subNodes = append(subNodes, item.Left)
		}
		if item.Right != nil {
			subNodes = append(subNodes, item.Right)
		}
	}
	res = append(res, findValue(subNodes...)...)
	return res
}

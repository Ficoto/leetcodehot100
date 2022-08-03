package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func levelOrder(root *TreeNode) [][]int {
	return findValue(root)
}

func findValue(nodes ...*TreeNode) [][]int {
	var (
		res      [][]int
		subNodes []*TreeNode
	)
	for _, item := range nodes {
		if item == nil {
			continue
		}
		if len(res) == 0 {
			res = [][]int{{}}
		}
		res[0] = append(res[0], item.Val)
		if item.Left != nil {
			subNodes = append(subNodes, item.Left)
		}
		if item.Right != nil {
			subNodes = append(subNodes, item.Right)
		}
	}
	if len(subNodes) != 0 {
		res = append(res, findValue(subNodes...)...)
	}
	return res
}

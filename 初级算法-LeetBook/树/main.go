// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x28wnt/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if r > l {
		res += r
	} else {
		res += l
	}
	return res
}

func main() {
	fmt.Println(maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
}

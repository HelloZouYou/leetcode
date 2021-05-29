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
	// 深度优先搜索，nil节点返回0
	if root == nil {
		return 0
	}
	// 递归搜索当前节点的左右子节点，当前节点深度记1，将子节点的结果加起来
	res := 1
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	// 只取最大值
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

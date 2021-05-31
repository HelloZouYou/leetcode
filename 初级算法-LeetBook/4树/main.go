// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x28wnt/
package main

import (
	"fmt"
	"math"
)

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

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isValidBst(root, math.MinInt64, math.MaxInt64)
}
func isValidBst(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	// 递归判断左右子节点的值是否在限定范围内
	if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= min) {
		return false
	}
	if root.Right != nil && (root.Right.Val <= root.Val || root.Right.Val >= max) {
		return false
	}
	if !isValidBst(root.Left, min, root.Val) || !isValidBst(root.Right, root.Val, max) {
		return false
	}
	return true
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelp(root.Left, root.Right)
}
func isSymmetricHelp(left, right *TreeNode) bool {
	// 判断当前要比较的两个节点值是否一致
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	// 为保证对称，分别比较左节点的左子节点和右节点的右子节点，以及左节点的右子节点和右节点的左子节点
	return isSymmetricHelp(left.Left, right.Right) && isSymmetricHelp(left.Right, right.Left)
}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// 临时数组保存下一层要遍历的有效节点，先将根节点初始化进去
	arr := []*TreeNode{root}
	for len(arr) != 0 {
		tmp1 := make([]int, 0, len(arr))         // 记录该层的值
		tmp2 := make([]*TreeNode, 0, len(arr)*2) // 记录下一层的有效节点
		// 遍历当前层
		for _, item := range arr {
			tmp1 = append(tmp1, item.Val)
			if item.Left != nil {
				tmp2 = append(tmp2, item.Left)
			}
			if item.Right != nil {
				tmp2 = append(tmp2, item.Right)
			}
		}
		// 将当前层的值加入返回值
		res = append(res, tmp1)
		// 将一下层要遍历的节点赋给arr，没有时len为0会停止循环
		arr = tmp2
	}
	return res
}

// 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	return sortedArrayToBst(nums, 0, l)
}

func sortedArrayToBst(nums []int, start, end int) *TreeNode {
	// 二者相等说明这一半没元素了，返回nil
	if start >= end {
		return nil
	}
	// 二分法，递归处理
	mid := (start + end) / 2
	left := sortedArrayToBst(nums, start, mid)
	right := sortedArrayToBst(nums, mid+1, end)
	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}

func main() {
	fmt.Printf("%+v", sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}

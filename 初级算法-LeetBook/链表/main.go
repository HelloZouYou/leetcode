// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2t7vj/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表中的节点
func deleteNode(node *ListNode) {
	// 直接用下一个节点的数据覆盖当前节点，即可间接实现删除当前节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义快慢指针
	fast := head
	slow := head
	// 快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	// if head == nil {
	// 	return head
	// }
	// var pre *ListNode
	// for head.Next != nil {
	// 	head.Next, pre, head = pre, head, head.Next
	// }
	// head.Next = pre
	// return head

	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next            // 保存一个指向下一个节点的副本
	reverse := reverseList(next) // 递归到最末尾后，返回的reverse即是末尾节点
	next.Next = head             // 递归返回后，根据留在这的next，更改其Next指向自己
	head.Next = nil              // 最后头节点变尾节点需要清空Next避免构成环
	return reverse
}

func main() {
	fmt.Println(reverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}))
}

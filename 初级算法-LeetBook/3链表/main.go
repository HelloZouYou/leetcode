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

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 某一个链表到尾了则终止递归
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 车轮战，轮番比试
	// 谁当前值小，就留下，然后递归比较其Next
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 回文链表
func isPalindrome(head *ListNode) bool {
	fast := head
	var reverse *ListNode
	// 先通过快慢指针将当前head移到中间位置，并同时直接将前半部分翻转
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head.Next, reverse, head = reverse, head, head.Next
	}
	// fast不为nil说明奇数个，此时head恰在中间，需要再进一位
	if fast != nil {
		head = head.Next
	}
	// 此时head和reverse就分别是后半部分和前半部分了
	for head != nil {
		// 直接比较值是否一致
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	return true
}

// 环形链表
func hasCycle(head *ListNode) bool {
	// 快慢指针，最终相遇则有环
	// if head == nil || head.Next == nil {
	// 	return false
	// }
	// fast, slow := head, head
	// for fast != nil && fast.Next != nil {
	// 	fast = fast.Next.Next
	// 	slow = slow.Next
	// 	if fast == slow { // 要比较节点本身
	// 		return true
	// 	}
	// }
	// return false

	// 先将节点的Next指向自己同时head指针后移，如果节点Next已经指向自己，则说明有环
	if head == nil || head.Next == nil {
		return false
	}
	for head != nil && head.Next != nil {
		head.Next, head = head, head.Next
		if head.Next == head {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}))
}

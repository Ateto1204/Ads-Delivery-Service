/**
 * https://leetcode.com/problems/reverse-linked-list/description
 * Problem ID: 206. Reverse Linked List
 *
 * Author: Ateto
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	list := []int{}
	dummy := head
	for dummy != nil {
		list = append(list, dummy.Val)
		dummy = dummy.Next
	}
	n := len(list)
	if n <= 1 {
		return head
	}
	dummy = head
	for i := n - 1; i >= 0; i-- {
		dummy.Val = list[i]
		dummy = dummy.Next
	}
	return head
}
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
	var prev *ListNode
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}
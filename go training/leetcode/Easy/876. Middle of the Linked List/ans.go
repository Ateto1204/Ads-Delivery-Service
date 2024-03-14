/**
 * https://leetcode.com/problems/middle-of-the-linked-list/
 * Problem ID: 876. Middle of the Linked List
 *
 * Author: Ateto
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	tmp := head
	sz := 0
	for tmp != nil {
		tmp = tmp.Next
		sz++
	}

	for i := 0; i < sz/2; i++ {
		head = head.Next
	}

	return head
}
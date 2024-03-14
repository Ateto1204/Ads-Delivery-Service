/**
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/
 * Problem ID: 19. Remove Nth Node From End of List
 *
 * Author: Ateto
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := head
	size := 0
	for tmp != nil {
		tmp = tmp.Next
		size++
	}
	if size-n == 0 {
		return head.Next
	}
	tmp = head
	for i := 0; i < size-n-1; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return head
}
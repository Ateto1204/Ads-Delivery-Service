/**
 * https://leetcode.com/problems/merge-in-between-linked-lists
 * Problem ID: 1669. Merge In Between Linked Lists
 *
 * Author: Ateto
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	left := list1
	for i := 0; i < a-1; i++ {
		left = left.Next
	}
	right := left.Next
	for i := 0; i <= b-a; i++ {
		right = right.Next
	}
	left.Next = list2
	for left.Next != nil {
		left = left.Next
	}
	left.Next = right
	return list1
}
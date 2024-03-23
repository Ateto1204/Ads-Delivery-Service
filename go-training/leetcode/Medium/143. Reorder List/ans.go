/*
 *
 * https://leetcode.com/problems/reorder-list
 * Problem ID: 143. Reorder List
 *
 * Author: Ateto
 *
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	dummy := head
	arr := []int{}
	for dummy != nil {
		arr = append(arr, dummy.Val)
		dummy = dummy.Next
	}
	n := len(arr)
	if n <= 2 {
		return
	}
	dummy = head
	for i := 0; i < n/2; i++ {
		dummy.Val = arr[i]
		dummy = dummy.Next
		dummy.Val = arr[n-i-1]
		dummy = dummy.Next
	}
	if n%2 == 1 {
		dummy.Val = arr[n/2]
	}
}
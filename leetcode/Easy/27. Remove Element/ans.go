/**
 * https://leetcode.com/problems/remove-element/
 * Problem ID: 27. Remove Element
 *
 * Author: Ateto
 */

func removeElement(nums []int, val int) int {
	count := len(nums)
	for _, e := range nums {
		if e == val {
			count--
		}
	}
	var dummy []int
	for _, e := range nums {
		if e != val {
			dummy = append(dummy, e)
		}
	}
	for i, _ := range dummy {
		nums[i] = dummy[i]
	}
	return count
}
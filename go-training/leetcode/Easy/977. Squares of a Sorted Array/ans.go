/**
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/
 * Problem ID: 977. Squares of a Sorted Array
 *
 * Author: Ateto
 *
 */

func sortedSquares(nums []int) []int {
	for idx, num := range nums {
		num *= num
		nums[idx] = num
	}
	sort.Ints(nums)
	return nums
}
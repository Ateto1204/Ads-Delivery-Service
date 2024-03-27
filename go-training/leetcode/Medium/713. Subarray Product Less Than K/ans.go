/**
 * https://leetcode.com/problems/subarray-product-less-than-k
 * Problem ID: 713. Subarray Product Less Than K
 *
 * Author: Ateto
 */

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	left := 0
	right := 0
	product := 1
	count := 0
	n := len(nums)
	for right < n {
		product *= nums[right]
		for product >= k {
			product /= nums[left]
			left++
		}
		count += right - left + 1
		right++
	}
	return count
}
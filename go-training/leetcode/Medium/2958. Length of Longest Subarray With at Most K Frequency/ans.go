/**
 * https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency
 * Problem ID: 2958. Length of Longest Subarray With at Most K Frequency
 *
 * Author: Ateto
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubarrayLength(nums []int, k int) int {
	mp := make(map[int]int)
	n := len(nums)
	left := 0
	right := 0
	result := 0
	for right < n {
		cur := nums[right]
		right++
		mp[cur]++
		for left < right && mp[cur] > k {
			dummy := nums[left]
			left++
			mp[dummy]--
		}
		result = max(result, right-left)
	}
	return result
}
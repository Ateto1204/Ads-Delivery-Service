/**
 * https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times
 * Problem ID: 2962. Count Subarrays Where Max Element Appears at Least K Times
 *
 * Author: Ateto
 */

func countSubarrays(nums []int, k int) int64 {
	mx := 0
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	left := 0
	right := 0
	result := int64(0)
	n := len(nums)
	count := 0
	for right < n || left > right {
		if nums[right] == mx {
			count++
		}
		for count >= k {
			if nums[left] == mx {
				count--
			}
			result += int64(n - right)
			left++
		}
		right++
	}
	return result
}
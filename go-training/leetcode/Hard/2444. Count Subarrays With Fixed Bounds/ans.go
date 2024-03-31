/**
 * https://leetcode.com/problems/count-subarrays-with-fixed-bounds
 * Problem ID: 2444. Count Subarrays With Fixed Bounds
 *
 * Author: Ateto
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	p := -1
	q := -1
	dummy := -1
	result := int64(0)
	for i, num := range nums {
		if (minK <= num && num <= maxK) == false {
			dummy = i
		}
		if num == minK {
			p = i
		}
		if num == maxK {
			q = i
		}
		tmp := int64(max(0, min(p, q)-dummy))
		result += tmp
	}
	return result
}
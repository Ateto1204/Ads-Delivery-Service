/**
 * https://leetcode.com/problems/subarrays-with-k-different-integers
 * Problem ID: 992. Subarrays with K Different Integers
 *
 * Author: Ateto
 */

func solve(nums []int, k int) int {
	mp := make(map[int]int)
	count := 0
	p := 0
	q := 0
	n := len(nums)
	result := 0
	for q < n {
		tmp := nums[q]
		mp[tmp]++
		if mp[tmp] == 1 {
			count++
		}
		for count > k {
			tmp := nums[p]
			mp[tmp]--
			if mp[tmp] == 0 {
				count--
			}
			p++
		}
		result += q - p + 1
		q++
	}
	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	x := solve(nums, k)
	y := solve(nums, k-1)
	result := x - y
	return result
}
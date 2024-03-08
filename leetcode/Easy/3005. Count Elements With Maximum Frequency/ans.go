/**
 * https://leetcode.com/problems/count-elements-with-maximum-frequency/
 * Problem ID: 3005. Count Elements With Maximum Frequency
 *
 * Author: Ateto
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFrequencyElements(nums []int) int {
	var cnt [105]int
	mx := 0
	i := 0
	sz := len(nums)
	for i = 0; i < sz; i++ {
		cnt[nums[i]]++
	}
	for i = 1; i <= 100; i++ {
		mx = max(mx, cnt[i])
	}
	ans := 0
	for i = 1; i <= 100; i++ {
		if cnt[i] == mx {
			ans += mx
		}
	}
	return ans
}
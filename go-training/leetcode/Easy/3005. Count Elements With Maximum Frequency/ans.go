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
	for _, num := range nums {
		cnt[num]++
	}
	for _, count := range cnt {
		mx = max(mx, count)
	}
	ans := 0
	for _, count := range cnt {
		if count == mx {
			ans += mx
		}
	}
	return ans
}
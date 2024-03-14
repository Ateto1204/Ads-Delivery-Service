/**
 *
 * https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
 * Problem ID: 1750. Minimum Length of String After Deleting Similar Ends
 *
 * Author: Ateto
 *
 */

func minimumLength(s string) int {
	left := 0
	right := len(s) - 1
	for left < right && s[left] == s[right] {
		tmp := s[left]
		for s[left+1] == tmp && left+1 < right {
			left++
		}
		left++
		for s[right-1] == tmp && left < right-1 {
			right--
		}
		right--
	}
	if right < left {
		return 0
	}
	return right - left + 1
}
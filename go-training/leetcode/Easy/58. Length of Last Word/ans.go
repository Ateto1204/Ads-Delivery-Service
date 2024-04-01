/*
 *
 * https://leetcode.com/problems/length-of-last-word
 * Problem ID: 58. Length of Last Word
 *
 * Author: Ateto
 *
 */

func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	j := i - 1
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}
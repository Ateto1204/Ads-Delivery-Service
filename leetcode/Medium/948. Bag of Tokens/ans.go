/**
 *
 * https://leetcode.com/problems/bag-of-tokens/
 * Problem ID: 948. Bag of Tokens
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

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	n := len(tokens)
	score := 0
	maxScore := 0
	left := 0
	right := n - 1

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			score++
			left++
			maxScore = max(maxScore, score)
		} else if score > 0 {
			power += tokens[right]
			score--
			right--
		} else {
			break
		}
	}

	return maxScore
}
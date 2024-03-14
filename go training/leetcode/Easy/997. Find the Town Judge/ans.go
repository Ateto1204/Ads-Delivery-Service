/**
 *
 * https://leetcode.com/problems/find-the-town-judge/
 * Problem ID: 997. Find the Town Judge
 *
 * Author: Ateto
 *
 */

func findJudge(n int, trust [][]int) int {
	var hasTrusted [1005]bool
	var beTrusted [1005]int

	var i int
	sz := len(trust)

	for i = 0; i < sz; i++ {
		hasTrusted[trust[i][0]] = true
		beTrusted[trust[i][1]]++
	}

	for i = 1; i <= n; i++ {
		if !hasTrusted[i] && beTrusted[i] == n-1 {
			return i
		}
	}
	return -1
}
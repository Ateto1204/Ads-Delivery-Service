/**
 * https://leetcode.com/problems/unique-paths
 * Problem ID: 62. Unique Paths
 *
 * Author: Ateto
 */

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
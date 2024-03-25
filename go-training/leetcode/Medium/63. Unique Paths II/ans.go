/**
 * https://leetcode.com/problems/unique-paths-ii
 * Problem ID: 63. Unique Paths II
 *
 * Author: Ateto
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width := len(obstacleGrid[0])
	dp := make([]int, width)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for j := 0; j < width; j++ {
			if row[j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[width-1]
}
/*
 *
 * https://leetcode.com/problems/find-the-duplicate-number
 * Problem ID: 287. Find the Duplicate Number
 *
 * Author: Ateto
 *
 */

func findDuplicate(nums []int) int {
	used := make(map[int]bool)
	for _, num := range nums {
		if used[num] {
			return num
		}
		used[num] = true
	}
	return -1
}
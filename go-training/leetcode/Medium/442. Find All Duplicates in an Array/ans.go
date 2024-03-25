/*
 *
 * https://leetcode.com/problems/find-all-duplicates-in-an-array
 * Problem ID: 442. Find All Duplicates in an Array
 *
 * Author: Ateto
 *
 */

func findDuplicates(nums []int) []int {
	used := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if used[num] {
			result = append(result, num)
		}
		used[num] = true
	}
	return result
}
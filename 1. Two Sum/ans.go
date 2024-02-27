/**
 * https://leetcode.com/problems/two-sum/
 * Problem ID: 1. Two Sum
 *
 * Author: Ateto
 */

func twoSum(nums []int, target int) []int {
	var result []int
	sz := len(nums)
	for i := 0; i < sz; i++ {
		for j := i + 1; j < sz; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}
	return result
}
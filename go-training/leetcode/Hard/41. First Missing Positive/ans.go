/**
 * https://leetcode.com/problems/first-missing-positive
 * Problem ID: 41. First Missing Positive
 *
 * Author: Ateto
 */

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	arr := []int{}
	for _, num := range nums {
		if num > 0 {
			arr = append(arr, num)
		}
	}
	n := len(arr)
	if n < 1 {
		return 1
	}
	first := true
	prev := 0
	for _, num := range arr {
		if first {
			first = false
			if num > 1 {
				return 1
			}
			prev = num
			continue
		}
		cur := num
		if cur-prev > 1 {
			return prev + 1
		}
		prev = cur
	}
	return prev + 1
}
/**
 * https://leetcode.com/problems/intersection-of-two-arrays/
 * Problem ID: 349. Intersection of Two Arrays
 *
 * Author: Ateto
 */

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	var cnt1 [1001]int
	var cnt2 [1001]int
	for _, val := range nums1 {
		cnt1[val]++
	}
	for _, val := range nums2 {
		cnt2[val]++
	}
	for i := 0; i <= 1000; i++ {
		if cnt1[i] > 0 && cnt2[i] > 0 {
			result = append(result, i)
		}
	}

	return result
}
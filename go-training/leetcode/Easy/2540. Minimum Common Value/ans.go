/**
 * https://leetcode.com/problems/minimum-common-value/
 * Problem ID: 2540. Minimum Common Value
 *
 * Author: Ateto
 *
 */

func getCommon(nums1 []int, nums2 []int) int {
	x := len(nums1)
	y := len(nums2)
	i := 0
	j := 0
	for i < x && j < y {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
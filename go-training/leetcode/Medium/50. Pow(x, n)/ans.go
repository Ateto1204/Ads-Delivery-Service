/**
 * https://leetcode.com/problems/intersection-of-two-arrays/
 * Problem ID: 349. Intersection of Two Arrays
 *
 * Author: Ateto
 */

func myPow(x float64, n int) float64 {
	sum := 1.0
	i := x
	j := n
	if j < 0 {
		j *= -1
	}
	for j > 0 {
		if (j & 1) == 1 {
			sum *= i
		}
		i *= i
		j >>= 1
	}
	if n < 0 {
		return 1.0 / sum
	}
	return sum
}
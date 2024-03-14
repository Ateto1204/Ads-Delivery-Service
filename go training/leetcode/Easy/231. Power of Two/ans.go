/**
 *
 * https://leetcode.com/problems/power-of-two/
 * Problem ID: 231. Power of Two
 *
 * Author: Ateto
 *
 */

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	tmp := n
	for tmp > 0 {
		if tmp == 1 {
			return true
		}
		if tmp%2 != 0 {
			return false
		}
		tmp /= 2
	}
	return true
}
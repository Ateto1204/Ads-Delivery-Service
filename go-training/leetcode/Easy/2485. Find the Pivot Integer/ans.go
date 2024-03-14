/**
 * https://leetcode.com/problems/find-the-pivot-integer/
 * Problem ID: 2485. Find the Pivot Integer
 *
 * Author: Ateto
 */

func cal(p, q int) int {
	tmp := (p + q) * (q - p + 1) / 2
	return tmp
}

func pivotInteger(n int) int {
	for i := n; i >= 0; i-- {
		x := cal(1, i)
		y := cal(i, n)
		if x == y {
			return i
		}
		if x < y {
			break
		}
	}
	return -1
}
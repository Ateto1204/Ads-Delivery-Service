/**
 *
 * https://leetcode.com/problems/maximum-odd-binary-number/
 * Problem ID: 2864. Maximum Odd Binary Number
 *
 * Author: Ateto
 *
 */

func maximumOddBinaryNumber(s string) string {
	cnt := 0
	sz := len(s)
	result := ""
	for i := 0; i < sz; i++ {
		if s[i] == '1' {
			cnt += 1
		}
	}
	for i := 0; i < cnt-1; i++ {
		result += "1"
	}
	for i := 0; i < sz-cnt; i++ {
		result += "0"
	}
	result += "1"
	return result
}
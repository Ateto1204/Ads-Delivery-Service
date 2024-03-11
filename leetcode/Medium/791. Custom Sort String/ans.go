/**
 * https://leetcode.com/problems/custom-sort-string/
 * Problem ID: 791. Custom Sort String
 *
 * Author: Ateto
 */

func customSortString(order string, s string) string {
	dic := map[rune]int{}

	for _, char := range s {
		dic[char]++
	}

	res := ""
	for _, char := range order {
		if _, ok := dic[char]; ok {
			res += strings.Repeat(string(char), dic[char])
			dic[char] = 0
		}
	}

	for key, value := range dic {
		if value != 0 {
			res += strings.Repeat(string(key), value)
		}
	}
	return res
}
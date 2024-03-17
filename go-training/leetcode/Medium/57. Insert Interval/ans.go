/**
 * https://leetcode.com/problems/insert-interval
 * Problem ID: 57. Insert Interval
 *
 * Author: Ateto
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	n := len(intervals)
	if n == 0 {
		result = append(result, newInterval)
		return result
	}
	p := intervals[0][0]
	q := intervals[0][1]
	used := false

	if newInterval[1] <= q {
		if newInterval[1] >= p {
			p = min(p, newInterval[0])
		} else {
			dummy := []int{newInterval[0], newInterval[1]}
			result = append(result, dummy)
		}
		used = true
	} else if newInterval[0] <= p {
		p = min(p, newInterval[0])
		q = max(q, newInterval[1])
		used = true
	}

	for i := 1; i < n; i++ {
		if !used {
			if newInterval[0] >= p && newInterval[1] <= intervals[i][1] {
				if newInterval[0] > q {
					dummy := []int{p, q}
					result = append(result, dummy)
					if newInterval[1] < intervals[i][0] {
						dummy := []int{newInterval[0], newInterval[1]}
						result = append(result, dummy)
						p = intervals[i][0]
					} else {
						p = newInterval[0]
					}
					q = intervals[i][1]
				} else {
					if newInterval[1] < intervals[i][0] {
						dummy := []int{p, newInterval[1]}
						result = append(result, dummy)
						p = intervals[i][0]
					}
					q = intervals[i][0]
				}
				used = true
			} else if newInterval[0] <= q {
				p = min(p, newInterval[0])
				q = max(q, newInterval[1])
				used = true
			}
		}
		if intervals[i][0] <= q {
			p = min(p, intervals[i][0])
			q = max(q, intervals[i][1])
		} else {
			dummy := []int{p, q}
			result = append(result, dummy)
			p = intervals[i][0]
			q = intervals[i][1]
		}
	}
	if !used {
		if newInterval[0] <= q {
			p = min(p, newInterval[0])
			q = max(q, newInterval[1])
		} else {
			dummy := []int{p, q}
			result = append(result, dummy)
			p = newInterval[0]
			q = newInterval[1]
		}
	}
	dummy := []int{p, q}
	result = append(result, dummy)
	return result
}
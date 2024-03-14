/**
 * https://leetcode.com/problems/find-bottom-left-tree-value/
 * Problem ID: 513. Find Bottom Left Tree Value
 *
 * Author: Ateto
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var leftmost, maxDepth int

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if depth > maxDepth {
			maxDepth = depth
			leftmost = root.Val
		}
	}
	dfs(root.Left, depth+1)
	dfs(root.Right, depth+1)
}

func findBottomLeftValue(root *TreeNode) int {
	leftmost = root.Val
	maxDepth = 0
	dfs(root, 0)
	return leftmost
}
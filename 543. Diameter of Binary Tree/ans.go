/**
 * https://leetcode.com/problems/diameter-of-binary-tree/
 * Problem ID: 543. Diameter of Binary Tree
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, diameter)
	right := dfs(root.Right, diameter)
	*diameter = max(*diameter, left+right)
	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	dfs(root, &diameter)
	return diameter
}
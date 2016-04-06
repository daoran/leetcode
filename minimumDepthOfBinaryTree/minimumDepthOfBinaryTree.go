// Source : https://oj.leetcode.com/problems/minimum-depth-of-binary-tree/
// Author : daoran
// Date   : 2016-04-06

/**********************************************************************************
*
* Given a binary tree, find its minimum depth.
*
* The minimum depth is the number of nodes along the shortest path from the root node
* down to the nearest leaf node.
*
**********************************************************************************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
import "math"
 
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := math.MaxInt32
	right := math.MaxInt32

	if root.Left != nil {
		left = minDepth(root.Left) + 1
	}
	if root.Right != nil {
		right = minDepth(root.Right) + 1
	}

	if left < right {
		return int(left)
	}

	return int(right)
}
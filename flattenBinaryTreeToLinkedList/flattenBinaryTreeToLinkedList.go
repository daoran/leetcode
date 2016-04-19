// Source : https://oj.leetcode.com/problems/flatten-binary-tree-to-linked-list/
// Author : daoran
// Date   : 2016-04-19

/**********************************************************************************
*
* Given a binary tree, flatten it to a linked list in-place.
*
* For example,
* Given
*
*          1
*         / \
*        2   5
*       / \   \
*      3   4   6
*
* The flattened tree should look like:
*
*    1
*     \
*      2
*       \
*        3
*         \
*          4
*           \
*            5
*             \
*              6
*
*
* Hints:
* If you notice carefully in the flattened tree, each node's right child points to
* the next node of a pre-order traversal.
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
package main

// TreeNode defines a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	v := []*TreeNode{}
	stack := []*TreeNode{}

	node := new(TreeNode)
	stack = append(stack, root)

	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		v = append(v, node)

		if node != nil && node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node != nil && node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	v = append(v, nil)
	for i := 0; i < len(v); i++ {
		if v[i] != nil {
			v[i].Left = nil
			v[i].Right = v[i+1]
		}
	}
}

func main() {

}

// Source : https://oj.leetcode.com/problems/balanced-binary-tree/
// Author : daoran
// Date   : 2016-04-10

/********************************************************************************** 
* 
* Given a binary tree, determine if it is height-balanced.
* 
* For this problem, a height-balanced binary tree is defined as a binary tree in which 
* the depth of the two subtrees of every node never differ by more than 1.
* 
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
func isBalanced(root *TreeNode) bool {
    return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    
    l := maxDepth(root.Left)
    if l == -1 {
        return -1
    }
    
    r := maxDepth(root.Right)
    if r == -1 {
        return -1
    }
    
    if abs(l - r) <= 1 {
        return max(l, r) + 1
    } 
    return -1
} 

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    
    if x == 0 {
        return 0
    }
    
    return x
}

func max(l,r int) int {
    if l > r {
        return l
    } 
    
    return r
}
// Source : https://leetcode.com/problems/invert-binary-tree/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
 * 
 * Invert a binary tree.
 *      4
 *    /   \
 *   2     7
 *  / \   / \
 * 1   3 6   9
 * 
 * to
 *      4
 *    /   \
 *   7     2
 *  / \   / \
 * 9   6 3   1
 * 
 * Trivia:
 * This problem was inspired by this original tweet by Max Howell:
 * (https://twitter.com/mxcl/status/608682016205344768)
 *
 *  | Google: 90% of our engineers use the software you wrote (Homebrew), 
 *  | but you can’t invert a binary tree on a whiteboard so fuck off.
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
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    
    node := invertTree(root.Left) 
    root.Left = invertTree(root.Right)
    
    root.Right = node 
    return root
}

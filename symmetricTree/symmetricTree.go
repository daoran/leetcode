// Source : https://oj.leetcode.com/problems/symmetric-tree/
// Author : daoran
// Date   : 2016-04-14

/********************************************************************************** 
* 
* Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
* 
* For example, this binary tree is symmetric:
* 
*     1
*    / \
*   2   2
*  / \ / \
* 3  4 4  3
* 
* But the following is not:
* 
*     1
*    / \
*   2   2
*    \   \
*    3    3
* 
* Note:
* Bonus points if you could solve it both recursively and iteratively.
* 
* confused what "{1,#,2,3}" means? > read more on how binary tree is serialized on OJ.
* 
* OJ's Binary Tree Serialization:
* 
* The serialization of a binary tree follows a level order traversal, where '#' signifies 
* a path terminator where no node exists below.
* 
* Here's an example:
* 
*    1
*   / \
*  2   3
*     /
*    4
*     \
*      5
* 
* The above binary tree is serialized as "{1,2,3,#,#,4,#,#,5}". 
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
 
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(p *TreeNode, q * TreeNode) bool {
    if p == nil && q == nil {
        return true 
    }
    
    if p == nil || q == nil {
        return false
    }
    
    return p.Val == q.Val && isSymmetricHelper(p.Left, q.Right) && 
                        isSymmetricHelper(p.Right, q.Left)
}
// Source : https://oj.leetcode.com/problems/validate-binary-search-tree/
// Author : Hao Chen
// Date   : 2014-07-05

/********************************************************************************** 
* 
* Given a binary tree, determine if it is a valid binary search tree (BST).
* 
* Assume a BST is defined as follows:
* 
* The left subtree of a node contains only nodes with keys less than the node's key.
* The right subtree of a node contains only nodes with keys greater than the node's key.
* Both the left and right subtrees must also be binary search trees.
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

var prev *TreeNode

func isValidBST(root *TreeNode) bool {
    prev = nil 
    return isMonotonicIncreasing(root)
}

func isMonotonicIncreasing(p *TreeNode) bool {
    if p == nil {
        return true
    }
    
    if isMonotonicIncreasing(p.Left) {
        if prev != nil && p.Val <= prev.Val {
            return false
        } 
        
        prev = p
        
        return isMonotonicIncreasing(p.Right)
    }
    return false
}
// Source : https://leetcode.com/problems/binary-tree-paths/
// Author : daoran
// Date   : 2016-04-14

/*************************************************************************************** 
 *
 * Given a binary tree, return all root-to-leaf paths.
 * 
 * For example, given the following binary tree:
 * 
 *    1
 *  /   \
 * 2     3
 *  \
 *   5
 * 
 * All root-to-leaf paths are:
 * ["1->2->5", "1->3"]
 * 
 * Credits:
 * Special thanks to @jianchao.li.fighter for adding this problem and creating all test 
 * cases.
 *               
 ***************************************************************************************/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
import "strconv" 

func binaryTreePathsHelper(root *TreeNode, path []int, result *[]string) {
    if root == nil {
        return 
    }
    
    path = append(path, root.Val)
    
    if root.Left == nil && root.Right == nil {
        if len(path) > 0 {
            s := []string{}
            for i := 0; i < len(path); i++ {
                s = append(s, strconv.Itoa(path[i]))
            }
            *result = append(*result, strings.Join(s, "->"))
        }
        return 
    }
    
    binaryTreePathsHelper(root.Left, path, result)
    binaryTreePathsHelper(root.Right, path, result)
} 

func binaryTreePaths(root *TreeNode) []string {
    result := &[]string{}
    path := []int{}
    binaryTreePathsHelper(root, path, result)
    return *result
}
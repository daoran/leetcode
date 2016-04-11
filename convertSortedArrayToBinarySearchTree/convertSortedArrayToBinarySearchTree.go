// Source : https://oj.leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Author : daoran
// Date   : 2016-04-11

/********************************************************************************** 
* 
* Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
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
func sortedArrayToBST(nums []int) *TreeNode {
    return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(arr int[], start, end int) *TreeNode {
    if start > end {
        return nil
    }
    
    mid := (start + end) / 2
    node := &TreeNode{Val:arr[mid]}
    node.Left = sortedArrayToBSTHelper(arr, start, mid-1)
    node.Right = sortedArrayToBSTHelper(arr, mid+1, end)
    return node
}
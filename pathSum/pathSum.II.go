/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func pathSum(root *TreeNode, sum int) [][]int {
    result := &[][]int{}
    path := []int{}
    
    generatePathSum(root, sum, path, result)
    return *result
}

func generatePathSum(root *TreeNode, sum int, path []int, result *[][]int) []int {
    if root == nil {
        return path
    }
    
    path = append(path, root.Val)

    if root.Left == nil && root.Right == nil {
        if root.Val == sum {
            *result = append(*result, path)
        }
    }
    
    path = generatePathSum(root.Left, sum - root.Val, path, result)
    path = generatePathSum(root.Right, sum - root.Val, path, result)
    
    temp := make([]int, len(path) - 1)
    for i := 0; i < len(path) - 1; i++ {
        temp[i] = path[i]
    }

    path = temp
    return path
}
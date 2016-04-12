// Source : https://oj.leetcode.com/problems/unique-paths-ii/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Follow up for "Unique Paths":
* 
* Now consider if some obstacles are added to the grids. How many unique paths would there be?
* 
* An obstacle and empty space is marked as 1 and 0 respectively in the grid.
* 
* For example,
* There is one obstacle in the middle of a 3x3 grid as illustrated below.
* 
* [
*   [0,0,0],
*   [0,1,0],
*   [0,0,0]
* ]
* 
* The total number of unique paths is 2.
* 
* Note: m and n will be at most 100.
*               
**********************************************************************************/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    
    if m == 0 {
        return 0
    }
    
    n := len(obstacleGrid[0])
    
    mat := make([][]int, m+1)    
    for i := 0; i < m+1; i++ {
        mat[i] = make([]int, n+1)
    }
    
    mat[m-1][n] = 1
    
    for r := m-1; r >= 0; r-- {
        for c := n-1; c >= 0; c-- {
            if obstacleGrid[r][c] == 1 {
                mat[r][c] = 0 
            } else {
                mat[r][c] = mat[r+1][c] + mat[r][c+1]
            }
        }
    }
    
    return mat[0][0]
}

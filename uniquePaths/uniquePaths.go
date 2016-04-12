// Source : https://oj.leetcode.com/problems/unique-paths/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
 * 
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
 * 
 * The robot can only move either down or right at any point in time. The robot is trying to reach 
 * the bottom-right corner of the grid (marked 'Finish' in the diagram below).
 *    
 *    
 *    start                                                  
 *    +---------+----+----+----+----+----+                   
 *    |----|    |    |    |    |    |    |                   
 *    |----|    |    |    |    |    |    |                   
 *    +----------------------------------+                   
 *    |    |    |    |    |    |    |    |                   
 *    |    |    |    |    |    |    |    |                   
 *    +----------------------------------+                   
 *    |    |    |    |    |    |    |----|                   
 *    |    |    |    |    |    |    |----|                   
 *    +----+----+----+----+----+---------+                   
 *                                   finish                  
 *    
 * 
 * How many possible unique paths are there?
 * 
 * Above is a 3 x 7 grid. How many possible unique paths are there?
 * 
 * Note: m and n will be at most 100.
 *               
 **********************************************************************************/

func uniquePaths(m int, n int) int {
    mat := make([][]int, m+1)    
    for i := 0; i < m+1; i++ {
        mat[i] = make([]int, n+1)
    }
    
    mat[m-1][n] = 1
    
    for r := m-1; r >= 0; r-- {
        for c := n-1; c >= 0; c-- {
            mat[r][c] = mat[r+1][c] + mat[r][c+1]
        }
    }
    
    return mat[0][0]
}


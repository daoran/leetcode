// Source : https://oj.leetcode.com/problems/spiral-matrix/
// Author : daoran
// Date   : 2016-04-11

/********************************************************************************** 
* 
* Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
* 
* For example,
* Given the following matrix:
* 
* [
*  [ 1, 2, 3 ],
*  [ 4, 5, 6 ],
*  [ 7, 8, 9 ]
* ]
* 
* You should return [1,2,3,6,9,8,7,4,5].
* 
*               
**********************************************************************************/

func spiralOrder(matrix [][]int) []int {
    result := []int{}
    
    if len(matrix) == 0 {
        return result
    }
    
    m := len(matrix)
    n := len(matrix[0])
    
    row := 0
    col := -1
    
    for true {
        for i := 0; i < n; i++ {
            col++
            result = append(result, matrix[row][col])
        }
        
        m-- 
        if m == 0 {
            break
        }
        for i := 0; i < m; i++ {
            row++
            result = append(result, matrix[row][col])
        }
        
        n--
        if n == 0 {
            break
        }
        for i := 0; i < n; i++ {
            col--
            result = append(result, matrix[row][col])
        }
        
        m-- 
        if m == 0 {
            break
        }
        for i := 0; i < m; i++ {
            row--
            result = append(result, matrix[row][col])
        }
        
        n-- 
        if n == 0 {
            break
        }
    }
    
    return result
}

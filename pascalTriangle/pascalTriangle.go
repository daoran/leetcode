// Source : https://oj.leetcode.com/problems/pascals-triangle/
// Author : daoran
// Date   : 2016-04-14

/********************************************************************************** 
* 
* Given numRows, generate the first numRows of Pascal's triangle.
* 
* For example, given numRows = 5,
* Return
* 
* [
*      [1],
*     [1,1],
*    [1,2,1],
*   [1,3,3,1],
*  [1,4,6,4,1]
* ]
* 
*               
**********************************************************************************/

func generate(numRows int) [][]int {
    result := [][]int{}
    
    for i := 0; i < numRows; i++ {
        v := []int{}
        if i == 0 {
            v = append(v, 1)
        } else {
            v = append(v, 1)
            for j := 0; j < len(result[i-1])-1; j++ {
                v = append(v, result[i-1][j]+result[i-1][j+1])
            }
            v = append(v,1)
         }
         result = append(result, v)
    }
    return result
}
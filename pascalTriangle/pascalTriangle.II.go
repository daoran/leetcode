// Source : https://oj.leetcode.com/problems/pascals-triangle-ii/
// Author : daoran
// Date   : 2016-04-14

/********************************************************************************** 
* 
* Given an index k, return the kth row of the Pascal's triangle.
* 
* For example, given k = 3,
* Return [1,3,3,1].
* 
* Note:
* Could you optimize your algorithm to use only O(k) extra space?
* 
*               
**********************************************************************************/

func getRow(rowIndex int) []int {
    v := make([]int, rowIndex+1)
    
    v[0] = 1
    
    for i := 0; i < rowIndex; i++ {
        for j := i + 1; j > 0; j-- {
            v[j] += v[j-1]
        }
    }
    
    return v
}
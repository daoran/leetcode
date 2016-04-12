// Source : https://oj.leetcode.com/problems/search-insert-position/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Given a sorted array and a target value, return the index if the target is found. 
* If not, return the index where it would be if it were inserted in order.
* 
* You may assume no duplicates in the array.
* 
* Here are few examples.
* [1,3,5,6], 5 → 2
* [1,3,5,6], 2 → 1
* [1,3,5,6], 7 → 4
* [1,3,5,6], 0 → 0
* 
*               
**********************************************************************************/

func searchInsert(nums []int, target int) int {
    L := 0
    R := len(nums) - 1
    for L < R {
        M := (L + R) / 2
        if nums[M] < target {
            L = M + 1
        } else {
            R = M
        }
    }
    
    if nums[L] < target {
        return L + 1
    }
    
    return L
}



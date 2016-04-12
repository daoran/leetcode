// Source : https://oj.leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Follow up for "Find Minimum in Rotated Sorted Array":
* What if duplicates are allowed?
* 
* Would this affect the run-time complexity? How and why?
* 
* Suppose a sorted array is rotated at some pivot unknown to you beforehand.
* 
* (i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
* 
* Find the minimum element.
* 
* The array may contain duplicates.
*               
**********************************************************************************/

func findMin(nums []int) int {
    L := 0
    R := len(nums) - 1
    
    for L < R && nums[L] >= nums[R] {
        M := (L + R) / 2
        if nums[M] > nums[R] {
            L = M + 1
        } else if nums[M] < nums[L] {
            R = M
        } else {
            L = L + 1
        }
    }
    
    return nums[L]
}

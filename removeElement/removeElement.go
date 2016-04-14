// Source : https://oj.leetcode.com/problems/remove-element/
// Author : daoran
// Date   : 2016-04-14

/********************************************************************************** 
* 
* Given an array and a value, remove all instances of that value in place and return the new length.
* 
* The order of elements can be changed. It doesn't matter what you leave beyond the new length.
* 
*               
**********************************************************************************/

func removeElement(nums []int, val int) int {
    pos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[pos] = nums[i]
            pos++
        }
    }
    
    return pos
}
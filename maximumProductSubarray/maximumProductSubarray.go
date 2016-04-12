// Source : https://oj.leetcode.com/problems/maximum-product-subarray/
// Author : daoran
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Find the contiguous subarray within an array (containing at least one number) 
* which has the largest product.
* 
* For example, given the array [2,3,-2,4],
* the contiguous subarray [2,3] has the largest product = 6.
* 
* More examples:
*   
*   Input: arr[] = {6, -3, -10, 0, 2}
*   Output:   180  // The subarray is {6, -3, -10}
*   
*   Input: arr[] = {-1, -3, -10, 0, 60}
*   Output:   60  // The subarray is {60}
*   
*   Input: arr[] = {-2, -3, 0, -2, -40}
*   Output:   80  // The subarray is {-2, -40}
*               
**********************************************************************************/

func maxProduct(nums []int) int {
    max := nums[0]
    min := nums[0]
    maxAns := nums[0]
    
    for i := 1; i < len(nums); i++ {
        mx := max 
        mn := min
        
        max = Max(Max(nums[i], mx * nums[i]), mn * nums[i])
        min = Min(Min(nums[i], mx * nums[i]), mn * nums[i])
        
        maxAns = Max(max, maxAns)
    } 
    
    return maxAns
    
}

func Max(x, y int) int {
    if x > y {
        return x 
    }
    
    return y
}

func Min(x, y int) int {
    if x < y {
        return x
    }
    
    return y
}


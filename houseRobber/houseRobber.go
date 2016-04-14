// Source : https://leetcode.com/problems/house-robber/
// Author : Hdaoran
// Date   : 2016-04-14

/********************************************************************************** 
 * 
 * You are a professional robber planning to rob houses along a street. Each house has 
 * a certain amount of money stashed, the only constraint stopping you from robbing 
 * each of them is that adjacent houses have security system connected and it will 
 * automatically contact the police if two adjacent houses were broken into on the same night.
 * 
 * Given a list of non-negative integers representing the amount of money of each house, 
 * determine the maximum amount of money you can rob tonight without alerting the police.
 * 
 *               
 **********************************************************************************/

/*
 * Dynamic Programming
 *
 * We can easy find the recurive fomular:
 *
 *     dp[n] = max( 
 *                    dp[n-1],   // the previous house has been robbed. 
 *                    dp[n-2] + money[n]  // the previous house has NOT been robbed.
 *                )
 *                  
 * The initalization is obvious:
 *     dp[1] = money[1]
 *     dp[2] = max(money[1], money[2])
 *
 */
 
func rob(nums []int) int {
    n2 := 0 // dp[i-2]
    n1 := 0 // dp[i-1]
    
    for i := 0; i < len(nums); i++ {
        current := max(n1, n2+nums[i])
        n2 = n1
        n1 = current
    }
    
    return n1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    
    return y
}
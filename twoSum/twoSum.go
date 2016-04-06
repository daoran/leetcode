// Source : https://oj.leetcode.com/problems/two-sum/
// Author : daoran
// Date   : 2016-4-6

/**********************************************************************************
*
* Given an array of integers, find two numbers such that they add up to a specific target number.
*
* The function twoSum should return indices of the two numbers such that they add up to the target,
* where index1 must be less than index2. Please note that your returned answers (both index1 and index2)
* are not zero-based.
*
* You may assume that each input would have exactly one solution.
*
* Input: numbers={2, 7, 11, 15}, target=9
* Output: index1=1, index2=2
*
*
**********************************************************************************/
package twoSum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hash[target-nums[i]]; ok {
			result = []int{j, i}
			break
		}
		hash[nums[i]] = i
	}

	return result
}

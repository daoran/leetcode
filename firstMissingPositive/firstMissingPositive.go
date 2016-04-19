// Source : https://oj.leetcode.com/problems/first-missing-positive/
// Author : daoran
// Date   : 2016-04-19

/**********************************************************************************
*
* Given an unsorted integer array, find the first missing positive integer.
*
* For example,
* Given [1,2,0] return 3,
* and [3,4,-1,1] return 2.
*
* Your algorithm should run in O(n) time and uses constant space.
*
*
**********************************************************************************/

/*
 *  Idea:
 *
 *    We can move the num to the place whcih the index is the num.
 *
 *    for example,  (considering the array is zero-based.
 *       1 => A[0], 2 => A[1], 3=>A[2]
 *
 *    Then, we can go through the array check the i+1 == A[i], if not ,just return i+1;
 *
 *    This solution comes from StackOverflow.com
 *    http://stackoverflow.com/questions/1586858/find-the-smallest-integer-not-in-a-list
 */

package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 1
	}

	var num int
	for i := 0; i < len(nums); i++ {
		num = nums[i]
		for num > 0 && num < n && nums[num-1] != num {
			nums[i], nums[num-1] = nums[num-1], nums[i]
			num = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

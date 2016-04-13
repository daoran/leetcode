// Source : https://leetcode.com/problems/rotate-array/
// Author : daoran
// Date   : 2016-04-13

/**********************************************************************************
*
* Rotate an array of n elements to the right by k steps.
* For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
*
* Note:
* Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
*
* Hint:
* Could you do it in-place with O(1) extra space?
*
* Related problem: Reverse Words in a String II
*
* Credits:Special thanks to @Freezen for adding this problem and creating all test cases.
*
**********************************************************************************/

package main

import "fmt"

/*
 * this solution is so-called three times rotate method
 * because (X^TY^T)^T = YX, so we can perform rotate operation three times to get the result
 * obviously, the algorithm consumes O(1) space and O(n) time
 */

func reverseArray(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate(nums []int, k int) {
	if k <= 0 {
		return
	}

	n := len(nums)
	k %= n
	reverseArray(nums, n-k, n-1)
	reverseArray(nums, 0, n-k-1)
	reverseArray(nums, 0, n-1)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

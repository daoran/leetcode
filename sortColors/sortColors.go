// Source : https://oj.leetcode.com/problems/sort-colors/
// Author : daoran
// Date   : 2016-04-16

/**********************************************************************************
*
* Given an array with n objects colored red, white or blue, sort them so that objects of
* the same color are adjacent, with the colors in the order red, white and blue.
*
* Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
*
* Note:
* You are not suppose to use the library's sort function for this problem.
*
* Follow up:
*  > A rather straight forward solution is a two-pass algorithm using counting sort.
*  > First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array
*    with total number of 0's, then 1's and followed by 2's.
*  > Could you come up with an one-pass algorithm using only constant space?
*
**********************************************************************************/

package main

import "fmt"

func sortColors(nums []int) {
	zero := 0
	two := len(nums) - 1

	for i := 0; i <= two; i++ {
		if nums[i] == 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			zero++
		}

		if nums[i] == 2 {
			nums[two], nums[i] = nums[i], nums[two]
			two--
			i--
		}
	}
}

func main() {
	test := []int{0, 1, 2, 2, 1, 1, 0}
	sortColors(test)
	fmt.Println(test)
}

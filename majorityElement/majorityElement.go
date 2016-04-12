// Source : https://oj.leetcode.com/problems/majority-element/
// Author : daoran
// Date   : 2016-04-12

/**********************************************************************************
 *
 * Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always exist in the array.
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import "fmt"

func majorityElement(nums []int) int {
	majority := 0
	cnt := 0

	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			majority = nums[i]
			cnt++
		} else {
			if majority == nums[i] {
				cnt++
			} else {
				cnt--
			}
		}
	}

	return majority
}

func main() {
	fmt.Println(majorityElement([]int{7, 7, 8, 8, 8}))
}

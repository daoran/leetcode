// Source : https://oj.leetcode.com/problems/single-number/
// Author : daoran
// Date   : 2016-04-11

/**********************************************************************************
*
* Given an array of integers, every element appears twice except for one. Find that single one.
*
* Note:
* Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
*
*
**********************************************************************************/

package main

import "fmt"

func singleNumber(nums []int) int {
	num := 0
	for _, n := range nums {
		num ^= n
	}
	return num
}

func main() {
	fmt.Println(singleNumber([]int{-4, 2, 2, 3, 3}))
}

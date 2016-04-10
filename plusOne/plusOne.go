// Source : https://oj.leetcode.com/problems/plus-one/
// Author : daoran
// Date   : 2016-04-10

/**********************************************************************************
*
* Given a non-negative number represented as an array of digits, plus one to the number.
*
* The digits are stored such that the most significant digit is at the head of the list.
*
**********************************************************************************/

package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if digit < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{1, 0, 0, 9}))
}

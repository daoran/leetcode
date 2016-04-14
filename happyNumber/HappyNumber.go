// Source : https://leetcode.com/problems/happy-number/
// Author : daoran
// Date   : 2016-04-14

/**********************************************************************************
 *
 * Write an algorithm to determine if a number is "happy".
 *
 * A happy number is a number defined by the following process: Starting with any positive integer,
 * replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1
 * (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this
 * process ends in 1 are happy numbers.
 *
 * Example: 19 is a happy number
 *
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 * Credits:Special thanks to @mithmatt and @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/
package main

import "fmt"

func squares(n int) int {
	result := 0
	sq := 0
	for ; n > 0; n /= 10 {
		sq = n % 10
		result += sq * sq
	}
	return result
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	m := make(map[int]struct{})
	m[n] = struct{}{}

	for n != 1 {
		n = squares(n)
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(isHappy(19))
}

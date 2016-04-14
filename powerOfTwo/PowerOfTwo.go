// Source : https://leetcode.com/problems/power-of-two/
// Author : daoran
// Date   : 2016-04-14

/**********************************************************************************
 *
 * Given an integer, write a function to determine if it is a power of two.
 *
 * Credits:Special thanks to @jianchao.li.fighter for adding this problem and creating
 * all test cases.
 *
 *
 *
 **********************************************************************************/

package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	} else {
		return (n & (n - 1)) == 0
	}
}

func main() {
	fmt.Println(isPowerOfTwo(0))
}

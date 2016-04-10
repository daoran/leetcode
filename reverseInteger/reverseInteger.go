// Source : https://oj.leetcode.com/problems/reverse-integer/
// Author : daoran
// Date   : 2016-04-10

/**********************************************************************************
*
* Reverse digits of an integer.
*
* Example1: x =  123, return  321
* Example2: x = -123, return -321
*
*
* Have you thought about this?
*
* Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!
*
* > If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.
*
* > Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer,
*   then the reverse of 1000000003 overflows. How should you handle such cases?
*
* > Throw an exception? Good, but what if throwing an exception is not an option?
*   You would then have to re-design the function (ie, add an extra parameter).
*
*
**********************************************************************************/

package main

import "fmt"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		if abs(ret) > 214748364 {
			return 0
		}

		ret = ret*10 + x%10
		x /= 10
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}

	return x
}

func main() {
	fmt.Println(reverse(-2147483412))
	fmt.Println(reverse(-123456))
}

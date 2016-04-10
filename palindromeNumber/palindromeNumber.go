// Source : https://oj.leetcode.com/problems/palindrome-number/
// Author : daoran
// Date   : 2016-04-10

/**********************************************************************************
*
* Determine whether an integer is a palindrome. Do this without extra space.
*
*
* Some hints:
*
* Could negative integers be palindromes? (ie, -1)
*
* If you are thinking of converting the integer to string, note the restriction of using extra space.
*
* You could also try reversing an integer. However, if you have solved the problem "Reverse Integer",
* you know that the reversed integer might overflow. How would you handle such case?
*
* There is a more generic way of solving this problem.
*
*
**********************************************************************************/

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	div := 1
	for x/div >= 10 {
		div *= 10
	}

	for x != 0 {
		l := x / div
		r := x % 10

		if l != r {
			return false
		}

		x = (x % div) / 10
		div /= 100
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(1231))
}

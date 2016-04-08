// Source : https://oj.leetcode.com/problems/valid-palindrome/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
*
* For example,
* "A man, a plan, a canal: Panama" is a palindrome.
* "race a car" is not a palindrome.
*
* Note:
* Have you consider that the string might be empty? This is a good question to ask during an interview.
*
* For the purpose of this problem, we define empty string as valid palindrome.
*
*
**********************************************************************************/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	slice := []rune{}

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			slice = append(slice, r)
		}
	}

	left, right := 0, len(slice)-1
	for left < right {
		if slice[left] != slice[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}

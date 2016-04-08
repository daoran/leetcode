// Source : https://oj.leetcode.com/problems/length-of-last-word/
// Author : daoran
// Date   : 2016-04-08

/**********************************************************************************
*
* Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
* return the length of last word in the string.
*
* If the last word does not exist, return 0.
*
* Note: A word is defined as a character sequence consists of non-space characters only.
*
* For example,
* Given s = "Hello World",
* return 5.
*
*
**********************************************************************************/

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	slice := strings.Split(s, " ")
	result := 0
	for i := len(slice) - 1; i >= 0; i-- {
		if len(slice[i]) > 0 {
			result = len(slice[i])
			break
		}
	}

	return result
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("a "))
}
